package IrcBot

import (
	"crypto/tls"
	//"errors"
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/quiteawful/g0/conf"
	"github.com/quiteawful/go-ircevent"
)

var chprefixes = map[uint8]bool{
	'#': true,
	'&': true,
	'!': true,
	'+': true,
}

var urlregex = regexp.MustCompile(`((([A-Za-z]{3,9}:(?:\/\/)?)(?:[-;:&=\+\$,\w]+@)?[A-Za-z0-9.-]+|(?:www.|[-;:&=\+\$,\w]+@)[A-Za-z0-9.-]+)((?:\/[\+~%\/.\w-_]*)?\??(?:[-\+=&;%@.\w_]*)#?(?:[\w]*))?)`)
var delregex = regexp.MustCompile(`!del (\d+)`)
var useCrypto = true

type Bot struct {
	Nickname    string "g0bot"
	Realname    string "g0bot"
	Connections []Conn
	LinkChannel chan Link
	DeleteImage chan int64
	SendChannel chan Send
}

type Conn struct {
	Connection *irc.Connection
	Address    string
	Network    string "unknownnetid"
	Channels   []string
}

type Send struct {
	IrcChan string
	Msg     string
}

type Link struct {
	URL     string
	Network string
	Channel string
	Poster  string
}

var (
	_bot *Bot = nil
)

// If shit's broken this is probably the reason for it.
func init() {
	if _bot == nil {
		_bot = new(Bot)
	}
	tmpBot := new(Bot)
	conf.Fill(tmpBot)
	_bot.Nickname = tmpBot.Nickname
	_bot.Realname = tmpBot.Realname
	_bot.Connections = tmpBot.Connections
}

func (b *Bot) Run() {
	ircCon := irc.IRC(b.Nickname, b.Realname)
	ircCon.VerboseCallbackHandler = false
	ircCon.UseTLS = true
	ircCon.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	for x, i := range b.Connections {
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
			parseIrcMsg(e, b)
		})
		b.Connections[x].Connection = ircCon
		go b.Connections[x].Connection.Loop()
	}
}

func parseIrcMsg(e *irc.Event, b *Bot) {
	// add images
	if urlregex.MatchString(e.Message()) {
		if strings.HasPrefix(e.Message(), "!nope") {
			return // nope-ing out
		}
		urlString := urlregex.FindStringSubmatch(e.Message())[0]
		//dont do shit if it is a aidkrebs link, but not i.aids...
		if strings.Contains(urlString, "aidskrebs") && !strings.Contains(urlString, "i.aidskrebs") {
			return
		}
		if ircch := e.Arguments[0]; chprefixes[ircch[0]] {
			b.LinkChannel <- Link{urlString, b.Connections[0].Network, ircch, e.Nick}
		}
	}
	//print help
	if e.Message() == "!halp" {
		if chprefixes[e.Arguments[0][0]] {
			printHalp(e.Arguments[0], b)
		}
	}
	// del image
	if delregex.MatchString(e.Message()) {
		id, err := strconv.ParseInt(delregex.FindStringSubmatch(e.Message())[1], 10, 64)
		if err != nil {
			log.Println("Parsing deletion-id:", err.Error())
			return
		}
		b.DeleteImage <- id
	}
}

func printHalp(ch string, b *Bot) {
	b.Connections[0].Connection.Privmsg(ch, "!nope <url>: skip link")
	b.Connections[0].Connection.Privmsg(ch, "!del <id>: delete image")
}
