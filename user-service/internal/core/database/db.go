package database

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"os"
)

func SetUpConnection() *sqlx.DB {
	connectStr := os.Getenv("SQL_CONTAINER_DSN")
	driver := os.Getenv("SQL_CONTAINER_DRIVER")

	db, err := sqlx.Open(driver, connectStr)
	if err != nil {
		logrus.Error(err)
		panic(err)
	}

	return db
}
