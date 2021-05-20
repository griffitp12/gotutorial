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

func NewDino(w http.ResponseWriter, r *http.Request) {
	DB, err = gorm.Open(postgres.Open(DSN), &gorm.Config{})
    if err != nil {
        panic("could not connect to the database")
    }

    vars := mux.Vars(r)
    name := vars["name"]
    food := vars["food"]

    DB.Create(&Dino{Name: name, Food: food})

    fmt.Fprintf(w, "New Dino successfully created")
}

func DeleteDino(w http.ResponseWriter, r *http.Request) {
	DB, err = gorm.Open(postgres.Open(DSN), &gorm.Config{})
    if err != nil {
        panic("could not connect to the database")
    }

    vars := mux.Vars(r)
    name := vars["name"]

    var dino Dino
    DB.Where("name = ?", name).Find(&dino)
    DB.Delete(&dino)

    fmt.Fprintf(w, "%v successfully deleted", dino.Name)
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

