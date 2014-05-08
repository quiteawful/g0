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

	conf.Fill(api)
	conf.Fill(bot)

	go api.Run()
	go bot.Run()

	bot.LinkChannel = make(chan IrcBot.Link)

	db := new(Db.DbConfig)
	conf.Fill(db)
	dbase, _ := Db.NewDb(db.DbFile)

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
