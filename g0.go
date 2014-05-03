// reminder.go
package main

import (
	"crypto/tls"
	//"errors"
	"fmt"
	"g0/api"
	"github.com/thoj/go-ircevent"
	"os"
	"regexp"
)

var urlregex = regexp.MustCompile(`((([A-Za-z]{3,9}:(?:\/\/)?)(?:[-;:&=\+\$,\w]+@)?[A-Za-z0-9.-]+|(?:www.|[-;:&=\+\$,\w]+@)[A-Za-z0-9.-]+)((?:\/[\+~%\/.\w-_]*)?\??(?:[-\+=&;%@.\w_]*)#?(?:[\w]*))?)`)

func main() {
	api, err := Api.NewApi(":31337")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(-1)
	}
	go api.Run()
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

func parseIrc(e *irc.Event, ircCon *irc.Connection) {
	if urlregex.MatchString(e.Message()) {
		urlString := urlregex.FindStringSubmatch(e.Message())
		ircCon.Privmsg(e.Arguments[0], ">"+urlString[0])
	}
}
