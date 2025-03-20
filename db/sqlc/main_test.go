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
	DBSource = "postgresql://root:secret@127.0.0.1:5432/simple_bank?sslmode=disable"
)

var testQueries *Queries
var testDB *pgxpool.Pool

func TestMain(m *testing.M) {
	var err error

	dbSource := getDBSource()

	config, err := pgxpool.ParseConfig(dbSource)
	if err != nil {
		log.Fatal("Cannot parse config: ", err)
	}

	config.ConnConfig.DefaultQueryExecMode = pgx.QueryExecModeSimpleProtocol

	testDB, err = pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}

func getDBSource() string {
	if envSource := os.Getenv("DB_SOURCE"); envSource != "" {
		return envSource
	}
	return DBSource
}
