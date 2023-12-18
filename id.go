package monday

import "encoding/json"

type ID json.Number

func (*ID) GetGraphQLType() string {
	return "ID"
}
