package shared

type Object struct {
	DestinationID      *int64        `json:"destination_id,omitempty"`
	DisplayName        *string       `json:"display_name,omitempty"`
	EndCustomerIDField *string       `json:"end_customer_id_field,omitempty"`
	ID                 *int64        `json:"id,omitempty"`
	Namespace          *string       `json:"namespace,omitempty"`
	ObjectFields       []ObjectField `json:"object_fields,omitempty"`
	TableName          *string       `json:"table_name,omitempty"`
}
