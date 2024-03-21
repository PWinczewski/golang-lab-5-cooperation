package main

import (
	"log"
	"os"

	"github.com/gocarina/gocsv"
)

// Definiujemy strukturę danych, które chcemy zapisać do pliku CSV
type Record struct {
	Name  string `csv:"Name"`
	Age   int    `csv:"Age"`
	Email string `csv:"Email"`
}

func main() {
	// Przykładowe dane
	records := []*Record{
		{Name: "John Doe", Age: 30, Email: "john@example.com"},
		{Name: "Jane Smith", Age: 25, Email: "jane@example.com"},
		{Name: "Bob Johnson", Age: 35, Email: "bob@example.com"},
	}

	// Otwieramy plik CSV do zapisu
	file, err := os.Create("output.csv")
	if err != nil {
		log.Fatalf("error creating CSV file: %s", err)
	}
	defer file.Close()

	// Zapisujemy dane do pliku
	err = gocsv.MarshalFile(&records, file)
	if err != nil {
		log.Fatalf("error writing CSV data: %s", err)
	}

}
