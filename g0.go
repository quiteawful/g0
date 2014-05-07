package main

import (
	"fmt"
	"github.com/aimless/g0/api"
	"github.com/aimless/g0/conf"
	"github.com/aimless/g0/db"
	"github.com/aimless/g0/ircbot"
	"github.com/aimless/g0/util"
	"github.com/aimless/g0/util/img"
)

type JSONconf struct {
	Imagepath string
	Thumbpath string
	DBpath    string
	Rest      *Api.Api
	Bot       *IrcBot.Bot
}

func main() {
	Init("init")
	conf.Conf.Bot.LinkChannel = make(chan IrcBot.Link)

	dbase, _ := db.NewDb(conf.Conf.DBpath)

	//hässliche blocking schleife ist hässlich
	for true {
		link := <-conf.Conf.Bot.LinkChannel
		f, hash, err := util.DownloadImage(link.URL)
		if err != nil {
			fmt.Println(err.Error())
			continue
		}

		imgbytes, _ := img.GetImageFromFile(f)
		thmb, _ := img.MakeThumbnail(imgbytes, 150, 150)
		img.SaveImageAsJPG("thumb-"+f, thmb)

		dbase.NewImage(hash, f, "thumb"+f, link.URL, link.Network, link.Channel, link.Poster)
	}
}

func Init(placeholder string) {
	go conf.Conf.Rest.Run()
	go conf.Conf.Bot.Run()
}
