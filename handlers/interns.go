package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/jinzhu/gorm"
)

// Person represents the model for an intern
// Default table name will be `persons`
type Person struct {
	gorm.Model

	ID			int		`json:"id"`
	firstName	string	`json:"f_name"`
	lastName	string	`json:"l_name"`
	Number		int		`json:"number"`
}

var persons []Person


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
	db.HasTable(&Person{})
	db.DropTableIfExists(&Person{})

	// Migration to create table for recharge_db
	db.AutoMigrate(&Person{})
}

// Get all interns
func Intern(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var persons []Person
	db.Find(&persons)
	json.NewEncoder(w).Encode(persons)
}
