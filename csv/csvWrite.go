package main

import (
	"log"
	"os"

	"github.com/gocarina/gocsv"
)

// Define a struct to represent your data
type Record struct {
	Name  string `csv:"Name"`
	Age   int    `csv:"Age"`
	Email string `csv:"Email"`
}

func main() {
	// Create some sample data
	records := []*Record{
		{Name: "John Doe", Age: 30, Email: "john@example.com"},
		{Name: "Jane Smith", Age: 25, Email: "jane@example.com"},
		{Name: "Bob Johnson", Age: 35, Email: "bob@example.com"},
	}

	// Open a CSV file for writing
	file, err := os.Create("output.csv")
	if err != nil {
		log.Fatalf("error creating CSV file: %s", err)
	}
	defer file.Close()

	// Write records to the CSV file
	err = gocsv.MarshalFile(&records, file)
	if err != nil {
		log.Fatalf("error writing CSV data: %s", err)
	}

}
