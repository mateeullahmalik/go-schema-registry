package srclient

// Schema defines schema response structure
type Schema struct {
	Subject    string      `json:"subject"`
	Version    int         `json:"version"`
	ID         int         `json:"id"`
	Type       string      `json:"SchemaType"`
	References []Reference `json:"references"`
	SchemaStr  string      `json:"schema"`
}

// Reference defines schema references
type Reference struct {
	Name    string `json:"name"`
	Subject string `json:"subject"`
	Version int    `json:"version"`
}
