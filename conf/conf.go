// conf project conf.go
package conf

import (
	"encoding/json"
	"log"
	"os"
)

func Fill(s interface{}) error {
	file, err := os.Open("config.json")
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
