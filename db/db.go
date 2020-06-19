package db

import (
	"fmt"
	"github.com/ofili/recharge/handlers"
	"log"

	"github.com/jinzhu/gorm"
)

// Initial DB setup
var db *gorm.DB

// InitDB ...InitDB
func InitDB() {
	var err error
	db, err = gorm.Open("mysql", "root:P@ssw0rd)@tcp(localhost:3306)/recharge_db?parseTime=True")

	if err != nil {
		fmt.Println(err)
		panic("failed to connect to database")
	}
	log.Println("connection established")

	// Create table in database
	db.HasTable(&handlers.Person{})
	db.DropTableIfExists(&handlers.Person{})

	// Migrateion to create table for recharge_db
	db.AutoMigrate(&handlers.Person{})
}