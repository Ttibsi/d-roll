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
	values []int
	Roll_string string
}

func (d Dice) addToString(x int) {
	if d.Roll_string != "" {
		d.Roll_string += "+"
	}
	d.Roll_string += strconv.Itoa(x)
	log.Println(d.Roll_string)
}

func (d Dice) homeHandler(w http.ResponseWriter, r *http.Request) {
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

func (d Dice) homePostHandler(w http.ResponseWriter, r *http.Request) {
	pressed_btn, _ := strconv.Atoi(r.FormValue("value"))

	switch pressed_btn {
	case 20:
		log.Println("accepted 20")
		d.values = append(d.values, 20)
		d.addToString(20)
	case 12:
		log.Println("accepted 12")
		d.values = append(d.values, 12)
		d.addToString(12)
	case 10:
		log.Print("accepted 10\n")
		d.values = append(d.values, 10)
		d.addToString(10)
	case 8:
		log.Print("accepted 8\n")
		d.values = append(d.values, 8)
		d.addToString(8)
	case 4:
		log.Print("accepted 4\n")
		d.values = append(d.values, 4)
		d.addToString(4)
	case 100:
		log.Print("accepted 100\n")
		d.values = append(d.values, 100)
		d.addToString(100)
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
