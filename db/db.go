package Db

import (
	"database/sql"
	"errors"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
)

type DbConfig struct {
	DbFile    string
	TblImages string
	// Tbl$name for more tables in the database
}
type Db struct {
	DbFile       string
	DbImageTable string

	conn *sql.DB
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

		log.Printf("NewDb: Failed to open DbFile. Error: %s\n", err.Error())
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
			log.Printf("NewDb: Failed to execute query: %s Error:%s\n", sql, err.Error())
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
		log.Printf("Exec: Error:%s\n", err.Error())
		return nil, err
	}
	result, err := db.conn.Exec(query)
	if err != nil {
		log.Printf("Exec: Failed to execute query: %s Error:%s\n", query, err.Error())
		return nil, err
	}
	return result, nil
}

func (db *Db) execute(query string, args ...interface{}) (result sql.Result, err error) {
	if query == "" {
		err = errors.New("Empty query")
		log.Printf("Db.execute: %s\n", err.Error())
		return result, err
	}

	result, err = db.conn.Exec(query, args...)
	if err != nil {
		log.Printf("Db.execute: %s\n", err.Error())
		return result, err
	}
	return result, nil
}
