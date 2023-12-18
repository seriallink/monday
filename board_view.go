package monday

import "encoding/json"

type BoardView struct {
	Id                  ID              `json:"id"                     graphql:"id"`
	Name                string          `json:"name"                   graphql:"name"`
	SettingsStr         json.RawMessage `json:"settings_str"           graphql:"settings_str"`
	Type                string          `json:"type"                   graphql:"type"`
	ViewSpecificDataStr string          `json:"view_specific_data_str" graphql:"view_specific_data_str"`
}
