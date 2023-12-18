package monday

type ColumnType string

func (*ColumnType) GetGraphQLType() string {
	return "ColumnType"
}

func ColumnType2Pnt(value ColumnType) *ColumnType {
	return &value
}

const (
	ColumnTypeAutoNumber    ColumnType = "auto_number"
	ColumnTypeBoardRelation ColumnType = "board_relation"
	ColumnTypeButton        ColumnType = "button"
	ColumnTypeCheckbox      ColumnType = "checkbox"
	ColumnTypeColorPicker   ColumnType = "color_picker"
	ColumnTypeCountry       ColumnType = "country"
	ColumnTypeCreationLog   ColumnType = "creation_log"
	ColumnTypeDate          ColumnType = "date"
	ColumnTypeDependency    ColumnType = "dependency"
	ColumnTypeDoc           ColumnType = "doc"
	ColumnTypeDropdown      ColumnType = "dropdown"
	ColumnTypeEmail         ColumnType = "email"
	ColumnTypeFile          ColumnType = "file"
	ColumnTypeFormula       ColumnType = "formula"
	ColumnTypeHour          ColumnType = "hour"
	ColumnTypeItemAssignees ColumnType = "item_assignees"
	ColumnTypeId            ColumnType = "id"
	ColumnTypeLastUpdated   ColumnType = "last_updated"
	ColumnTypeLink          ColumnType = "link"
	ColumnTypeLocation      ColumnType = "location"
	ColumnTypeLongText      ColumnType = "long_text"
	ColumnTypeMirror        ColumnType = "mirror"
	ColumnTypeName          ColumnType = "name"
	ColumnTypeNumbers       ColumnType = "numbers"
	ColumnTypePeople        ColumnType = "people"
	ColumnTypePhone         ColumnType = "phone"
	ColumnTypeProgress      ColumnType = "progress"
	ColumnTypeRating        ColumnType = "rating"
	ColumnTypeStatus        ColumnType = "status"
	ColumnTypeSubtasks      ColumnType = "subtasks"
	ColumnTypeTags          ColumnType = "tags"
	ColumnTypeTeam          ColumnType = "team"
	ColumnTypeText          ColumnType = "text"
	ColumnTypeTimeline      ColumnType = "timeline"
	ColumnTypeTimeTracking  ColumnType = "time_tracking"
	ColumnTypeVote          ColumnType = "vote"
	ColumnTypeWeek          ColumnType = "week"
	ColumnTypeWorldClock    ColumnType = "world_clock"
	ColumnTypeUnsupported   ColumnType = "unsupported"
)
