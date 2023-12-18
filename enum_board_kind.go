package monday

type BoardKind string

func (*BoardKind) GetGraphQLType() string {
	return "BoardKind"
}

func BoardKind2Pnt(value BoardKind) *BoardKind {
	return &value
}

const (
	BoardKindPublic  BoardKind = "public"
	BoardKindPrivate BoardKind = "private"
	BoardKindShare   BoardKind = "share"
)
