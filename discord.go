package gologgerdiscord

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

func sendMessage(channelID string, message *Message, authConfig *AuthConfig) error {
	var client = &http.Client{}
	var content interface{}

	data, err := json.Marshal(message)
	if err != nil {
		log.Println(err)
		return err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf(SEND_MESSAGE_CHANNEL_URL, channelID), strings.NewReader(string(data)))
	if err != nil {
		log.Println(err)
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("%s %s", authConfig.Type, authConfig.Token))

	res, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return err
	}

	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&content)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
