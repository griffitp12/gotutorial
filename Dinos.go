package main

import (
	/* "encoding/json" */
	"fmt"
	/* "github.com/gorilla/mux" */
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	/* "io/ioutil" */
	/* "net/http" */
)


var DB *gorm.DB
var err error
var DSN = "host=localhost user=postgres password=shinmone dbname=dinos port=5432 sslmode=disable TimeZone=Asia/Shanghai"

type Dino struct {
    gorm.Model
    Name string
    Food string
}

func InitialMigration() {
    DB, err := gorm.Open(postgres.Open(DSN), &gorm.Config{})
    if err != nil {
        fmt.Println(err.Error())
        panic("Failed to connect to database")
    } 
   /*  defer db.Close() */

    DB.AutoMigrate(&Dino{})
}


