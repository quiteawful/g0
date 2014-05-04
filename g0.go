// reminder.go
package main

import (
	//"errors"
	"encoding/json"
	"fmt"
	"g0/api"
	"g0/ircbot"
	"g0/util"
	"os"
	//"time"
)

type JSONconf struct {
	Imagepath string
	Thumbpath string
	Rest      *Api.Api
	Bot       *IrcBot.Bot
}

func main() {
	var err error
	conf := new(JSONconf)
	Init(conf)
	conf.Bot.LinkChannel = make(chan string)
	//hässliche blocking schleife ist hässlich
	for true {
		_, err = util.DownloadImage(<-conf.Bot.LinkChannel)
		if err != nil {
			fmt.Println(err.Error())
		}
	}
}

func Init(conf *JSONconf) {
	file, _ := os.Open("config.json")
	decoder := json.NewDecoder(file)
	err := decoder.Decode(conf)
	if err != nil {
		fmt.Println(err.Error())
	}
	go conf.Rest.Run()
	go conf.Bot.Run()
}
