package main

import (
	"github.com/aimless/g0/conf"
	"github.com/aimless/g0/db"
	"github.com/aimless/g0/ircbot"
	"github.com/aimless/g0/util"
	"github.com/aimless/g0/util/img"
	"log"
)

func main() {
	Init("init")
	conf.Bot.LinkChannel = make(chan IrcBot.Link)

	dbase, _ := Db.NewDb(conf.Data.DbFile)

	//hässliche blocking schleife ist hässlich
	for true {
		link := <-conf.Bot.LinkChannel
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

func Init(placeholder string) {
	go conf.Rest.Run()
	go conf.Bot.Run()
}
