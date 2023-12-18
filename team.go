package monday

type Team struct {
	Id         int    `json:"id"          graphql:"id"`
	Name       string `json:"name"        graphql:"name"`
	PictureUrl string `json:"picture_url" graphql:"picture_url"`
	Users      []User `json:"users"       graphql:"users"`
}
