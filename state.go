package monday

type State string

const (
	StateAll      State = "all"
	StateActive   State = "active"
	StateArchived State = "archived"
	StateDeleted  State = "deleted"
)
