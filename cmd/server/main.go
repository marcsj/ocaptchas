package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/marcsj/ocaptchas/repo"
	"log"
)

func main() {
	//TODO: change where this database is located by an arg
	db, err := gorm.Open("sqlite3", "stored.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = repo.NewSessionRepo(db)
	if err != nil {
		log.Fatal(err)
	}
}