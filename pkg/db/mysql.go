package db

import (
	"database/sql"
	"fmt"

	"github.com/sambasivareddy-ch/db_backup_utility/context"
	"github.com/sambasivareddy-ch/db_backup_utility/pkg/logging"

	_ "github.com/go-sql-driver/mysql"
)

func PingSQL(params *context.DBSessionContext) error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		params.DBUsername,
		params.DBPassword,
		params.DBHost,
		params.DBPort,
		params.DBName,
	)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		logging.Logger.LogError("Unable to connect to MySQL")
		return err
	}

	return db.Ping()
}
