package monday

type ItemsQueryRuleOperator string

func (*ItemsQueryRuleOperator) GetGraphQLType() string {
	return "ItemsQueryRuleOperator"
}

func ItemsQueryRuleOperator2Pnt(value ItemsQueryRuleOperator) *ItemsQueryRuleOperator {
	return &value
}

const (
	ItemsQueryRuleOperatorAnyOf               ItemsQueryRuleOperator = "any_of"
	ItemsQueryRuleOperatorNotAnyOf            ItemsQueryRuleOperator = "not_any_of"
	ItemsQueryRuleOperatorIsEmpty             ItemsQueryRuleOperator = "is_empty"
	ItemsQueryRuleOperatorIsNotEmpty          ItemsQueryRuleOperator = "is_not_empty"
	ItemsQueryRuleOperatorGreaterThan         ItemsQueryRuleOperator = "greater_than"
	ItemsQueryRuleOperatorGreaterThanOrEquals ItemsQueryRuleOperator = "greater_than_or_equals"
	ItemsQueryRuleOperatorLowerThan           ItemsQueryRuleOperator = "lower_than"
	ItemsQueryRuleOperatorLowerThanOrEquals   ItemsQueryRuleOperator = "lower_than_or_equal"
	ItemsQueryRuleOperatorBetween             ItemsQueryRuleOperator = "between"
	ItemsQueryRuleOperatorNotContainsText     ItemsQueryRuleOperator = "not_contains_text"
	ItemsQueryRuleOperatorContainsText        ItemsQueryRuleOperator = "contains_text"
	ItemsQueryRuleOperatorContainsTerms       ItemsQueryRuleOperator = "contains_terms"
	ItemsQueryRuleOperatorStartsWith          ItemsQueryRuleOperator = "starts_with"
	ItemsQueryRuleOperatorEndsWith            ItemsQueryRuleOperator = "ends_with"
	ItemsQueryRuleOperatorWithinTheNext       ItemsQueryRuleOperator = "within_the_next"
	ItemsQueryRuleOperatorWithinTheLast       ItemsQueryRuleOperator = "within_the_last"
)
