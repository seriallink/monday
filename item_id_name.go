package monday

type ItemBasic struct {
	Id           string        `json:"id"            graphql:"id"`
	Name         string        `json:"name"          graphql:"name"`
	ColumnValues []ColumnValue `json:"column_values" graphql:"column_values"`
}
