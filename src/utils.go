package src

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func rollDice(d *Dice) {
	var results []int
	var total int
    rand.Seed(time.Now().UnixNano())

	for _, val := range d.Values {
		results = append(results, rand.Intn(val) + 1)
	}

	d.Roll_string = strings.Trim(fmt.Sprint(results), "[]")

	for _, val := range results {
		log.Printf("Val: %v, total: %v", val, total)
		total += val
	}

	d.Roll_string += "<br>" //This works in html instead of a \n
	d.Roll_string += "Result: "
	d.Roll_string += strconv.Itoa(total)
}

func insertToDB() {}
