package notify

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os"

	"github.com/sambasivareddy-ch/db_backup_utility/pkg/utils"
)

func SendNotificationOnDiscord(message string) error {
	discordWebHook := os.Getenv("DISCORD_WEBHOOK_URI")
	payload := utils.DiscordPayload{
		Content: message,
	}

	payloadBytes, _ := json.Marshal(payload)

	resp, err := http.Post(discordWebHook, "application/json", bytes.NewBuffer(payloadBytes))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}
