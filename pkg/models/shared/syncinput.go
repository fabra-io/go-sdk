package shared

type SyncInput struct {
	CursorField    *string            `json:"cursor_field,omitempty"`
	CustomJoin     *string            `json:"custom_join,omitempty"`
	DestinationID  int64              `json:"destination_id"`
	DisplayName    string             `json:"display_name"`
	FieldMappings  []FieldMapping     `json:"field_mappings"`
	Frequency      int64              `json:"frequency"`
	FrequencyUnits FrequencyUnitsEnum `json:"frequency_units"`
	Namespace      *string            `json:"namespace,omitempty"`
	ObjectID       int64              `json:"object_id"`
	PrimaryKey     *string            `json:"primary_key,omitempty"`
	SourceID       int64              `json:"source_id"`
	TableName      *string            `json:"table_name,omitempty"`
}
