package models

type EmbedFooter struct {
	Text         string `json:"text"`
	IconURL      string `json:"icon_url"`
	ProxyIconURL string `json:"proxy_icon_url"`
}

type EmbedThumbnail struct {
	URL    string `json:"url"`
	Height int    `json:"height"`
	Width  int    `json:"width"`
}

type EmbedField struct {
	Name   string `json:"name"`
	Value  string `json:"value"`
	Inline bool   `json:"inline"`
}

type Embed struct {
	Color       string          `json:"color,omitempty"`
	Title       string          `json:"title"`
	Description string          `json:"description"`
	URL         string          `json:"url"`
	Footer      *EmbedFooter    `json:"footer"`
	Thumbnail   *EmbedThumbnail `json:"thumbnail"`
	Fields      []*EmbedField   `json:"fields"`
}

type WebhookBuilder struct {
	Embeds []*Embed `json:"embeds"`
}
