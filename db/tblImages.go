package Db

import (
	_ "database/sql"
	"errors"
	"log"
	"strconv"
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

func (img *Image) Setup() error {
	// Creates the table inside the databasefile
	// if the table exists, nothing will be done
	query := "CREATE TABLE g0_images();"
	result, err := structDb.Exec2(query)
	return nil
}

func (img *Image) Save(i Image) (id int, err error) {
	if i.Name == "" {
		err = errors.New("Name property is not set.")
		log.Printf("Image.Save: %s\n", err.Error())
		return 0, err
	}

	query := "INSERT INTO g0_images" + // + conf.Data.TblImages +
		"(hash, name, thumbnail, url, network, chan, user) values(" +
		"?, ?, ?, ?, ?, ?, ?);"

	result, err := structDb.Exec(query, i.Hash, i.Name, i.Thumbnail, i.Url,
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
	// miiiight fail :>
	const maxuint = ^uint(0)
	const maxint = int(maxuint >> 1)
	if id64 > int64(maxint) { // overflow check
		log.Printf("Image.Save: Inserted id too big for 32-bit integer. %s\n", err.Error())
		return 0, nil // err=nil due to the fact, that the insert succeeded.
	}
	return int(id64), nil
}

func (img *Image) Open(id int) (r Image, err error) {
	if id < 1 {
		err = errors.New("Input parameter id is not set.")
		log.Printf("Image.Open: %s\n", err.Error())
		return Image{}, err
	}
	i := Image{}
	query := "SELECT * FROM g0_images WHERE id = ?"
	row := structDb.Exec(query, id)

	err = row.Scan(&i.Id, &i.Hash, &i.Name, &i.Thumbnail, &i.Timestamp, &i.Url,
		&i.Network, &i.Channel, &i.User)

	if err != nil {
		log.Printf("Image.Open: %s\n", err.Error())
		return Image{}, err
	}

	return i, nil
}

func (img *Image) OpenAll() ([]Image, error) {
	var result []Image
	query := "SELECT * FROM g0_images;"
	rows, err := structDb.Query(query)
	if err != nil {
		log.Printf("Image.OpenAll: %s\n", err.Error())
		return result, err
	}

	// scanning
	for rows.Next() {
		i := new(Image)
		err = rows.Scan(&i.Id, &i.Hash, &i.Name, &i.Thumbnail, &i.Timestamp, &i.Url,
			&i.Network, &i.Channel, &i.User)

		if err != nil {
			log.Printf("Image.OpenAll: %s\n", err.Error())
			continue
		}
		result = append(result, i)
	}

	return result, nil
}

func (img *Image) OpenLatest(id, n int) ([]Image, error) {
	var result []Image
	var query string

	if id > 0 {
		idend := id - n
		if idend < 1 { // do not accept negative values in where clause
			idend = 1
		}
		query = "select * from g0_images where id <= " + strconv.Itoa(id) + " and id > " + strconv.Itoa(idend) + " order by tstamp desc"

	} else {
		query = "select * from g0_images order by id desc limit 0, " + strconv.Itoa(n)
	}

	rows, err := structDb.Query(query)
	if err != nil {
		log.Printf("Image.OpenLatest: %s\n", err.Error())
		return result, err
	}

	// scanning
	for rows.Next() {
		i := new(Image)
		err = row.Scan(&i.Id, &i.Hash, &i.Name, &i.Thumbnail, &i.Timestamp, &i.Url,
			&i.Network, &i.Channel, &i.User)

		if err != nil {
			log.Printf("Image.OpenAll: %s\n", err.Error())
			continue
		}
		result = append(result, i)
	}

	return result, nil
}
