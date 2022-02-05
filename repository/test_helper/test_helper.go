package test_helper

import (
	"affordmed/dbhelper"
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"testing"
)

func GetDB() *sqlx.DB {
	if err := godotenv.Load("/Users/hariharans/IdeaProjects/affordmed/.envrc"); err != nil {
		logrus.Info("error loading env config")
	}
	connectionString := dbhelper.BuildConnectionString()
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		logrus.Panic("Could not connect to affordmed DB", err)
	}
	return sqlx.NewDb(db, "postgres")
}

func ClearDB(db *sqlx.DB, t *testing.T) {
	tableOrder := []string{"users"}
	for _, table := range tableOrder {
		_, err := db.Exec(fmt.Sprintf("delete from %v", table))
		if err != nil {
			t.Errorf("unable to cleanup db %v", err)
		}
	}
}
