package src

import (
	"net/http"
	"strconv"
	"text/template"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Dice struct {
	Values []int
	Roll_string string
}

func (d *Dice) addToString(x int) {
	if d.Roll_string != "" {
		d.Roll_string += "+"
	}
	d.Roll_string += strconv.Itoa(x)
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
}

func (d *Dice) homePostHandler(w http.ResponseWriter, r *http.Request) {
	pressed_btn, _ := strconv.Atoi(r.FormValue("value"))

	switch pressed_btn {
	case 20:
		d.Values = append(d.Values, 20)
		d.addToString(20)
	case 12:
		d.Values = append(d.Values, 12)
		d.addToString(12)
	case 10:
		d.Values = append(d.Values, 10)
		d.addToString(10)
	case 8:
		d.Values = append(d.Values, 8)
		d.addToString(8)
	case 4:
		d.Values = append(d.Values, 4)
		d.addToString(4)
	case 100:
		d.Values = append(d.Values, 100)
		d.addToString(100)
	case 1: // Roll button pressed
		rollDice(d)
		insertToDB()
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func Serve() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	d := Dice{}

    r.Get("/", d.homeHandler)
    r.Post("/", d.homePostHandler)
    // r.Get("/results", resultsHandler)
    http.ListenAndServe(":3000", r)
}
