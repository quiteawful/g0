// conf project conf.go
package conf

import (
	"encoding/json"
	"log"
	"os"
)

<<<<<<< HEAD
type JSONconf struct {
	Imagepath string
	Thumbpath string
	Db        *DbConfig
	Rest      *Api.Api
	Bot       *IrcBot.Bot
}

// Public Config struct for json parser.
type DbConfig struct {
	DbEngine  string
	DbFile    string
	TblImages string
	// Tbl$name for more tables in the database
	// and add Tbl$name in config.json
}

var (
	Imagepath = ""
	Thumbpath = ""
	Data      = new(DbConfig)
	Rest      = new(Api.Api)
	Bot       = new(IrcBot.Bot)
)

var c = new(JSONconf)

func init() {
	file, _ := os.Open("config.json")
	decoder := json.NewDecoder(file)
	err := decoder.Decode(c)
=======
func Fill(s interface{}) error {
	file, err := os.Open("config.json")
>>>>>>> dbfeefae85b90a4d0f6b5f05dececc0fde9cdb79
	if err != nil {
		log.Println("Could not open config file", err)
	}

	dec := json.NewDecoder(file)
	err = dec.Decode(s)
	if err != nil {
		log.Println("Config parser: ", err, s)
		return err
	}

	return nil
}
