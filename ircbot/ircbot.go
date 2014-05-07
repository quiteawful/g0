package IrcBot

import (
	"crypto/tls"
	//"errors"
	"github.com/thoj/go-ircevent"
	"log"
	"regexp"
)

var urlregex = regexp.MustCompile(`((([A-Za-z]{3,9}:(?:\/\/)?)(?:[-;:&=\+\$,\w]+@)?[A-Za-z0-9.-]+|(?:www.|[-;:&=\+\$,\w]+@)[A-Za-z0-9.-]+)((?:\/[\+~%\/.\w-_]*)?\??(?:[-\+=&;%@.\w_]*)#?(?:[\w]*))?)`)
var useCrypto = true

type Bot struct {
	Nickname    string "g0bot"
	Realname    string "g0bot"
	Connections []Conn
	LinkChannel chan Link
}

type Conn struct {
	Connection *irc.Connection
	Address    string
	Channels   []string
}

type Link struct {
	URL     string
	Network string
	Channel string
	Poster  string
}

func (b *Bot) Run() {
	ircCon := irc.IRC(b.Nickname, b.Realname)
	ircCon.VerboseCallbackHandler = false
	ircCon.UseTLS = true
	ircCon.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	ircErr := ircCon.Connect(b.Connections[0].Address)
	if ircErr != nil {
		log.Println(ircErr.Error())
	}
	ircCon.AddCallback("001", func(e *irc.Event) {
		for _, i := range b.Connections[0].Channels {
			log.Println("Joining: " + i)
			ircCon.Join(i)
		}

	})
	ircCon.AddCallback("PRIVMSG", func(e *irc.Event) {
		if urlregex.MatchString(e.Message()) {
			urlString := urlregex.FindStringSubmatch(e.Message())
			ircCon.Privmsg(e.Arguments[0], ">"+urlString[0])
			b.LinkChannel <- Link{urlString[0], "unknown", "#unknown", e.Nick}
		}
	})

	b.Connections = append(b.Connections, Conn{ircCon, b.Connections[0].Address, b.Connections[0].Channels})
	ircCon.Loop()
	log.Printf("IRC loop exited")
}

func parseIrc(e *irc.Event, ircCon *irc.Connection) {

}
