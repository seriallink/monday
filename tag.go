package monday

type Tag struct {
	Color string `json:"color" graphql:"color"`
	Id    int    `json:"id"    graphql:"id"`
	Name  string `json:"name"  graphql:"name"`
}
