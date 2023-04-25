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
	rand.Seed(time.Now().UnixNano())

	for _, val := range d.Values {
		d.Rolls = append(d.Rolls, rand.Intn(val)+1)
	}

	d.Display = strings.Trim(fmt.Sprint(d.Rolls), "[]")

	for _, val := range d.Rolls {
		d.Total += val
	}

	d.Display += "<br>" //This works in html instead of a \n
	d.Display += "Result: "
	d.Display += strconv.Itoa(d.Total)
}

func insertToDB(d *Dice) error {
	db, err := sql.Open("sqlite3", "./db.db")
	if err != nil {
		return err
	}

	stmt, err := db.Prepare("INSERT INTO results VALUES (NULL, ?, ?, ?);")
	if err != nil {
		return err
	}

	_, err = stmt.Exec(
		time.Now().Format("2006-01-02"),
		d.Roll_string,
		strconv.Itoa(d.Total),
	)
	if err != nil {
		return err
	}

	return nil
}
