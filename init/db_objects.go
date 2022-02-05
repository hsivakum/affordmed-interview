package init

import (
	"affordmed/dbhelper"
	"database/sql"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

func Db() *sqlx.DB {
	connectionString := dbhelper.BuildConnectionString()
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		logrus.Panic("Could not connect to affordmed DB", err)
	}
	return sqlx.NewDb(db, "postgres")
}
