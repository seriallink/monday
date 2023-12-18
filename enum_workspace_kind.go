package monday

type WorkspaceKind string

func (*WorkspaceKind) GetGraphQLType() string {
	return "WorkspaceKind"
}

func WorkspaceKind2Pnt(value WorkspaceKind) *WorkspaceKind {
	return &value
}

const (
	WorkspaceKindOpen   WorkspaceKind = "open"
	WorkspaceKindClosed WorkspaceKind = "closed"
)
