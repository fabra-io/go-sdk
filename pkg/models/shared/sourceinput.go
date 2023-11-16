// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type SourceInput struct {
	BigqueryConfig  *BigQueryConfig  `json:"bigquery_config,omitempty"`
	ConnectionType  ConnectionType   `json:"connection_type"`
	DisplayName     string           `json:"display_name"`
	EndCustomerID   string           `json:"end_customer_id"`
	MongodbConfig   *MongoDbConfig   `json:"mongodb_config,omitempty"`
	RedshiftConfig  *RedshiftConfig  `json:"redshift_config,omitempty"`
	SnowflakeConfig *SnowflakeConfig `json:"snowflake_config,omitempty"`
}

func (o *SourceInput) GetBigqueryConfig() *BigQueryConfig {
	if o == nil {
		return nil
	}
	return o.BigqueryConfig
}

func (o *SourceInput) GetConnectionType() ConnectionType {
	if o == nil {
		return ConnectionType("")
	}
	return o.ConnectionType
}

func (o *SourceInput) GetDisplayName() string {
	if o == nil {
		return ""
	}
	return o.DisplayName
}

func (o *SourceInput) GetEndCustomerID() string {
	if o == nil {
		return ""
	}
	return o.EndCustomerID
}

func (o *SourceInput) GetMongodbConfig() *MongoDbConfig {
	if o == nil {
		return nil
	}
	return o.MongodbConfig
}

func (o *SourceInput) GetRedshiftConfig() *RedshiftConfig {
	if o == nil {
		return nil
	}
	return o.RedshiftConfig
}

func (o *SourceInput) GetSnowflakeConfig() *SnowflakeConfig {
	if o == nil {
		return nil
	}
	return o.SnowflakeConfig
}
