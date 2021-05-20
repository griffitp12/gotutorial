package main

import (
	/* "encoding/json" */
	"fmt"
	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	/* "io/ioutil" */
	"net/http"
    "html/template"
)

var tmpl = template.Must(template.ParseFiles("static/index.html"))


func AllDinos(w http.ResponseWriter, r *http.Request) {
	DB, err = gorm.Open(postgres.Open(DSN), &gorm.Config{})
    if err != nil {
        panic("could not connect to the database")
    }

    var dinos []Dino
    DB.Find(&dinos)
    /* json.NewEncoder(w).Encode(dinos) */
    data := AllDinoPageData{
		PageTitle: "My Dinos list",
		Dinos: dinos,
	}
	tmpl.Execute(w, data)
}

func AddDino(w http.ResponseWriter, r *http.Request) {
	DB, err = gorm.Open(postgres.Open(DSN), &gorm.Config{})
    if err != nil {
        panic("could not connect to the database")
    }
	
	tmpl := template.Must(template.ParseFiles("static/add_a_dino.html"))
	if r.Method != http.MethodPost {
		tmpl.Execute(w, nil)
		return
	}
	details := Dino{
		Name:   r.FormValue("name"),
		Food: r.FormValue("food"),	
	}

	DB.Create(&Dino{Name: details.Name, Food: details.Food})
	tmpl.Execute(w, struct{ Success bool }{true})
}

func DeleteDinoPageServe(w http.ResponseWriter, r *http.Request) {
	DB, err = gorm.Open(postgres.Open(DSN), &gorm.Config{})
    if err != nil {
        panic("could not connect to the database")
    }

    tmpl := template.Must(template.ParseFiles("static/delete_a_dino.html"))
	if r.Method != http.MethodPost {
		tmpl.Execute(w, nil)
		return
	}
	details := Dino{
		Name:   r.FormValue("Name"),
	}

    var dino Dino
    DB.Where("Name = ?", details.Name).Find(&dino)
    DB.Delete(&dino)

    tmpl.Execute(w, struct{ Success bool }{true})
}

func UpdateDino(w http.ResponseWriter, r *http.Request) {
	DB, err = gorm.Open(postgres.Open(DSN), &gorm.Config{})
    if err != nil {
        panic("could not connect to the database")
    }

    vars := mux.Vars(r)
    name := vars["name"]
    food := vars["food"]

    var dino Dino
    DB.Where("name = ?", name).Find(&dino)
    dino.Food = food

    DB.Save(&dino)

    fmt.Fprintf(w, "%v successfully updated", dino.Name)
}

