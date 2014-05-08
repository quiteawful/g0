package Db

import (
	"database/sql"
	"errors"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

var (
	DbEngine string = "sqlite3"
	DbFile   string = "g0.db"

	Connection *sql.DB
)

/*
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
*/
func Open() error {
	var err error
	if Connection == nil {
		Connection, err = sql.Open(DbEngine, DbFile)
		if err != nil {
			log.Printf("Db.Open: %s\n", err.Error())
			return err
		}

	}
	return nil
}

func Close() {
	Connection.Close()
}

func Exec(query string, args ...interface{}) (sql.Result, error) {
	var err error
	if query == "" {
		err = errors.New("Query parameter is emtpy.")
		log.Printf("Db.Exec2: %s\n", err.Error())
		return nil, err
	}

	err = Open() // open db Connection
	if err != nil {

	}
	result, err := Connection.Exec(query, args...)
	return result, err
}

func Query(query string, args ...interface{}) (*sql.Rows, error) {
	var err error
	if query == "" {
		err = errors.New("Query parameter is emtpy.")
		log.Printf("Db.Query: %s\n", err.Error())
		return nil, err
	}

	err = Open() // open db Connection
	if err != nil {
		log.Printf("Db.Query: %s\n", err.Error())
		return nil, err
	}
	rows, err := Connection.Query(query, args...)
	if err != nil {
		log.Printf("Db.Query: %s\n", err.Error())
		return nil, err
	}
	return rows, nil
}

// Ist mit Vorsicht zu genießen, hrm hm mmhm
func Select(fields, from, where string) (*sql.Rows, error) {
	query := "SELECT " + fields + " FROM " + from
	if where != "" {
		query += " WHERE " + where
	}
	rows, err := Query(query)
	return rows, err
}
