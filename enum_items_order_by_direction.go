package monday

type ItemsOrderByDirection string

func (*ItemsOrderByDirection) GetGraphQLType() string {
	return "ItemsOrderByDirection"
}

func ItemsOrderByDirection2Pnt(value ItemsOrderByDirection) *ItemsOrderByDirection {
	return &value
}

const (
	ItemsOrderByDirectionAsc  ItemsOrderByDirection = "asc"
	ItemsOrderByDirectionDesc ItemsOrderByDirection = "desc"
)
