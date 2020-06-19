package handlers

import (
	"encoding/json"
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

// Get all interns
func Intern(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var persons []Person
	do.Find(&persons)
	json.NewEncoder(w).Encode(persons)
}
