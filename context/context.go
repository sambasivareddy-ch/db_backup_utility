package context

import (
	"encoding/json"
	"os"

	"github.com/sambasivareddy-ch/db_backup_utility/pkg/logging"
)

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

var GlobalSessionCtx *DBSessionContext = &DBSessionContext{}

func (dbCtx *DBSessionContext) SaveToFile() {
	_, err := os.OpenFile("db_context.json", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)
	if err != nil {
		logging.Logger.LogError("error occurred while creating session file")
	}

	data, _ := json.Marshal(dbCtx)
	os.WriteFile("db_context.json", data, 0644)
}

func LoadSession() error {
	data, err := os.ReadFile("db_context.json")
	if err != nil {
		return err
	}

	json.Unmarshal(data, &GlobalSessionCtx)
	logging.Logger.LogInfo("Loaded Session Details")

	return nil
}
