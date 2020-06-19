package controller

import (
	"encoding/json"
	_ "fmt"

	_ "github.com/jinzhu/gorm"

	_ "log"
	"net/http"

	_ "github.com/jinzhu/gorm"
)

// Send airtime
func Send(w http.ResponseWriter, r *http.Request)  {
	var intern Person
	json.NewEncoder(r.Body).Decode(&intern)
	//Create new interns by inserting records in the interns table
	db.Create(&intern)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(intern)
}