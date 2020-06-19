package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"recharge/handlers"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/ofili/recharge/db"
)

func main() {

	url := "https://sandbox.wallets.africa/bills/airtime/providers/"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("x-rapidapi-host", "santamca-mobile-carrier-info-and-recharge-plans-v1.p.rapidapi.com")
	req.Header.Add("x-rapidapi-key", "SIGN-UP-FOR-KEY")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("error occured")
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

	//Router
	router := mux.NewRouter()

	// Send airtime
	router.HandleFunc("/send", handlers.Send).Methods("POST")

	// Get balance
	router.HandleFunc("/balance", handlers.Balance).Methods("GET")

	// Get interns
	router.HandleFunc("/Intern", handlers.Intern).Methods("GET")

	db.InitDB()

	log.Fatal(http.ListenAndServe(":9000", router))

}
