package monday

type ItemsQueryOperator string

func (*ItemsQueryOperator) GetGraphQLType() string {
	return "ItemsQueryOperator"
}

func ItemsQueryOperator2Pnt(value ItemsQueryOperator) *ItemsQueryOperator {
	return &value
}

const (
	ItemsQueryOperatorAnd ItemsQueryOperator = "and"
	ItemsQueryOperatorOr  ItemsQueryOperator = "or"
)
