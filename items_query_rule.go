package monday

type ItemsQueryRule struct {
	ColumnID         string                 `json:"column_id,omitempty"         graphql:"column_id"`
	CompareAttribute string                 `json:"compare_attribute,omitempty" graphql:"compare_attribute"`
	CompareValue     string                 `json:"compare_value,omitempty"     graphql:"compare_value"`
	Operator         ItemsQueryRuleOperator `json:"operator,omitempty"          graphql:"operator"`
}
