package shared

type ObjectInput struct {
	DestinationID      int64         `json:"destination_id"`
	DisplayName        string        `json:"display_name"`
	EndCustomerIDField string        `json:"end_customer_id_field"`
	Namespace          string        `json:"namespace"`
	ObjectFields       []ObjectField `json:"object_fields,omitempty"`
	TableName          string        `json:"table_name"`
}
