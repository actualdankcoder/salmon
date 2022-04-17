package main

import (
	"encoding/json"
	"errors"
	"log"
	"os"
)

func LoadJSON() SalDat {
	// Check if data.json exists
	if _, err := os.Stat("data.json"); errors.Is(err, os.ErrNotExist) {

		// Create data.json if not exists
		rJS, err := json.MarshalIndent(SalDat{}, "", "    ")
		if err != nil {
			panic(err)
		}
		os.WriteFile("data.json", rJS, 0644)

	} else if err != nil {

		log.Fatalf("Unknown Error: %s", err)

	}

	// Read JSON
	file, err1 := os.ReadFile("data.json")
	if err1 != nil {
		log.Fatalf("Read Error: %s", err1)
	}

	// Map JSON to Struct
	data := SalDat{}
	err := json.Unmarshal([]byte(file), &data)
	if err != nil {
		log.Printf("Unmarshal Error: %s", err)
	}
	return data
}

func WriteJSON(data SalDat) SalDat {
	// Struct to JSON Map
	respJSON, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		log.Printf("Marshal Error: %s", err)
	}

	// Write to File
	err = os.WriteFile("data.json", respJSON, 0644)
	if err != nil {
		log.Printf("Write Error: %s", err)
	}
	return data
}
