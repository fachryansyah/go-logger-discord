package gologgerdiscord

type AuthConfig struct {
	Token string
	Type  string // Bot token, OAuth2
}

type Message struct {
	Content string
	Embeds  []*EmbedObject
}

type EmbedObject struct {
	Title       string
	Type        string
	Description string
	URL         string
	Color       int
}
