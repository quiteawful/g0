package Db

import (
	"database/sql"
	"errors"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

type DbConfig struct {
	DbEngine  string
	DbFile    string
	TblImages string
}

var (
	Connection *sql.DB = nil
	conf       DbConfig
)

func Init(c DbConfig) {
	conf.DbEngine = c.DbEngine
	conf.DbFile = c.DbFile
	conf.TblImages = c.TblImages
}

func Open() error {
	var err error
	if Connection == nil {
		Connection, err = sql.Open(conf.DbEngine, conf.DbFile)
		if err != nil {
			log.Printf("Db.Open: %s\n", err.Error())
			return err
		}
	}
	if Connection != nil {
		return nil
	}
	err = errors.New("hrmpf")
	return err
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
		log.Printf("Db.Exec: %s\n", err.Error())
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
