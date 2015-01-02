package main

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/lucron/g0/util/img"
	"github.com/quiteawful/g0/api"
	"github.com/quiteawful/g0/conf"
	"github.com/quiteawful/g0/db"
	"github.com/quiteawful/g0/ircbot"
	"github.com/quiteawful/g0/util"
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

	for {
		select {
		case link := <-bot.LinkChannel:
			saveImage(link, dbase)
		case id := <-bot.DeleteImage:
			dbase.DeleteImage(id)
		}
	}
}

func saveImage(link IrcBot.Link, dbase *Db.Db) {
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
		thumbfile = "tmp.jpg"
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
