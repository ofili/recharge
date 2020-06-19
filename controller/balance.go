package controller

import (
	"encoding/json"
	"net/http"
)

type Airtime struct {
	Amount		int		`json:"airtime"`
}

func Balance(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "application/json")

	m := Airtime{15000}
	b, err := json.Marshal(m)
}