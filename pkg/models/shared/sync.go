package shared

type Sync struct {
	CursorField    *string             `json:"cursor_field,omitempty"`
	CustomJoin     *string             `json:"custom_join,omitempty"`
	DestinationID  *int64              `json:"destination_id,omitempty"`
	DisplayName    *string             `json:"display_name,omitempty"`
	FieldMappings  []FieldMapping      `json:"field_mappings,omitempty"`
	Frequency      *int64              `json:"frequency,omitempty"`
	FrequencyUnits *FrequencyUnitsEnum `json:"frequency_units,omitempty"`
	ID             *int64              `json:"id,omitempty"`
	Namespace      *string             `json:"namespace,omitempty"`
	ObjectID       *int64              `json:"object_id,omitempty"`
	PrimaryKey     *string             `json:"primary_key,omitempty"`
	SourceID       *int64              `json:"source_id,omitempty"`
	TableName      *string             `json:"table_name,omitempty"`
}
