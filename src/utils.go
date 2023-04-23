package src

import (
	"database/sql"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

	// _ "github.com/mattn/go-sqlite3"
)

func rollDice(d *Dice) {
	var results []int
	var total int
    rand.Seed(time.Now().UnixNano())

	for _, val := range d.Values {
		results = append(results, rand.Intn(val) + 1)
	}

	d.Display = strings.Trim(fmt.Sprint(results), "[]")

	for _, val := range results {
		total += val
	}

	d.Display += "<br>" //This works in html instead of a \n
	d.Display += "Result: "
	d.Display += strconv.Itoa(total)
}

func insertToDB() error {
	db, err := sql.Open("sqlite3", "./db.db")
	if err != nil {
		return err
	}

	stmt, err := db.Prepare("INSERT INTO results VALUES (NULL, ?, ?, ?);")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(time.Now().Format("2006-01-02"), )

	
	return nil
}
