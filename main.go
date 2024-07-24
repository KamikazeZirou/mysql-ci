package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	user := flag.String("u", "root", "the user to connect mysql server")
	password := flag.String("pass", "", "the user to connect mysql server")
	port := flag.Int("port", 3306, "the port of mysql server")
	host := flag.String("h", "127.0.0.1", "the address of mysql server")
	query := flag.String("q", "", "the sql query")
	file := flag.String("f", "", "the sql script file path")
	flag.Parse()

	//
	// Connect DB
	//
	db, err := sql.Open("mysql", fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/?parseTime=true&loc=Local&multiStatements=true",
		*user,
		*password,
		*host,
		*port,
	))
	if err != nil {
		log.Fatal("error connecting to mysql:" + err.Error())
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal("error pinging mysql:" + err.Error())
	}

	//
	// Read query
	//
	if len(*file) > 0 {
		fi, err := os.Open(*file)
		if err != nil {
			log.Fatal("error opening file:" + err.Error())
		}
		defer fi.Close()

		data, err := io.ReadAll(fi)
		if err != nil {
			log.Fatal("error reading file:" + err.Error())
		}
		q := string(data)
		query = &q
	}
	if len(*query) == 0 {
		log.Fatal("query or script-file must be set")
	}

	//
	// Run query
	//
	_, err = db.Exec(*query)
	if err != nil {
		log.Fatal("error executing query:" + err.Error())
	}
}
