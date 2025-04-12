package context

import (
	"encoding/json"
	"os"

	"github.com/sambasivareddy-ch/db_backup_utility/pkg/logging"
)

// DBSessionContext holds the database connection details
// and other related information for the backup and restore operations.
// This structure is used to save the session context
// in a JSON file (db_context.json) for later use.
type DBSessionContext struct {
	DBType      string
	DBHost      string
	DBPort      string
	DBUsername  string
	DBPassword  string
	DBName      string
	BackupDir   string
	RestoreFile string // Used in restore case
}

/*
	GlobalSessionCtx is a global variable that holds the current session context.
	This variable is used to store the database connection details
	and other related information for the backup and restore operations.
*/
var GlobalSessionCtx *DBSessionContext = &DBSessionContext{}

/*
	SaveToFile method is used to save the current session context
	to a JSON file (db_context.json).
	Helps to persist the database connection details and other related information
	for backup and restore operations.
	This method creates a file called db_context.json
*/
func (dbCtx *DBSessionContext) SaveToFile() {
	_, err := os.OpenFile("db_context.json", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)
	if err != nil {
		logging.Logger.LogError("error occurred while creating session file")
	}

	data, _ := json.Marshal(dbCtx)
	os.WriteFile("db_context.json", data, 0644)
}

/*
	LoadSession method is used to load the session context
	from a JSON file (db_context.json).
	This method reads the file and unmarshals the JSON data to GlobalSessionCtx.
	It helps to restore the database connection details and other related information
	for backup and restore operations.
	This method is called when the user runs the backup or restore command.
*/
func LoadSession() error {
	data, err := os.ReadFile("db_context.json")
	if err != nil {
		return err
	}

	json.Unmarshal(data, &GlobalSessionCtx)
	logging.Logger.LogInfo("Loaded Session Details")

	return nil
}
