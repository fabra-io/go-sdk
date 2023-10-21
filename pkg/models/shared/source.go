// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type Source struct {
	Connection    *Connection `json:"connection,omitempty"`
	DisplayName   *string     `json:"display_name,omitempty"`
	EndCustomerID *string     `json:"end_customer_id,omitempty"`
	ID            *int64      `json:"id,omitempty"`
}

func (o *Source) GetConnection() *Connection {
	if o == nil {
		return nil
	}
	return o.Connection
}

func (o *Source) GetDisplayName() *string {
	if o == nil {
		return nil
	}
	return o.DisplayName
}

func (o *Source) GetEndCustomerID() *string {
	if o == nil {
		return nil
	}
	return o.EndCustomerID
}

func (o *Source) GetID() *int64 {
	if o == nil {
		return nil
	}
	return o.ID
}
