package db

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	localDBSource = "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"
	ciDBSource    = "postgresql://root:secret@postgres:5432/simple_bank?sslmode=disable"
)

var testQueries *Queries
var testDB *pgxpool.Pool

func TestMain(m *testing.M) {
	var err error
	dbSource := localDBSource
	if os.Getenv("GITHUB_ACTIONS") == "true" {
		dbSource = ciDBSource
	}

	config, err := pgxpool.ParseConfig(dbSource)
	if err != nil {
		log.Fatal("Cannot parse config: ", err)
	}

	config.ConnConfig.DefaultQueryExecMode = pgx.QueryExecModeSimpleProtocol

	testDB, err = pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		log.Fatal("Cannot connect to db: ", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}
