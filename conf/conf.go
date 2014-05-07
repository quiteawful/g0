// conf project conf.go
package conf

import (
	"encoding/json"
	"fmt"
	"github.com/aimless/g0/api"
	"github.com/aimless/g0/ircbot"
	"os"
)

type JSONconf struct {
	Imagepath string
	Thumbpath string
	DBpath    string
	Rest      *Api.Api
	Bot       *IrcBot.Bot
}

var Conf = new(JSONconf)

func init() {

	file, _ := os.Open("config.json")
	decoder := json.NewDecoder(file)
	err := decoder.Decode(Conf)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
