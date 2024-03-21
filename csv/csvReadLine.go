// w Golangu do czytania pliku csv mozemy uzyc package encoding/csv
// mozemy wczytac plik csv caly od razu (csv.Reader.ReadAll()) badz linijka po linijce (csv.Reader.Read() )
// minus z ta paczka jest taki ze zwraca nam zawsze typ danych string, mozemu uzyc jeszcze paczki GoCsv

// przyklad linjka po linijce
// zwraca [][]string, gdzie kazdy element to linijka z pliku csv
package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
    // otwieramy plik csv
    f, err := os.Open("data.csv")
    if err != nil {
        log.Fatal(err)
    }

    // zamykamy plik na koniec programu
    defer f.Close()

    // czytamy wartosci uzywajac csv.Reader
    csvReader := csv.NewReader(f)
    for {
        rec, err := csvReader.Read()
        if err == io.EOF {
            break
        }
        if err != nil {
            log.Fatal(err)
        }

        fmt.Printf("%+v\n", rec)
    }
}

// Output:
// [warzywo owoc]
// [marchew banan]
// [brokuł truskawka]

