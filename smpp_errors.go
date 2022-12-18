// Copyright (c) 2021 Ameed Jamous - TelecomXChange LLC.
// This program is licensed under the MIT License and is available for distribution and use.
// You are allowed to modify, distribute, and use this program as per the terms of the MIT License.
// Please contact the copyright holder (Ameed Jamous - a.jamous@telecomsxchange.com) if you have any questions or require further information.

package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Open a connection to the database
	db, err := sql.Open("mysql", "user:pass@tcp(127.0.0.1:3306)/tcxc")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create the HTTP handler function that selects all records.

	http.HandleFunc("/errorcodes", func(w http.ResponseWriter, r *http.Request) {

		// Log request

		log.Println("Received request:", r.Method, r.URL, r.RemoteAddr)

		// Execute the SELECT query to retrieve all rows from the smpp_error_codes table
		rows, err := db.Query("SELECT * FROM smpp_error_codes")
		if err != nil {
			log.Println(err)
			http.Error(w, http.StatusText(500), 500)
			return
		}
		defer rows.Close()

		// Iterate over the rows and scan the values into a slice of structs
		var errorCodes []struct {
			ID                int    `json:"id"`
			CommandStatusName string `json:"command_status_name"`
			Value             string `json:"value"`
			Description       string `json:"description"`
		}
		for rows.Next() {
			var ec struct {
				ID                int    `json:"id"`
				CommandStatusName string `json:"command_status_name"`
				Value             string `json:"value"`
				Description       string `json:"description"`
			}
			err := rows.Scan(&ec.ID, &ec.CommandStatusName, &ec.Value, &ec.Description)
			if err != nil {
				log.Println(err)
				http.Error(w, http.StatusText(500), 500)
				return
			}
			errorCodes = append(errorCodes, ec)
		}

		// Return the error codes as a JSON response
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(errorCodes)
	})

	// Start the server
	log.Fatal(http.ListenAndServe(":8083", nil))
	log.Println("Server is running")

}
