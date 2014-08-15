package Db

import (
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func (db *Db) GetStatistics() map[string]int {
	result := make(map[string]int)
	query := "select count(*) as count, user from g0_images group by user union select count(*) as count, 'total' as user from g0_images order by count;"

	rows, err := db.query(query)
	if err != nil {
		log.Printf("Db.GetStatistics: %s\n", err.Error())
		return nil
	}

	var usr string
	var count int
	for rows.Next() {
		err = rows.Scan(&count, &usr)
		if err != nil {
			log.Printf("Db.GetStatistics: %s\n", err.Error())
			return nil
		}
		result[usr] = count
	}

	return result

}
