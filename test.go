package main

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"os"
	"reflect"
	"strings"
	"testing"
)

func TestConvertCSVtoJSONLines(t *testing.T) {
	// Create a temporary CSV file for testing
	csvData := `value,income,age,rooms,bedrooms,pop,hh
	"100000",50000,25,4,2,1000,2
	"150000",60000,30,5,3,1200,3
	"200000",70000,35,6,4,1500,4
	`
	csvFile, err := os.CreateTemp("", "test.csv")
	if err != nil {
		t.Fatalf("Error creating temporary CSV file: %v", err)
	}
	defer os.Remove(csvFile.Name())

	_, err = csvFile.WriteString(csvData)
	if err != nil {
		t.Fatalf("Error writing to temporary CSV file: %v", err)
	}
	csvFile.Close()

	// Create a buffer to store the output JSON lines
	var outputBuffer bytes.Buffer

	// Call the function with the temporary CSV file and output buffer
	err = convertCSVtoJSONLines(csvFile.Name(), &outputBuffer)
	if err != nil {
		t.Fatalf("Error converting CSV to JSON lines: %v", err)
	}

	// Decode each line of JSON and compare with expected results
	expected := []House{
		{"100000", 50000, 25, 4, 2, 1000, 2},
		{"150000", 60000, 30, 5, 3, 1200, 3},
		{"200000", 70000, 35, 6, 4, 1500, 4},
	}

	decoder := json.NewDecoder(&outputBuffer)
	for _, exp := range expected {
		var house House
		err := decoder.Decode(&house)
		if err != nil {
			t.Fatalf("Error decoding JSON line: %v", err)
		}

		if !reflect.DeepEqual(house, exp) {
			t.Errorf("Expected %v, got %v", exp, house)
		}
	}

	// Check if all JSON lines have been read
	var house House
	err = decoder.Decode(&house)
	if err == nil {
		t.Error("Expected end of JSON lines, got more data")
	}
	if err.Error() != "EOF" {
		t.Errorf("Expected EOF, got %v", err)
	}
}
