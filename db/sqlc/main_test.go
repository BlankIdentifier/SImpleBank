package db

import (
	"os"
	"testing"
	"log"
	_ "github.com/lib/pq"
	"database/sql"
	
)

var testQueries *Queries
var testDB *sql.DB

// var (
// 	dbDriver = os.Getenv("DB_DRIVER")
// 	dbSource = os.Getenv("DB_SOURCE")
// )

const (
	dbDriver = "postgres"
	dbSource = "postgresql://sb:secret1@localhost:5432/simplebank?sslmode=disable"
)



func TestMain(m *testing.M) {

	var err error 
	testDB, err = sql.Open(dbDriver, dbSource)

	if err != nil {
		log.Fatal("could not connect to the database", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}