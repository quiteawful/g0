package Db

import (
	_ "database/sql"
	"errors"
	"github.com/aimless/g0/conf"
	"log"
	"time"
)

type Image struct {
	// public Fields
	Id        int
	Hash      string
	Name      string
	Thumbnail string
	Timestamp time.Time
	Url       string
	Network   string
	Channel   string
	User      string
}

var (
	structDb *Db
)

func (img *Image) Save(i Image) (id int, err error) {
	if i.Name == "" {
		err = errors.New("Name property is not set.")
		log.Printf("Image.Save: %s\n", err.Error())
		return 0, err
	}

	query := "INSERT INTO " + conf.Data.TblImages +
		"(hash, name, thumbnail, url, network, chan, user) values(" +
		"?, ?, ?, ?, ?, ?, ?);"

	result, err := structDb.Exec2(query, i.Hash, i.Name, i.Thumbnail, i.Url,
		i.Network, i.Channel, i.User)

	if err != nil {
		log.Printf("Image.Save: %s\n", err.Error())
		return 0, err
	}

	id64, err := result.LastInsertId() // Get inserted id
	if err != nil {
		log.Printf("Image.Save: %s\n", err.Error())
		return 0, err
	}

	// http://stackoverflow.com/a/6878625/1374884
	const maxuint = ^uint(0)
	const maxint = int(maxuint >> 1)
	if id64 > int64(maxint) { // overflow check
		log.Printf("Image.Save: Inserted id too big for 32-bit integer. %s\n", err.Error())
		return 0, nil // err=nil due to the fact, that the insert succeeded.
	}
	return int(id64), nil
}

func (img *Image) Open(id int) (Image, error) {
	return Image{}, nil
}

func (img *Image) OpenAll(id, count int) ([]Image, error) {
	return []Image{}, nil
}
