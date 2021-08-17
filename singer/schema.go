package singer

// Schema defines the shape and properties of a stream
type Schema struct {
	Properties map[string]SchemaProperties `json:"properties,omitempty"`
	Type       string                      `json:"type,omitempty"`

	/*
		AdditionalProperties bool              `json:"additionalProperties,omitempty"`
		KeyProperties        []string          `json:"keyProperties,omitempty"`
		Selected             bool              `json:"selected,omitempty"`
		Description          string            `json:"description,omitempty"`
		ExclusiveMinimum     float32           `json:"exclusiveMinimum,omitempty"`
		ExclusiveMaximum     float32           `json:"exclusiveMaximum,omitempty"`
		MultipleOf           float32           `json:"multipleOf,omitempty"`
		MaxLength            int               `json:"maxLength,omitempty"`
		MinLength            int               `json:"minLength,omitempty"`
		Anyof                string            `json:"anyOf,omitempty"`
		PatternProperties    string            `json:"patternProperties,omitempty"`
	*/
}

type SchemaProperties struct {
	Inclusion string   `json:"inclusion,omitempty"`
	Format    string   `json:"format,omitempty"`
	Minimum   float32  `json:"minimum,omitempty"`
	Maximum   float32  `json:"maximum,omitempty"`
	Type      []string `json:"type,omitempty"`
}
