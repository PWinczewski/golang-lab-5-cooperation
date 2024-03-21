# GOlang ~ .INI files and you

## INI Syntax

INI files can have sections, which contain key-value pairs. The syntax for an INI file is as follows:

lab5.ini file:
```ini

[section]
key = value
;This is how you comment!

[section.subsection]
key2 = value2

[database]
user = dbuser
password = dbpassword
dbname = dbname

[server]
port = 8080
host = localhost

```

## Installing Golang INI
Packages have to be installed within a module.
To install the Golang INI package, run the following commands in the terminal:
```bash
go mod init module_name
go get gopkg.in/ini.v1
```

## Reading INI Files

The ini package provides a Load function that takes a file path and returns a *ini.File object. This object contains all the sections and key-value pairs in the INI file.

```go
package main

import (
    "fmt"
    "os"
    "gopkg.in/ini.v1"
)

func main() {
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
}
```
returns:
```
dbname
dbuser
dbpassword
8080
```

## Writing INI Files

The ini package provides a SaveTo function that takes a file path and writes the *ini.File object to the file.

```go
func main() {
//... previous code ...

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

    err = inidata_new.SaveTo("config2.ini")
    if err != nil {
        panic(err)
    }
}
```

config.ini:
```ini
[database]
host = localhost

[database.options]
option1 = false
```

same principle can be used to update existing *.ini files:
```go
func main() {
//... previous code ...
    sec = inidata_new.Section("database")

    sec.Key("host").SetValue("127.0.0.0")
    sec.Key("port").SetValue("3306")

    err = inidata_new.SaveTo("config.ini")
    if err != nil {
        panic(err)
    }
}
```

config.ini:
```ini
[database]
host = 127.0.0.0
port = 3306

[database.options]
option1 = false
```