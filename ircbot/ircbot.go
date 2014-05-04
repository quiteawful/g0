package IrcBot

import (
	"crypto/tls"
	//"errors"
	"fmt"
	"github.com/thoj/go-ircevent"
	"regexp"
)

var urlregex = regexp.MustCompile(`((([A-Za-z]{3,9}:(?:\/\/)?)(?:[-;:&=\+\$,\w]+@)?[A-Za-z0-9.-]+|(?:www.|[-;:&=\+\$,\w]+@)[A-Za-z0-9.-]+)((?:\/[\+~%\/.\w-_]*)?\??(?:[-\+=&;%@.\w_]*)#?(?:[\w]*))?)`)
var useCrypto = true

type Bot struct {
	Nickname    string "g0bot"
	Realname    string "g0bot"
	Connections []Conn
}

type Conn struct {
	Connection *irc.Connection
	Address    string
	Channels   []string
}

func NewBot(nickname string, realname string) *Bot {
	return &Bot{nickname, realname, make([]Conn, 0)}
}

func (b *Bot) Run(server string, channels ...string) {
	ircCon := irc.IRC(b.Nickname, b.Realname)
	ircCon.VerboseCallbackHandler = false
	ircCon.UseTLS = true
	ircCon.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	ircErr := ircCon.Connect(server)
	if ircErr != nil {
		fmt.Println(ircErr.Error())
	}
	ircCon.AddCallback("001", func(e *irc.Event) {
		for _, i := range channels {
			fmt.Println("Joining: " + i)
			ircCon.Join(i)
		}

	})
	ircCon.AddCallback("PRIVMSG", func(e *irc.Event) { parseIrc(e, ircCon) })
	b.Connections = append(b.Connections, Conn{ircCon, server, channels})
	ircCon.Loop()
	fmt.Printf("IRC loop exited")
}

func parseIrc(e *irc.Event, ircCon *irc.Connection) {
	if urlregex.MatchString(e.Message()) {
		urlString := urlregex.FindStringSubmatch(e.Message())
		ircCon.Privmsg(e.Arguments[0], ">"+urlString[0])
	}
}
