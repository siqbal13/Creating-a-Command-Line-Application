package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

// House represents a single house data
type House struct {
	Value     string  `json:"value"`
	Income    float64 `json:"income"`
	Age       int     `json:"age"`
	Rooms     int     `json:"rooms"`
	Bedrooms  int     `json:"bedrooms"`
	Pop       int     `json:"pop"`
	Hh        int     `json:"hh"`
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: csvtojl input.csv output.jl")
		os.Exit(1)
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	// Open CSV file
	csvFile, err := os.Open(inputFile)
	if err != nil {
		log.Fatalf("Error opening CSV file: %v", err)
	}
	defer csvFile.Close()

	// Create JSON lines file
	jsonFile, err := os.Create(outputFile)
	if err != nil {
		log.Fatalf("Error creating JSON lines file: %v", err)
	}
	defer jsonFile.Close()

	// Create CSV reader
	reader := csv.NewReader(csvFile)

	// Read CSV headers
	headers, err := reader.Read()
	if err != nil {
		log.Fatalf("Error reading CSV headers: %v", err)
	}

	// Read CSV records and convert to JSON lines
	encoder := json.NewEncoder(jsonFile)
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error reading CSV record: %v", err)
		}

		house := House{}
		for i, value := range record {
			switch headers[i] {
			case "value":
				house.Value = value
			case "income":
				house.Income = parseFloat(value)
			case "age":
				house.Age = parseInt(value)
			case "rooms":
				house.Rooms = parseInt(value)
			case "bedrooms":
				house.Bedrooms = parseInt(value)
			case "pop":
				house.Pop = parseInt(value)
			case "hh":
				house.Hh = parseInt(value)
			}
		}

		err = encoder.Encode(house)
		if err != nil {
			log.Fatalf("Error encoding JSON line: %v", err)
		}
	}

	fmt.Println("Conversion completed successfully")
}

func parseFloat(s string) float64 {
	var f float64
	_, err := fmt.Sscanf(s, "%f", &f)
	if err != nil {
		log.Fatalf("Error parsing float: %v", err)
	}
	return f
}

func parseInt(s string) int {
	var i int
	_, err := fmt.Sscanf(s, "%d", &i)
	if err != nil {
		log.Fatalf("Error parsing int: %v", err)
	}
	return i
}
