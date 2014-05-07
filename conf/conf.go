// conf project conf.go
package conf

import (
	"encoding/json"
	"github.com/aimless/g0/api"
	"github.com/aimless/g0/ircbot"
	"log"
	"os"
)

type JSONconf struct {
	Imagepath string
	Thumbpath string
	DBpath    string
	Rest      *Api.Api
	Bot       *IrcBot.Bot
}

var (
	Imagepath = ""
	Thumbpath = ""
	DBpath    = ""
	Rest      = new(Api.Api)
	Bot       = new(IrcBot.Bot)
)

var c = new(JSONconf)

func init() {
	file, _ := os.Open("config.json")
	decoder := json.NewDecoder(file)
	err := decoder.Decode(c)
	if err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}

	Imagepath = c.Imagepath
	Thumbpath = c.Thumbpath
	DBpath = c.DBpath
	Rest = c.Rest
	Bot = c.Bot

	log.Println("Parsed the following values: ", Imagepath, Thumbpath, DBpath, Rest, Bot)
}
