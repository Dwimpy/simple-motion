package tests

import (
	"database/sql"
	db "github.com/Dwimpy/simple-motion/db/sqlc"
	_ "github.com/lib/pq"
	"log"
	"os"
	"testing"
)

var testQuery *db.Queries

func TestMain(m *testing.M) {
	conn, err := sql.Open("postgres", "postgresql://andreirobu:reload@localhost:5432/test_db?sslmode=disable")
	if err != nil {
		log.Fatal("Could not connect to database")
	}

	testQuery = db.New(conn)
	os.Exit(m.Run())
}
