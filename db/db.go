package Db

import (
	"database/sql"
	"errors"
	"github.com/aimless/g0/conf"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

type DbConfig struct {
	DbEngine  string
	DbFile    string
	TblImages string
	// Tbl$name for more tables in the database
}
type Db struct {
	DbEngine     string
	DbFile       string
	DbImageTable string

	conn *sql.DB
}

var (
	_db   *Db  = nil // private var of *Db
	debug bool = true
)

func NewDb() (*Db, error) {
	var err error
	// singleton
	if _db == nil {
		_db = new(Db)
	}
	// get config
	tmpConf := new(DbConfig)
	conf.Fill(tmpConf)

	// set values from config.json
	_db.DbEngine = tmpConf.DbEngine
	_db.DbFile = tmpConf.DbFile
	_db.DbImageTable = tmpConf.TblImages

	// open connection, and create tables if needed.
	_db.conn, err = sql.Open(_db.DbEngine, _db.DbFile /*+":locked.sqlite?chache=shared&mode=rwc"*/)
	if err != nil {
		log.Printf("Db.NewDb: Failed to open %s via driver: %s. Error: %s\n", _db.DbFile, _db.DbEngine, err.Error())
		return nil, err
	}
	result, err := _db.tblImagesSetup() // setup image table
	if err != nil {
		log.Printf("Db.NewDb: %s\n", err.Error())
		return nil, err
	}
	if result != true {
		err = errors.New("Could not create table " + _db.DbImageTable)
		log.Printf("Db.NewDb: %s\n", err.Error())
		return nil, err
	}

	// setup other tables if needed.
	return _db, nil
}

func (db *Db) Close() {
	db.conn.Close()
}

func (db *Db) execute(query string, args ...interface{}) (result sql.Result, err error) {
	if query == "" {
		err = errors.New("Empty query")
		log.Printf("Db.execute: %s\n", err.Error())
		return result, err
	}

	if debug {
		log.Printf("Query: (%v) %s\n", &db, query)
	}
	result, err = db.conn.Exec(query, args...)
	if err != nil {
		log.Printf("Db.execute: %s\n", err.Error())
		return result, err
	}
	return result, nil
}

func (db *Db) query(query string, args ...interface{}) (result *sql.Rows, err error) {
	if query == "" {
		err = errors.New("Empty query")
		log.Printf("Db.query: %s\n", err.Error())
		return result, err
	}
	if debug {
		log.Printf("Query: (%v) %s\n", &db, query)
	}
	result, err = db.conn.Query(query, args...)
	if err != nil {
		log.Printf("Db.query: %s: %s\n", err.Error(), query)
		return result, err
	}

	return result, nil
}
