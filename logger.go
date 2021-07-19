package gologgerdiscord

func Init(authConfig AuthConfig, channelID string) *Logger {
	return &Logger{
		Auth: AuthConfig{
			Token: authConfig.Token,
			Type:  authConfig.Type,
		},
		ChannelID: channelID,
	}
}

func (l Logger) Success(message, title, description, URL string) error {
	var embeds []*EmbedObject
	embeds = append(embeds, &EmbedObject{
		Title:       title,
		Description: description,
		URL:         URL,
		Color:       0x5CBD33,
	})

	err := sendMessage(l.ChannelID, &Message{
		Content: message,
		Embeds:  embeds,
	}, &AuthConfig{
		Token: l.Auth.Token,
		Type:  l.Auth.Type,
	})

	return err
}

func (l Logger) Info(message, title, description, URL string) error {
	var embeds []*EmbedObject
	embeds = append(embeds, &EmbedObject{
		Title:       title,
		Description: description,
		URL:         URL,
		Color:       0x555DFA,
	})

	err := sendMessage(l.ChannelID, &Message{
		Content: message,
		Embeds:  embeds,
	}, &AuthConfig{
		Token: l.Auth.Token,
		Type:  l.Auth.Type,
	})

	return err
}

func (l Logger) Warn(message, title, description, URL string) error {
	var embeds []*EmbedObject
	embeds = append(embeds, &EmbedObject{
		Title:       title,
		Description: description,
		URL:         URL,
		Color:       0xFAC655,
	})

	err := sendMessage(l.ChannelID, &Message{
		Content: message,
		Embeds:  embeds,
	}, &AuthConfig{
		Token: l.Auth.Token,
		Type:  l.Auth.Type,
	})

	return err
}

func (l Logger) Fatal(message, title, description, URL string) error {
	var embeds []*EmbedObject
	embeds = append(embeds, &EmbedObject{
		Title:       title,
		Description: description,
		URL:         URL,
		Color:       0xFF0000,
	})

	err := sendMessage(l.ChannelID, &Message{
		Content: message,
		Embeds:  embeds,
	}, &AuthConfig{
		Token: l.Auth.Token,
		Type:  l.Auth.Type,
	})

	return err
}

func (l Logger) Error(message, title, description, URL string) error {
	var embeds []*EmbedObject
	embeds = append(embeds, &EmbedObject{
		Title:       title,
		Description: description,
		URL:         URL,
		Color:       0xFF0000,
	})

	err := sendMessage(l.ChannelID, &Message{
		Content: message,
		Embeds:  embeds,
	}, &AuthConfig{
		Token: l.Auth.Token,
		Type:  l.Auth.Type,
	})

	return err
}
