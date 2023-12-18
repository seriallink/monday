package monday

import "encoding/json"

type Column struct {
	Archived    bool            `json:"archived"     graphql:"archived"`
	Description string          `json:"description"  graphql:"description"`
	Id          string          `json:"id"           graphql:"id"`
	SettingsStr json.RawMessage `json:"settings_str" graphql:"settings_str"`
	Title       string          `json:"title"        graphql:"title"`
	Type        ColumnType      `json:"type"         graphql:"type"`
	Width       int             `json:"width"        graphql:"width"`
}
