package Db

import (
	"database/sql"
	"errors"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"strconv"
	"time"
)

type Image struct {
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

func (db *Db) tblImagesSetup() (bool, error) {
	query := "CREATE TABLE IF NOT EXISTS " + db.DbImageTable + "(" +
		"id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT, hash TEXT, " +
		"name TEXT, thumbnail TEXT, " +
		"tstamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP, url TEXT, " +
		"network TEXT, chan TEXT, user TEXT);"

	_, err := db.execute(query)
	if err != nil {
		log.Printf("db.tblImagesSetup: %s\n", err.Error())
		return false, err
	}
	return true, nil
}

func (db *Db) NewImage(hash, name, thumbnail, url, network, channel, user string) (id int64, err error) {
	if name == "" {
		err = errors.New("Empty filename")
		log.Printf("NewImage: %s\n", err.Error())
		return 0, err
	}

	query := "INSERT INTO " + db.DbImageTable + "(hash, name, thumbnail, url, network, chan, user) " +
		"values(?, ?, ?, ?, ?, ?, ?)"

	result, err := db.execute(query, hash, name, thumbnail, url, network, channel, user)
	if err != nil {
		log.Printf("Db.NewImage: %s\n", err.Error())
		return 0, err
	}
	id, err = result.LastInsertId()
	if err != nil {
		log.Printf("Db.NewImage: %s\n", err.Error())
		return id, err
	}
	return id, nil
}

func (db *Db) GetImage(id int) (result Image, err error) {
	if id < 1 {
		err = errors.New("No id found.")
		log.Printf("GetImage: %s\n", err.Error())
		return Image{}, err
	}

	query := "SELECT * FROM " + db.DbImageTable + " WHERE id = ?;"
	rows, err := db.query(query, id)

	if err != nil {
		log.Printf("Db.GetImage: %s\n", err.Error())
		return result, err
	}
	rows.Next() // no loop, only first result
	err = rows.Scan(&result.Id, &result.Hash, &result.Name, &result.Thumbnail,
		&result.Timestamp, &result.Url, &result.Network, &result.Channel, &result.Url)

	if err != nil {
		log.Printf("Db.GetImage: %s\n", err.Error())
		return result, err
	}
	if err == sql.ErrNoRows {
		err = errors.New("Query returned zero rows.")
		log.Printf("Db.GetImage: %s\n", err.Error())
		return result, err
	}
	return result, nil
}
func (db *Db) GetPreviousImagesBefore(id, n int) (result []Image, err error) {
	var strId string = strconv.Itoa(id)
	var strN string = strconv.Itoa(n)
	var query string = "SELECT * FROM " +
		db.DbImageTable +
		" where id > " +
		strId +
		" and id <= (" + strId + " + " + strN + ") order by id desc"

	rows, err := db.query(query)
	if err != nil {
		log.Printf("Db.GetLatestImages: %s\n", err)
		return nil, err
	}

	for rows.Next() {
		img := Image{}
		err = rows.Scan(
			&img.Id,
			&img.Hash,
			&img.Name,
			&img.Thumbnail,
			&img.Timestamp,
			&img.Url,
			&img.Network,
			&img.Channel,
			&img.User)

		if err != nil {
			log.Printf("Db.GetLatestImages: %s\n", err)
			return nil, err
		}

		result = append(result, img)
	}
	return result, nil

}

func (db *Db) GetLatestImages(id, n int) (result []Image, err error) {
	var query string = "SELECT * FROM " + db.DbImageTable
	if id > 0 {
		idend := id - n
		if idend < 1 { // do not accept negative values in where clause
			idend = 1
		}
		query += " WHERE id < " + strconv.Itoa(id) + " AND id > " + strconv.Itoa(idend) + " ORDER BY tstamp DESC, id DESC"

	} else {
		query += " ORDER BY id DESC LIMIT 0, " + strconv.Itoa(n)
	}

	rows, err := db.query(query)
	if err != nil {
		log.Printf("Db.GetLatestImages: %s\n", err)
		return nil, err
	}

	for rows.Next() {
		img := Image{}
		err = rows.Scan(
			&img.Id,
			&img.Hash,
			&img.Name,
			&img.Thumbnail,
			&img.Timestamp,
			&img.Url,
			&img.Network,
			&img.Channel,
			&img.User)

		if err != nil {
			log.Printf("Db.GetLatestImages: %s\n", err)
			return nil, err
		}

		result = append(result, img)
	}
	return result, nil
}

func (db *Db) GetImages(offset, n int) (result []Image, err error) {
	if offset < 1 {
		err = errors.New("Start id too low.")
		log.Printf("Db.GetImages: %s\n", err)
		return nil, err
	}
	if n < 1 {
		err = errors.New("Offset too low.")
		log.Printf("Db.GetImages: %s\n", err)
		return nil, err
	}

	query := "SELECT * FROM " + db.DbImageTable + " WHERE id >= ? AND id < ?"
	rows, err := db.query(query, offset, (offset + n))
	if err != nil {
		log.Printf("Db.GetImages: %s\n", err)
		return nil, err
	}

	for rows.Next() {
		img := Image{}
		err = rows.Scan(
			&img.Id,
			&img.Hash,
			&img.Name,
			&img.Thumbnail,
			&img.Timestamp,
			&img.Url,
			&img.Network,
			&img.Channel,
			&img.User)

		if err != nil {
			log.Printf("Db.GetImages: %s\n", err)
			return nil, err
		}

		result = append(result, img)
	}

	return result, nil
}

func (db *Db) DeleteImage(id int) bool {
	if id < 1 {
		return false
	}

	query := "DELETE FROM " + db.DbImageTable + " WHERE id = ?"
	result, err := db.execute(query, id)
	if err != nil {
		log.Printf("Db.DeleteImage: %s\n", err.Error())
		return false
	}
	affected, err := result.RowsAffected()
	if err != nil {
		log.Printf("Db.DeleteImage: %s\n", err.Error())
		return false
	}
	if affected != 1 {
		return false
	}
	return true
}

func (db *Db) GetImageCount() (c int, err error) {
	query := "SELECT count(*) FROM " + db.DbImageTable
	rows, err := db.query(query)
	if err != nil {
		log.Printf("Db.GetImageCount: %s\n", err.Error())
		return 0, err
	}
	rows.Next()
	err = rows.Scan(&c)

	if err == sql.ErrNoRows {
		err = errors.New("Query returned zero rows.")
		log.Printf("Db.GetImageCount: %s %s\n", err.Error(), query)
		return 0, err
	}
	return c, nil
}

func (db *Db) GetHashCount(hash string) (c int, err error) {
	query := "SELECT count(*) AS c FROM " + db.DbImageTable + " WHERE hash = ?"
	rows, err := db.query(query, hash)
	if err != nil {
		log.Printf("Db.GetHashCount: %s\n", err.Error())
		return -1, err
	}

	for rows.Next() {
		err = rows.Scan(&c)
		if err != nil {
			log.Printf("Db.GetHashCount: %s\n", err.Error())
			return -1, err
		}
	}

	return c, nil

}
