// reminder.go
package main

import (
	//"errors"
	"fmt"
	"g0/api"
	"g0/ircbot"
	"os"
	"time"
)

func main() {
	api, err := Api.NewApi(":31337")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(-1)
	}
	go api.Run()

	ircbot := IrcBot.NewBot("g0bot", "g0bot")
	go ircbot.Run("tardis.nerdlife.de:6697", "#amelie", "#g0")

	//hässliche blocking schleife ist hässlich
	for true {
		time.Sleep(10000000)
	}
}
