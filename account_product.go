package monday

type AccountProduct struct {
	Id   ID     `json:"id"   graphql:"id"`
	Kind string `json:"kind" graphql:"kind"`
}
