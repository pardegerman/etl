package singer

type Metadata struct {
	Breadcrumb    []string            `json:"breadcrumb"`
	MetadataProps *MetadataProperties `json:"metadata"`
}

type MetadataProperties struct {
	SelectedByDefault  *bool    `json:"selected-by-default,omitempty"`
	DatabaseName       string   `json:"database-name,omitempty"`
	RowCount           *int     `json:"row-count,omitempty"`
	IsView             *bool    `json:"is-view,omitempty"`
	TableKeyProperties []string `json:"table-key-properties,omitempty"`
	SqlDataType        string   `json:"sql-datatype,omitempty"`
	Selected           *bool    `json:"selected,omitempty"`
	ReplicationMethod  string   `json:"replication-method,omitempty"`
	ReplicationKey     string   `json:"replication-key,omitempty"`

	/*
		Inclusion               string   `json:"inclusion"`
		StreamName              string   `json:"stream_name"`
	*/
}
