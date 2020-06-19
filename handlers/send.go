package handlers

import (
	"recharge/db"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/jinzhu/gorm"
)



// Send airtime
func Send(w http.ResponseWriter, r *http.Request)  {
	var person Person
	json.NewEncoder(r.Body).Decode(&person)
	//Create new interns by inserting records in the interns table
	db.Create(&person)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(person)
}