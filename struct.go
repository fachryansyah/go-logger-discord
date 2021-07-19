package gologgerdiscord

type Logger struct {
	Auth      AuthConfig
	ChannelID string
}

type AuthConfig struct {
	Token string
	Type  string // Bot token or OAuth2
}

type Message struct {
	Content string         `json:"content"`
	Embeds  []*EmbedObject `json:"embeds"`
}

type EmbedObject struct {
	Title       string `json:"title"`
	Type        string `json:"type"`
	Description string `json:"description"`
	URL         string `json:"url"`
	Color       int    `json:"color"`
}
