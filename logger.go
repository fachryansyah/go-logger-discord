package gologgerdiscord

import "fmt"

type Logger struct {
	Auth      AuthConfig
	ChannelID string
}

func Init(authConfig AuthConfig, channelID string) (*Logger, error) {
	return &Logger{
		Auth: AuthConfig{
			Token: authConfig.Token,
			Type:  authConfig.Type,
		},
		ChannelID: channelID,
	}, nil
}

func (l Logger) Info(msg, env string) {
	fmt.Println("info", msg)
}

func (l Logger) Warn(msg, env string) {
	fmt.Println("warn", msg)
}

func (l Logger) Fatal(msg, env string) {
	fmt.Println("fatal", msg)
}

func (l Logger) Error(message, title, description, URL, env string) {
	var embeds []*EmbedObject
	embeds = append(embeds, &EmbedObject{
		Title:       title,
		Description: description,
		URL:         URL,
		Color:       0xFF0000,
	})

	sendMessage(l.ChannelID, &Message{
		Content: message,
		Embeds:  embeds,
	})
}
