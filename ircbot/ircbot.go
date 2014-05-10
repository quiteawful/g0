package IrcBot

import (
	"crypto/tls"
	//"errors"
	"github.com/thoj/go-ircevent"
	"log"
	"regexp"
	"strings"
)

var chprefixes = map[uint8]bool{
	'#': true,
	'&': true,
	'!': true,
	'+': true,
}

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
	Network    string "unknownnetid"
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

	for _, i := range b.Connections {
		ircErr := ircCon.Connect(i.Address)
		if ircErr != nil {
			log.Println(ircErr.Error())
		}

		ircCon.AddCallback("005", func(e *irc.Event) {
			for _, j := range strings.Fields(e.Raw) {
				if len(j) > 7 && j[:7] == "NETWORK" {
					i.Network = j[8:]
					log.Println("Network name is: " + i.Network)
				}
			}
		})

		ircCon.AddCallback("001", func(e *irc.Event) {
			for _, j := range i.Channels {
				log.Println("Joining: " + j + " on " + i.Address)
				ircCon.Join(j)
			}
		})
		ircCon.AddCallback("PRIVMSG", func(e *irc.Event) {
			if urlregex.MatchString(e.Message()) {
				urlString := urlregex.FindStringSubmatch(e.Message())
				//ircCon.Privmsg(e.Arguments[0], ">"+urlString[0])

				if ircch := e.Arguments[0]; chprefixes[ircch[0]] {
					b.LinkChannel <- Link{urlString[0], i.Network, ircch, e.Nick}
				}
			}
		})
		i.Connection = ircCon
		//ircCon.Loop()
	}
}

func parseIrc(e *irc.Event, ircCon *irc.Connection) {

}
