package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/marcsj/ocaptchas/repo"
	"github.com/namsral/flag"
	"log"
)

var (
	databaseFile = flag.String(
		"database_file", "stored.db", "location for database file")
)
func main() {
	db, err := gorm.Open("sqlite3", *databaseFile)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = repo.NewSessionRepo(db)
	if err != nil {
		log.Fatal(err)
	}
}