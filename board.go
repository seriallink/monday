package monday

import (
	"encoding/json"
	"time"
)

type Board struct {
	//ActivityLogs    []ActivityLogType `json:"activity_logs"    graphql:"activity_logs"`
	BoardFolderId   ID              `json:"board_folder_id"  graphql:"board_folder_id"`
	BoardKind       BoardKind       `json:"board_kind"       graphql:"board_kind"`
	Columns         []Column        `json:"columns"          graphql:"columns"`
	Communication   json.RawMessage `json:"communication"    graphql:"communication"`
	Creator         *User           `json:"creator"          graphql:"creator"`
	Description     string          `json:"description"      graphql:"description"`
	Groups          []Group         `json:"groups"           graphql:"groups"`
	Id              ID              `json:"id"               graphql:"id"`
	ItemTerminology string          `json:"item_terminology" graphql:"item_terminology"`
	ItemsCount      int             `json:"items_count"      graphql:"items_count"`
	ItemsPage       *ItemsResponse  `json:"items_page"       graphql:"items_page"`
	Name            string          `json:"name"             graphql:"name"`
	Owner           *User           `json:"owner"            graphql:"owner"`
	Owners          []User          `json:"owners"           graphql:"owners"`
	Permissions     string          `json:"permissions"      graphql:"permissions"`
	State           string          `json:"state"            graphql:"state"`
	Subscribers     []User          `json:"subscribers"      graphql:"subscribers"`
	Tags            []Tag           `json:"tags"             graphql:"tags"`
	TopGroup        *Group          `json:"top_group"        graphql:"top_group"`
	Type            string          `json:"type"             graphql:"type"`
	UpdatedAt       time.Time       `json:"updated_at"       graphql:"updated_at"`
	Updates         []Update        `json:"updates"          graphql:"updates"`
	Views           []BoardView     `json:"views"            graphql:"views"`
	Workspace       *Workspace      `json:"workspace"        graphql:"workspace"`
	WorkspaceId     ID              `json:"workspace_id"     graphql:"workspace_id"`
}

type BoardArguments struct {
	BoardKind    *BoardKind     `structs:"board_kind"    graphql:"board_kind"`
	Ids          *[]ID          `structs:"ids"           graphql:"ids"`
	Limit        *int           `structs:"limit"         graphql:"limit"`
	OrderBy      *BoardsOrderBy `structs:"order_by"      graphql:"order_by"`
	Page         *int           `structs:"page"          graphql:"page"`
	State        *State         `structs:"state"         graphql:"state"`
	WorkspaceIds *[]ID          `structs:"workspace_ids" graphql:"workspace_ids"`
}

// GetBoards returns []Board for all boards.
func (c *Client) GetBoards(arguments *BoardArguments) ([]Board, error) {
	var response struct {
		Boards []Board `json:"boards" graphql:"boards(ids: $ids, board_kind: $board_kind, limit: $limit, order_by: $order_by, page: $page, state: $state, workspace_ids: $workspace_ids)"`
	}
	err := c.RunQuery(&response, ArgumentsToMap(arguments))
	return response.Boards, err
}
