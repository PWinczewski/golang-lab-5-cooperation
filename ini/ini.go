package main

import (
    "fmt"
    "os"
    "gopkg.in/ini.v1"
)

func main() {
	// reading
    inidata, err := ini.Load("lab5.ini")
    if err != nil {
        fmt.Printf("Failed to read file: %v", err)
        os.Exit(1)
    }
    db_section := inidata.Section("database")
    server_section := inidata.Section("server")

    fmt.Println(db_section.Key("dbname").String())
    fmt.Println(db_section.Key("user").String())
    fmt.Println(db_section.Key("password").String())
	port, err := server_section.Key("port").Int()
	if err != nil {
        fmt.Printf("Failed to parse port: %v", err)
        os.Exit(1)
    }
    fmt.Println(port)

	//writing
    inidata_new := ini.Empty()
    sec, err := inidata_new.NewSection("database")
    if err != nil {
        panic(err)
    }
    _, err = sec.NewKey("host", "localhost")
    if err != nil {
        panic(err)
    }

    sec, err = inidata_new.NewSection("database.options")
    if err != nil {
        panic(err)
    }

    _, err = sec.NewKey("option1", "false")
    if err != nil {
        panic(err)
    }

    err = inidata_new.SaveTo("config.ini")
    if err != nil {
        panic(err)
    }

	//updating
	sec = inidata_new.Section("database")

    sec.Key("host").SetValue("127.0.0.0")
    sec.Key("port").SetValue("3306")

    err = inidata_new.SaveTo("config.ini")
    if err != nil {
        panic(err)
    }
}