package db

import (
	"database/sql"
	"fmt"

	"github.com/sambasivareddy-ch/db_backup_utility/context"
	"github.com/sambasivareddy-ch/db_backup_utility/pkg/logging"

	_ "github.com/lib/pq"
)

/*
	Utility function to Ping the PostgreSQL Database
	@params: params - DBSessionContext
	This function is used to check the connection to the PostgreSQL database.
	It takes the database connection parameters from the context and attempts to connect to the database.
	If the connection is successful, it returns nil. Otherwise, it returns an error.
*/
func PingPG(params *context.DBSessionContext) error {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		params.DBHost,
		params.DBPort,
		params.DBUsername,
		params.DBPassword,
		params.DBName,
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		logging.Logger.LogError("Unable to connect to Postgres")
		return err
	}

	logging.Logger.LogInfo("Successfully Connected to Postgres")

	return db.Ping()
}
