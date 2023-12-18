package monday

import "time"

type Reply struct {
	Body      string    `json:"body"       graphql:"body"`
	CreatedAt time.Time `json:"created_at" graphql:"created_at"`
	Creator   *User     `json:"creator"    graphql:"creator"`
	CreatorId string    `json:"creator_id" graphql:"creator_id"`
	Id        string    `json:"id"         graphql:"id"`
	TextBody  string    `json:"text_body"  graphql:"text_body"`
	UpdatedAt time.Time `json:"updated_at" graphql:"updated_at"`
}
