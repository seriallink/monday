package monday

type Group struct {
	Archived bool   `json:"archived" graphql:"archived"`
	Color    string `json:"color"    graphql:"color"`
	Deleted  bool   `json:"deleted"  graphql:"deleted"`
	Id       string `json:"id"       graphql:"id"`
	//Items    []Item `json:"items"    graphql:"items"`
	//ItemsResponse *ItemsResponse `json:"items_page" graphql:"items_page"`
	Position string `json:"position" graphql:"position"`
	Title    string `json:"title"    graphql:"title"`
}
