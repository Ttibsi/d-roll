package src

import (
	"log"
	"net/http"
	"strconv"
	"text/template"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Dice struct {
	Values      []int
	Rolls       []int
	Total       int
	Display     string
	Roll_string string
	DbData      *[]DbRow
	Toggle      bool
}

func (d *Dice) addToString(x int) {
	if d.Roll_string != "" {
		d.Roll_string += "+"
	}
	d.Roll_string += strconv.Itoa(x)
}

func (d *Dice) reset(reset bool) {
	d.Values = d.Values[:0]
	d.Rolls = d.Rolls[:0]
	d.Roll_string = ""
	d.Toggle = false

	if reset {
		d.Display = ""
	}
}

func (d *Dice) homeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("www/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, d)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if d.Toggle {
		d.reset(true)
	}
}

func (d *Dice) homePostHandler(w http.ResponseWriter, r *http.Request) {
	pressed_btn, _ := strconv.Atoi(r.FormValue("value"))

	switch pressed_btn {
	case 20:
		d.Values = append(d.Values, 20)
		d.addToString(20)
		d.Display = d.Roll_string
	case 12:
		d.Values = append(d.Values, 12)
		d.addToString(12)
		d.Display = d.Roll_string
	case 10:
		d.Values = append(d.Values, 10)
		d.addToString(10)
		d.Display = d.Roll_string
	case 8:
		d.Values = append(d.Values, 8)
		d.addToString(8)
		d.Display = d.Roll_string
	case 4:
		d.Values = append(d.Values, 4)
		d.addToString(4)
		d.Display = d.Roll_string
	case 100:
		d.Values = append(d.Values, 100)
		d.addToString(100)
		d.Display = d.Roll_string
	case -1: // Roll button pressed
		rollDice(d)
		err := insertToDB(d)
		if err != nil {
			log.Println(err.Error())
		}
		d.reset(false)
	case -2: // Reset
		d.reset(true)
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func (d *Dice) resultsHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	d.DbData, err = GetData()
	if err != nil {
		return
	}

	tmpl, err := template.ParseFiles("www/results.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, d)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func Serve() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	d := Dice{Toggle: false}

	r.Get("/", d.homeHandler)
	r.Post("/", d.homePostHandler)
	r.Get("/results", d.resultsHandler)
	log.Fatal(http.ListenAndServe(":3000", r))
}
