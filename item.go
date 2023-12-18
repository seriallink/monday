package monday

import (
	"time"
)

type Item struct {
	//AccountId int `json:"account_id" graphql:"account_id"`
	//Assets []Asset `json:"assets" graphql:"assets"`
	//Board        *Board        `json:"board"         graphql:"board"`
	ColumnValues []ColumnValue `json:"column_values" graphql:"column_values"`
	CreatedAt    time.Time     `json:"created_at"    graphql:"created_at"`
	Creator      *User         `json:"creator"       graphql:"creator"`
	CreatorId    string        `json:"creator_id"    graphql:"creator_id"`
	Email        string        `json:"email"         graphql:"email"`
	Group        *Group        `json:"group"         graphql:"group"`
	Id           string        `json:"id"            graphql:"id"`
	//LinkedItems     []Item        `json:"linked_items"      graphql:"linked_items"`
	Name string `json:"name" graphql:"name"`
	//ParentItem      *Item         `json:"parent_item"       graphql:"parent_item"`
	//RelativeLink    string        `json:"relative_link"     graphql:"relative_link"`
	//State           *State        `json:"state"             graphql:"state"`
	//Subitems []Item `json:"subitems" graphql:"subitems"`
	//Subscribers     []User        `json:"subscribers"       graphql:"subscribers"`
	UpdatedAt time.Time `json:"updated_at" graphql:"updated_at"`
	Updates   []Update  `json:"updates"           graphql:"updates"`
}

type ItemsArguments struct {
	ExcludeNonactive *bool `structs:"exclude_nonactive" graphql:"exclude_nonactive"`
	Ids              *[]ID `structs:"ids"               graphql:"ids"`
	Limit            *int  `structs:"limit"             graphql:"limit"`
	NewestFirst      *bool `structs:"newest_first"      graphql:"newest_first"`
	Page             *int  `structs:"page"              graphql:"page"`
}

// GetItems returns []item for all Items.
func (c *Client) GetItems(arguments *ItemsArguments) ([]Item, error) {
	var response struct {
		Items []Item `json:"items" graphql:"items(exclude_nonactive: $exclude_nonactive, ids: $ids, limit: $limit, newest_first: $newest_first, page: $page)"`
	}
	err := c.RunQuery(&response, ArgumentsToMap(arguments))
	return response.Items, err
}

// GetBoardItems returns all Items for a specific board.
func (c *Client) GetBoardItems(boardId ID, arguments *ItemsPageArguments) ([]Item, error) {
	var response struct {
		Boards []struct {
			ItemsPage struct {
				Items []Item `json:"items" graphql:"items"`
			} `json:"items_page" graphql:"items_page(cursor: $cursor, limit: $limit, query_params: $query_params)"`
		} `json:"boards" graphql:"boards(ids: $board_ids)"`
	}
	args := ArgumentsToMap(arguments)
	args["board_ids"] = []ID{boardId}
	err := c.RunQuery(&response, args)
	if len(response.Boards) == 0 {
		return nil, err
	}
	return response.Boards[0].ItemsPage.Items, err
}

// GetSubItems returns all sub items for a specific item.
func (c *Client) GetSubItems(arguments *ItemsArguments) ([]Item, error) {
	var response struct {
		Items []struct {
			Items []Item `json:"subitems" graphql:"subitems"`
		} `json:"items" graphql:"items(exclude_nonactive: $exclude_nonactive, ids: $ids, limit: $limit, newest_first: $newest_first, page: $page)"`
	}
	err := c.RunQuery(&response, ArgumentsToMap(arguments))
	if len(response.Items) == 0 {
		return nil, err
	}
	return response.Items[0].Items, err
}
