package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/NovikovAndrew/SimpleBank/util"

	_ "github.com/lib/pq"
)

var testQuery *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	config, err := util.LoadConfig("../..")
	if err != nil {
		log.Fatal("connot to load config, error: ", err)
	}

	testDB, err := sql.Open(config.DBDriver, config.DBSource)

	if err != nil {
		log.Fatal("can not connect to db, error: ", err)
	}

	testQuery = New(testDB)

	os.Exit(m.Run())
}
