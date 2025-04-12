package notify

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os"

	"github.com/sambasivareddy-ch/db_backup_utility/pkg/utils"
)

/*
	SendNotificationOnDiscord function sends a notification message to a Discord channel
	via a webhook URL.
	@params: message - The message to be sent to the Discord channel
	@returns: error - Returns an error if the message could not be sent
	@description: This function constructs a JSON payload with the message and sends it
	to the Discord webhook URL using an HTTP POST request.
	It uses the os package to get the webhook URL from an environment variable
*/
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
