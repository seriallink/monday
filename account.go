package monday

type Account struct {
	CountryCode          string           `json:"country_code"           graphql:"country_code"`
	FirstDayOfTheWeek    string           `json:"first_day_of_the_week"  graphql:"first_day_of_the_week"`
	Id                   ID               `json:"id"                     graphql:"id"`
	Logo                 string           `json:"logo"                   graphql:"logo"`
	Name                 string           `json:"name"                   graphql:"name"`
	Plan                 *Plan            `json:"plan"                   graphql:"plan"`
	Products             []AccountProduct `json:"products"               graphql:"products"`
	ShowTimelineWeekends bool             `json:"show_timeline_weekends" graphql:"show_timeline_weekends"`
	SignUpProductKind    string           `json:"sign_up_product_kind"   graphql:"sign_up_product_kind"`
	Slug                 string           `json:"slug"                   graphql:"slug"`
	Tier                 string           `json:"tier"                   graphql:"tier"`
}

// GetAccount returns the Account for current user.
func (c *Client) GetAccount() (Account, error) {
	var response struct {
		Me struct {
			Account Account `json:"account" graphql:"account"`
		} `json:"me" graphql:"me"`
	}
	err := c.RunQuery(&response, nil)
	return response.Me.Account, err
}
