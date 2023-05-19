// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type DestinationInput struct {
	BigqueryConfig  *BigQueryConfig  `json:"bigquery_config,omitempty"`
	ConnectionType  ConnectionType   `json:"connection_type"`
	DisplayName     string           `json:"display_name"`
	MongodbConfig   *MongoDbConfig   `json:"mongodb_config,omitempty"`
	RedshiftConfig  *RedshiftConfig  `json:"redshift_config,omitempty"`
	SnowflakeConfig *SnowflakeConfig `json:"snowflake_config,omitempty"`
}
