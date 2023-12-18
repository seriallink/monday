package monday

type ItemsResponse struct {
	Cursor string      `json:"cursor" graphql:"cursor"`
	Items  []ItemBasic `json:"items"  graphql:"items"`
}

type ItemsPageArguments struct {
	Cursor     *string     `structs:"cursor"       graphql:"cursor"`
	Limit      *int        `structs:"limit"        graphql:"limit"`
	QueryParam *ItemsQuery `structs:"query_params" graphql:"query_params"`
}
