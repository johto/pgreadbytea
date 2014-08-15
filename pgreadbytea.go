package main

import (
	_ "github.com/lib/pq"
	"database/sql"
	"log"
	"os"
)

func connectToPostgres() *sql.DB {
	var conninfo string
	if os.Getenv("PGSSLMODE") == "" {
		conninfo = "sslmode=disable"
	}
	dbh, err := sql.Open("postgres", conninfo)
	if err != nil {
		log.Fatal(err)
	}
	return dbh
}

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("usage: %s query", os.Args[0])
	}
	dbh := connectToPostgres()
	defer dbh.Close()

	var data []byte

	err := dbh.QueryRow(os.Args[1]).Scan(&data)
	if err != nil {
		log.Fatalf("SQL query failed: %s", err)
	}

	_, err = os.Stdout.Write(data)
	if err != nil {
		log.Fatalf("could not write output to stdout: %s", err)
	}
}
