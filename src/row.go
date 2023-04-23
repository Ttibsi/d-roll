package src

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type DbRow struct {
	Id    int
	Date  string
	Roll  string
	Total string
}

func GetData() (*[]DbRow, error) {
	db, err := sql.Open("sqlite3", "./db.db")
	if err != nil {
		return nil, err
	}

	rows, err := db.Query("SELECT * FROM results ASCENDING")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var data []DbRow
	for rows.Next() {
		var d DbRow
		err = rows.Scan(&d.Id, &d.Date, &d.Roll, &d.Total)
		if err != nil {
			return nil, err
		}
		data = append(data, d)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return &data, nil
}
