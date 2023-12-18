package monday

type BoardsOrderBy string

func (*BoardsOrderBy) GetGraphQLType() string {
	return "BoardsOrderBy"
}

func BoardsOrderBy2Pnt(value BoardsOrderBy) *BoardsOrderBy {
	return &value
}

const (
	BoardsOrderByCreatedAt BoardsOrderBy = "created_at"
	BoardsOrderByUsedAt    BoardsOrderBy = "used_at"
)
