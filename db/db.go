package db

import (
	"database/sql"
	"errors"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
	"time"
)

type Db struct {
	DbFile       string
	DbImageTable string
	conn         *sql.DB
}

type Image struct {
	id        int
	hash      string
	name      string
	thumbnail string
	timestamp time.Time
	url       string
	network   string
	channel   string
	user      string
}

func NewDb(DbFile string) (*Db, error) {
	var err error
	if DbFile == "" {
		return nil, errors.New("empty db DbFile")
	}
	_db := &Db{}
	_db.DbFile = DbFile
	_db.DbImageTable = "g0_images"
	_db.conn, err = sql.Open("sqlite3", _db.DbFile)
	if err != nil {
		log.Fatalf("NewDb: Failed to open DbFile. Error: %s\n", err.Error())
		return nil, err
	}
	if _, err = os.Stat(_db.DbFile); os.IsNotExist(err) {
		// db file does not exist, create new
		sql := "create table " + _db.DbImageTable + "(" +
			"id integer not null primary key autoincrement, " +
			"hash text, " +
			"name text, " +
			"thumbnail text, " +
			"tstamp timestamp default current_timestamp, " +
			"url text, " +
			"network text, " +
			"chan text, " +
			"user text" +
			")"
		_, err = _db.Exec(sql)
		if err != nil {
			log.Fatalf("NewDb: Failed to execute query: %s Error:%s\n", sql, err.Error())
			return nil, err
		}
	}

	return _db, nil
}

func (db *Db) Close() {
	db.conn.Close()
}

func (db *Db) Exec(query string) (sql.Result, error) {
	var err error
	if query == "" {
		err = errors.New("Empty query")
		log.Fatalf("Exec: Error:%s\n", err.Error())
		return nil, err
	}
	result, err := db.conn.Exec(query)
	if err != nil {
		log.Fatalf("Exec: Failed to execute query: %s Error:%s\n", query, err.Error())
		return nil, err
	}
	return result, nil
}

func (db *Db) NewImage(hash, name, thumbnail, url, network, channel, user string) (int64, error) {
	var err error
	if name == "" {
		err = errors.New("Empty filename")
		log.Fatalf("NewImage: %s\n", err.Error())
		return 0, err
	}

	sql := "insert into " + db.DbImageTable + "(hash, name, thumbnail, url, network, chan, user) values('" +
		hash + "', '" +
		name + "', '" +
		thumbnail + "', '" +
		url + "', '" +
		network + "', '" +
		channel + "', '" +
		user + "');"

	result, err := db.Exec(sql)
	if err != nil {
		log.Fatalf("NewImage: %s\n", err.Error())
		return 0, err
	}
	id, _ := result.LastInsertId()
	return id, nil
}

func (db *Db) GetImage(id int) (Image, error) {
	var err error
	if id < 1 {
		err = errors.New("No id found.")
		log.Fatalf("GetImage: %s\n", err.Error())
		return Image{}, err
	}

	sql := "select * from " + db.DbImageTable + " where id = ?"
	row := db.conn.QueryRow(sql, id)

	result := Image{}
	row.Scan(
		&result.id,
		&result.hash,
		&result.name,
		&result.thumbnail,
		&result.timestamp,
		&result.url,
		&result.network,
		&result.channel,
		&result.user)
	return result, nil
}

func (db *Db) GetImages(start, offset int) ([]Image, error) {
	var err error
	if start < 1 {
		err = errors.New("Start id too low.")
		log.Fatalf("GetImages: %s\n", err)
		return nil, err
	}
	if offset < 1 {
		err = errors.New("Offset too low.")
		log.Fatalf("GetImages: %s\n", err)
		return nil, err
	}

	sql := "select * from " + db.DbImageTable + " where id >= ? and id < ?"
	rows, err := db.conn.Query(sql, start, (start + offset))
	if err != nil {
		log.Fatalf("GetImages: %s\n", err)
		return nil, err
	}

	var result []Image
	for rows.Next() {
		img := Image{}
		err = rows.Scan(
			&img.id,
			&img.hash,
			&img.name,
			&img.thumbnail,
			&img.timestamp,
			&img.url,
			&img.network,
			&img.channel,
			&img.user)

		if err != nil {
			log.Fatalf("GetImages: %s\n", err)
			return nil, err
		}

		result = append(result, img)
	}

	return result, nil
}
