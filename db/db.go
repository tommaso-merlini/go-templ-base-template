package db

import (
	"database/sql"
	"os"

	_ "github.com/lib/pq"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/sqlitedialect"
)

var DB *bun.DB

func Init() error {
	connection := os.Getenv("SUPABASE_CONNECTION_STRING_TEST")
	if os.Getenv("ENV") == "production" {
		connection = os.Getenv("SUPABASE_CONNECTION_STRING_PROD")
	}

	sqldb, err := sql.Open("postgres", connection)
	if err != nil {
		panic(err)
	}
	err = sqldb.Ping()
	if err != nil {
		panic(err)
	}
	db := bun.NewDB(sqldb, sqlitedialect.New())
	DB = db
	return nil
}
