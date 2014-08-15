package Db

import (
	_ "github.com/mattn/go-sqlite3"
	"log"
)

type Pair struct {
	User  string
	Count int
}

func (db *Db) GetStatistics() []Pair {
	result := []Pair{}
	query := "select count(*) as count, user from g0_images group by user order by count;"

	rows, err := db.query(query)
	if err != nil {
		log.Printf("Db.GetStatistics: %s\n", err.Error())
		return nil
	}

	for rows.Next() {
		p := Pair{}
		err = rows.Scan(&p.Count, &p.User)
		if err != nil {
			log.Printf("Db.GetStatistics: %s\n", err.Error())
			return nil
		}
		result = append(result, p)
	}

	return result

}
