package main

import (
	"github.com/aimless/g0/api"
	"github.com/aimless/g0/conf"
	"github.com/aimless/g0/ircbot"
	"github.com/aimless/g0/util"
	"github.com/aimless/g0/util/img"
	"log"
)

func main() {
	api := new(Api.Api)
	bot := new(IrcBot.Bot)
	db := new(Db.DbConfig)

	// load config structs
	conf.Fill(api)
	conf.Fill(bot)
	conf.Fill(db)

	// init objects
	Db.Init(db)
	images := new(Db.Image)
	err := images.Setup()
	if err != nil {
		log.Printf("Main: %s\n", err.Error())
	}

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

		imgbytes, _ := img.GetImageFromFile(f)
		thmb, _ := img.MakeThumbnail(imgbytes, 150, 150)
		img.SaveImageAsJPG("thumb-"+f, thmb)

		dbase.NewImage(hash, f, "thumb"+f, link.URL, link.Network, link.Channel, link.Poster)
	}
}
