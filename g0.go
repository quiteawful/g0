package main

import (
	"github.com/quiteawful/g0/api"
	"github.com/quiteawful/g0/conf"
	"github.com/quiteawful/g0/db"
	"github.com/quiteawful/g0/ircbot"
	"github.com/quiteawful/g0/util"
	"github.com/quiteawful/g0/util/img"
	"io/ioutil"
	"log"
	"net/http"
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

	//hässliche blocking schleife ist hässlich
	for true {
		link := <-bot.LinkChannel
		f, hash, err := util.DownloadImage(link.URL)
		if err != nil {
			log.Println(err.Error())
			continue
		}

		// check if the imagehash is already in the database
		hashcount, err := dbase.GetHashCount(hash)
		if err != nil {
			log.Printf("Main: %s\n", err.Error())
			continue
		}
		if hashcount > 0 || hashcount == -1 {
			// TODO remove image
			// soda kann net warten, sonst haett ichs noch gemacht
			continue
		}

		thumbfile := f
		tmpbyte, err := ioutil.ReadFile(_util.Imagepath + thumbfile)
		if err != nil {
			log.Printf("main open file: %s\n", err.Error())
			continue
		}
		mime := http.DetectContentType(tmpbyte)
		if mime == "video/webm" {
			thumbfile = "tmp.jpeg"
		}
		imgbytes, err := img.GetImageFromFile(thumbfile)
		if err != nil {
			log.Printf("Main: %s\n", err.Error())
			continue
		}

		thmb, err := img.MakeThumbnail(imgbytes, 150, 150)
		if err != nil {
			dbase.NewImage(hash, f, "error.jpg", link.URL, link.Network, link.Channel, link.Poster)
			log.Printf("Main: %s\n", err.Error())
			continue
		}
		err = img.SaveImageAsJPG("thumb-"+hash+".jpg", thmb)
		if err != nil {
			log.Printf("Main: %s\n", err.Error())
			continue
		}
		dbase.NewImage(hash, f, "thumb-"+hash+".jpg", link.URL, link.Network, link.Channel, link.Poster)

	}
}
