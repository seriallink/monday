package monday

import "time"

/*
Field	Description
assets [Asset]	The update's assets/files.
body String!	The update's HTML-formatted body.
created_at Date	The update's creation date.
creator User	The update's creator.
creator_id String	The unique identifier of the update's creator.
id ID!	The update's unique identifier.
item_id String	The update's item ID.
replies [Reply]	The update's replies.
text_body String	The update's text body.
updated_at Date	The update's last edit date.
*/

type Update struct {
	Assets    []Asset   `json:"assets"     graphql:"assets"`
	Body      string    `json:"body"       graphql:"body"`
	CreatedAt time.Time `json:"created_at" graphql:"created_at"`
	Creator   *User     `json:"creator"    graphql:"creator"`
	CreatorId string    `json:"creator_id" graphql:"creator_id"`
	Id        string    `json:"id"         graphql:"id"`
	ItemId    string    `json:"item_id"    graphql:"item_id"`
	Replies   []Reply   `json:"replies"    graphql:"replies"`
	TextBody  string    `json:"text_body"  graphql:"text_body"`
	UpdatedAt time.Time `json:"updated_at" graphql:"updated_at"`
}
