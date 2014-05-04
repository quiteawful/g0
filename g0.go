// reminder.go
package main

import (
	//"errors"
	"fmt"
	"g0/api"
<<<<<<< Updated upstream
	"g0/ircbot"
=======
	"g0/util"
	"github.com/thoj/go-ircevent"
>>>>>>> Stashed changes
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
<<<<<<< Updated upstream
=======
	_, err = util.DownloadImage("http://i.imgur.com/Uxq2dPU.gif")
	if err != nil {
		fmt.Println(err.Error())
	}
	os.Exit(0)
	ircCon := irc.IRC("Churchill", "PrimeMinister")
	ircCon.VerboseCallbackHandler = false
	ircCon.UseTLS = true
	ircCon.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	ircErr := ircCon.Connect("tardis.nerdlife.de:6697")
	if ircErr != nil {
		fmt.Println(ircErr.Error())
	}
	ircCon.AddCallback("001", func(e *irc.Event) {
		ircCon.Join("#g0")
	})
	ircCon.AddCallback("PRIVMSG", func(e *irc.Event) { parseIrc(e, ircCon) })
	ircCon.Loop()
}
>>>>>>> Stashed changes

	ircbot := IrcBot.NewBot("g0bot", "g0bot")
	go ircbot.Run("tardis.nerdlife.de:6697", "#amelie", "#g0")

	//hässliche blocking schleife ist hässlich
	for true {
		time.Sleep(10000000)
	}
}
