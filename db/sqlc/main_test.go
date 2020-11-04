package db

import (
	"os"
	"testing"
	"log"
	_ "github.com/lib/pq"
	"database/sql"
	
)

var testQueries *Queries

var (
	dbDriver = os.Getenv("DB_DRIVER")
	dbSource = os.Getenv("DB_SOURCE")
)



func TestMain(m *testing.M) {

	conn, err := sql.Open(dbDriver, dbSource)

	if err != nil {
		log.Fatal("could not connect to the database", err)
	}

	testQueries = New(conn)

	os.Exit(m.Run())
}