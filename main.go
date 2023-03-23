package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "file:data.db")

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Initialize the database
	repo := PayoutRepo{db}
	err = repo.InitPayoutTable()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("initialized")

	// Load all payouts from the file and inserts them
	payouts, err := LoadPayoutsFromFile("payouts.json")
	if err != nil {
		log.Fatal(err)
	}

	for _, payout := range payouts {
		err = repo.InsertPayout(payout)
	}
	if err != nil {
		log.Fatal(err)
	}

	// Get all payouts from the db and display them in console
	payouts, err = repo.GetAllPayouts()
	if err != nil {
		log.Fatal(err)
	}
	for _, payout := range payouts {
		fmt.Printf("%v\n", payout)
	}
}

func LoadPayoutsFromFile(filename string) ([]Payout, error) {
	// Open our jsonFile
	jsonFile, err := os.Open(filename)
	// if we os.Open returns an error then handle it
	if err != nil {
		return nil, err
	}
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}

	var result []Payout
	err = json.Unmarshal(byteValue, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
