package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/quiteawful/g0/api"
	"github.com/quiteawful/g0/conf"
	"github.com/quiteawful/g0/db"
	"github.com/quiteawful/g0/ircbot"
	"github.com/quiteawful/g0/util"
	"github.com/quiteawful/g0/util/img"
)

var (
	_util *util.ConfImg = nil
)

func main() {

	api := new(Api.Api)
	bot := new(IrcBot.Bot)
	dbase, _ := Db.NewDb()
	if _util == nil {
		_util = new(util.ConfImg)
	}
	conf.Fill(api)
	conf.Fill(bot)
	conf.Fill(_util)
	go api.Run()
	go bot.Run()

	bot.LinkChannel = make(chan IrcBot.Link)
	bot.DeleteImage = make(chan int64)
	bot.SendChannel = make(chan IrcBot.Send)

	for {
		select {
		case link := <-bot.LinkChannel:
			saveImage(link, dbase, bot)
		case id := <-bot.DeleteImage:
			dbase.DeleteImage(id)
			// TODO: delete images from harddrive
		case msg := <-bot.SendChannel:
			bot.Connections[0].Connection.Privmsg(msg.IrcChan, msg.Msg)
		}
	}
}

// saveImage(...) is a mess.
// TODO: refactor.
func saveImage(link IrcBot.Link, dbase *Db.Db, bot *IrcBot.Bot) {
	//download image, hash it
	f, hash, err := util.DownloadImage(link.URL)
	if err != nil {
		log.Println("Downloading image:", err.Error())
		return
	}
	//check if image already exists
	hashcount, err := dbase.GetHashCount(hash)
	if err != nil {
		log.Println("Checking for duplicate:", err.Error())
		return
	}
	if hashcount > 0 || hashcount == -1 {
		// TODO: remove image from disk

		//when does hashcount return -1? O.o
		if hashcount > 0 {
			img, err := dbase.GetImageByHash(hash)
			if err != nil {
				log.Println("Sending AAAALT-Infos:", err.Error())
			}
			fmtstr := fmt.Sprintf("AAAALT: http://aidskrebs.net/?offset=%s", img.Id)
			bot.SendChannel <- IrcBot.Send{link.Channel, fmtstr}
		}
		return
	}
	//open image
	thumbfile := f
	tmpbyte, err := ioutil.ReadFile(_util.Imagepath + f)
	if err != nil {
		log.Println("Opening file:", err.Error())
		return
	}
	//check mimetype
	mime := http.DetectContentType(tmpbyte)
	if mime == "video/webm" {
		//workaround for something i dont remember
		thumbfile = "tmp.jpeg"
	}
	//open image...?
	imgbytes, err := img.GetImageFromFile(thumbfile)
	if err != nil {
		log.Println("Opening image:", err.Error())
		return
	}
	//create thumbnail
	thmb, err := img.MakeThumbnail(imgbytes, 150, 150)
	if err != nil {
		log.Println("Creating thumbnail:", err.Error())
		return
	}
	err = img.SaveImageAsJPG("thumb-"+hash+".jpg", thmb)
	if err != nil {
		log.Println("Saving thumb:", err.Error())
		return
	}
	dbase.NewImage(hash, f, "thumb-"+hash+".jpg", link.URL, link.Network, link.Channel, link.Poster)
}
