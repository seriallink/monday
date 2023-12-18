package monday

type WorkspacesOrderBy string

func (*WorkspacesOrderBy) GetGraphQLType() string {
	return "WorkspacesOrderBy"
}

func WorkspacesOrderBy2Pnt(value WorkspacesOrderBy) *WorkspacesOrderBy {
	return &value
}

const (
	WorkspacesOrderByCreatedAt WorkspacesOrderBy = "created_at"
)
