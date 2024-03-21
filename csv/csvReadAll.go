// przyklad czytania pliku csv caly od razu
// zwraca wycinek stringow czyli []string

package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func readCsvFile(filePath string) [][]string {
    f, err := os.Open(filePath)
    if err != nil {
        log.Fatal("Unable to read input file " + filePath, err)
    }
    defer f.Close()

    csvReader := csv.NewReader(f)
    records, err := csvReader.ReadAll()
    if err != nil {
        log.Fatal("Unable to parse file as CSV for " + filePath, err)
    }

    return records
}

func main() {
    records := readCsvFile("data.csv")
    fmt.Println(records)
}

// Output:
// [[warzywo owoc] [marchew banan] [brokuł truskawka]]