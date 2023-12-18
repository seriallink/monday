package monday

type Plan struct {
	MaxUsers int    `json:"max_users" graphql:"max_users"`
	Period   string `json:"period"    graphql:"period"`
	Tier     string `json:"tier"      graphql:"tier"`
	Version  int    `json:"version"   graphql:"version"`
}

// GetPlan returns the Plan for the account.
func (c *Client) GetPlan() (Plan, error) {
	var response struct {
		Account struct {
			Plan Plan `json:"plan" graphql:"plan"`
		} `json:"account" graphql:"account"`
	}
	err := c.RunQuery(&response, nil)
	return response.Account.Plan, err
}
