package monday

import (
	"time"
)

type Workspace struct {
	AccountProduct    *AccountProduct `json:"account_product"    graphql:"account_product"`
	CreatedAt         time.Time       `json:"created_at"         graphql:"created_at"`
	Description       string          `json:"description"        graphql:"description"`
	Id                ID              `json:"id"                 graphql:"id"`
	Kind              WorkspaceKind   `json:"kind"               graphql:"kind"`
	Name              string          `json:"name"               graphql:"name"`
	OwnersSubscribers []User          `json:"owners_subscribers" graphql:"owners_subscribers"`
	State             string          `json:"state"              graphql:"state"`
	TeamsSubscribers  []Team          `json:"teams_subscribers"  graphql:"teams_subscribers"`
	UsersSubscribers  []User          `json:"users_subscribers"  graphql:"users_subscribers"`
}

type WorkspaceArguments struct {
	Ids     *[]ID              `structs:"ids"      graphql:"ids"`
	Kind    *WorkspaceKind     `structs:"kind"     graphql:"kind"`
	Limit   *int               `structs:"limit"    graphql:"limit"`
	OrderBy *WorkspacesOrderBy `structs:"order_by" graphql:"order_by"`
	Page    *int               `structs:"page"     graphql:"page"`
	State   *State             `structs:"state"    graphql:"state"`
}

// GetWorkspaces returns []Workspace for all Workspaces.
func (c *Client) GetWorkspaces(arguments *WorkspaceArguments) ([]Workspace, error) {
	var response struct {
		Workspaces []Workspace `json:"workspaces" graphql:"workspaces(ids: $ids, kind: $kind, limit: $limit, order_by: $order_by, page: $page, state: $state)"`
	}
	err := c.RunQuery(&response, ArgumentsToMap(arguments))
	return response.Workspaces, err
}
