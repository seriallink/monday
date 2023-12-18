package monday

type ItemsQuery struct {
	Ids      []string           `json:"ids,omitempty"      graphql:"ids"`
	Rules    []ItemsQueryRule   `json:"rules,omitempty"    graphql:"rules"`
	Operator ItemsQueryOperator `json:"operator,omitempty" graphql:"operator"`
	OrderBy  []ItemsQueryOrder  `json:"order_by,omitempty" graphql:"order_by"`
}
