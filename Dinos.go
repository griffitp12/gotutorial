package main

import (
	/* "encoding/json" */
	"fmt"
	/* "github.com/gorilla/mux" */
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
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

/* type DinoDetails struct {
    Name   string
    Food string
} */

type AllDinoPageData struct {
    PageTitle string
    Dinos     []Dino
}

func InitialMigration() {
    DB, err := gorm.Open(postgres.Open(DSN), &gorm.Config{
        Logger: logger.Default.LogMode(logger.Info),
    })
    if err != nil {
        fmt.Println(err.Error())
        panic("Failed to connect to database")
    } 
   /*  defer db.Close() */

    DB.AutoMigrate(&Dino{})
}


