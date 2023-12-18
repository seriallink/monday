package monday

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const token = ""

var client *Client

func init() {
	client = NewClient(token, ApiVersion202310)
}

// test for GetWorkspaces
func TestGetWorkspaces(t *testing.T) {
	response, err := client.GetWorkspaces(&WorkspaceArguments{
		//Ids:  &[]ID{""},
		Kind: WorkspaceKind2Pnt(WorkspaceKindOpen),
	})
	assert.NoError(t, err)
	assert.NotEmpty(t, response)
}

// test for GetBoards
func TestGetBoards(t *testing.T) {
	response, err := client.GetBoards(&BoardArguments{
		//Ids:          &[]ID{""},
		//WorkspaceIds: &[]ID{""},
		//Page:         Int2Pnt(1),
	})
	assert.NoError(t, err)
	assert.NotEmpty(t, response)
}

// test for GetPlan
func TestGetPlan(t *testing.T) {
	response, err := client.GetPlan()
	assert.NoError(t, err)
	assert.NotEmpty(t, response)
}

// test for GetAccount
func TestGetAccount(t *testing.T) {
	response, err := client.GetAccount()
	assert.NoError(t, err)
	assert.NotEmpty(t, response)
}

// test for GetItems
func TestGetItems(t *testing.T) {
	response, err := client.GetItems(nil)
	assert.NoError(t, err)
	assert.NotEmpty(t, response)
}

func TestGetBoardItems(t *testing.T) {
	response, err := client.GetBoardItems("", &ItemsPageArguments{
		Limit: Int2Pnt(50),
	})
	assert.NoError(t, err)
	assert.NotEmpty(t, response)
}

func TestGetSubItems(t *testing.T) {
	response, err := client.GetSubItems(&ItemsArguments{
		Ids: &[]ID{""},
	})
	assert.NoError(t, err)
	assert.NotEmpty(t, response)
}
