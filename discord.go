package gologgerdiscord

import (
	"fmt"
	"io"
	"net/http"
)

func sendMessage(channelID string, message *Message) {
	var body io.Reader
	http.Post(fmt.Sprint(SEND_MESSAGE_CHANNEL_URL, channelID), "application/json", body)
	fmt.Println(body)
}
