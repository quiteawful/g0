package main

import (
	"github.com/aimless/g0/api"
	"github.com/aimless/g0/conf"
	"github.com/aimless/g0/db"
	"github.com/aimless/g0/ircbot"
	"github.com/aimless/g0/util"
	"github.com/aimless/g0/util/img"
	"log"
)

func main() {
	api := new(Api.Api)
	bot := new(IrcBot.Bot)
	dbase, _ := Db.NewDb()

	conf.Fill(api)
	conf.Fill(bot)

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

		imgbytes, err := img.GetImageFromFile(f)
		if err != nil {
			log.Printf("Main: %s\n", err.Error())
		}
		thmb, err := img.MakeThumbnail(imgbytes, 150, 150)
		if err != nil {
			log.Printf("Main: %s\n", err.Error())
		}
		img.SaveImageAsJPG("thumb-"+f, thmb)

		// check if the imagehash is already in the database
		hashcount, err := dbase.GetHashCount(hash)
		if err != nil {
			log.Printf("Main: %s\n", err.Error())
			continue
		}
		if hashcount > 0 || hashcount == -1 {
			// TODO remove thumbnail and image
			continue
		}

		dbase.NewImage(hash, f, "thumb-"+f, link.URL, link.Network, link.Channel, link.Poster)
	}
}
