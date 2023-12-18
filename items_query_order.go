package monday

type ItemsQueryOrder struct {
	ColumnID  string                `json:"column_id,omitempty" graphql:"column_id"`
	Direction ItemsOrderByDirection `json:"direction,omitempty" graphql:"direction"`
}
