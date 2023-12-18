package monday

import "time"

type User struct {
	Id                 ID        `json:"id"                   graphql:"id"`
	Name               string    `json:"name"                 graphql:"name"`
	Email              string    `json:"email"                graphql:"email"`
	Enabled            bool      `json:"enabled"              graphql:"enabled"`
	IsAdmin            bool      `json:"is_admin"             graphql:"is_admin"`
	IsGuest            bool      `json:"is_guest"             graphql:"is_guest"`
	IsPending          bool      `json:"is_pending"           graphql:"is_pending"`
	IsViewOnly         bool      `json:"is_view_only"         graphql:"is_view_only"`
	IsVerified         bool      `json:"is_verified"          graphql:"is_verified"`
	JoinDate           time.Time `json:"join_date"            graphql:"join_date"`
	LastActivity       time.Time `json:"last_activity"        graphql:"last_activity"`
	Location           string    `json:"location"             graphql:"location"`
	MobilePhone        string    `json:"mobile_phone"         graphql:"mobile_phone"`
	Phone              string    `json:"phone"                graphql:"phone"`
	PhotoOriginal      string    `json:"photo_original"       graphql:"photo_original"`
	PhotoSmall         string    `json:"photo_small"          graphql:"photo_small"`
	PhotoThumb         string    `json:"photo_thumb"          graphql:"photo_thumb"`
	PhotoThumbSmall    string    `json:"photo_thumb_small"    graphql:"photo_thumb_small"`
	PhotoTiny          string    `json:"photo_tiny"           graphql:"photo_tiny"`
	SignUpProductKind  string    `json:"sign_up_product_kind" graphql:"sign_up_product_kind"`
	TimeZoneIdentifier string    `json:"time_zone_identifier" graphql:"time_zone_identifier"`
	Title              string    `json:"title"                graphql:"title"`
	Url                string    `json:"url"                  graphql:"url"`
	UtcHoursDiff       int       `json:"utc_hours_diff"       graphql:"utc_hours_diff"`
}
