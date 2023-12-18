package monday

import "encoding/json"

/*
Field	Description
column Column!	The column the value belongs to.
id ID!	The column's unique identifier.
text String	The text representation of the column's value. Not every column supports the text value. Please note that the complexity of this field for mirror, link, and dependency columns is increased in API version 2023-10 and later.
type ColumnType!	The column's type.
value JSON	The column's raw value.
*/
type ColumnValue struct {
	Column Column          `json:"column,omitempty" graphql:"column"`
	Id     string          `json:"id,omitempty"     graphql:"id"`
	Text   string          `json:"text,omitempty"   graphql:"text"`
	Type   ColumnType      `json:"type,omitempty"   graphql:"type"`
	Value  json.RawMessage `json:"value,omitempty"  graphql:"value"`
}
