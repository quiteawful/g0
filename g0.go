package main

import (
	//"errors"
	"encoding/json"
	"fmt"
	"g0/api"
	"g0/db"
	"g0/ircbot"
	"g0/util"
	"g0/util/img"
	"os"
	//"time"
)

type JSONconf struct {
	Imagepath string
	Thumbpath string
	DBpath    string
	Rest      *Api.Api
	Bot       *IrcBot.Bot
}

var Conf = new(JSONconf)

func main() {
	//Conf := new(JSONconf)
	Init(Conf)
	Conf.Bot.LinkChannel = make(chan IrcBot.Link)

	dbase, _ := db.NewDb(Conf.DBpath)

	//hässliche blocking schleife ist hässlich
	for true {
		link := <-Conf.Bot.LinkChannel
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

func Init(Conf *JSONconf) {
	file, _ := os.Open("config.json")
	decoder := json.NewDecoder(file)
	err := decoder.Decode(Conf)
	if err != nil {
		fmt.Println(err.Error())
	}
	go Conf.Rest.Run()
	go Conf.Bot.Run()
}
