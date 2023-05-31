// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type ObjectInput struct {
	CursorField   *string `json:"cursor_field,omitempty"`
	DestinationID int64   `json:"destination_id"`
	DisplayName   string  `json:"display_name"`
	// This is where Fabra will insert the End Customer ID specified when creating a source.
	EndCustomerIDField string         `json:"end_customer_id_field"`
	Frequency          int64          `json:"frequency"`
	FrequencyUnits     FrequencyUnits `json:"frequency_units"`
	Namespace          string         `json:"namespace"`
	ObjectFields       []ObjectField  `json:"object_fields,omitempty"`
	PrimaryKey         *string        `json:"primary_key,omitempty"`
	TableName          string         `json:"table_name"`
}
