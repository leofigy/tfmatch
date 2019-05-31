package terra

const stateVersion = 4
const TFVersion = "0.12.0"

/* Items in the main terrafor state file */
type ValType struct {
	Value string `json:"value"`
	Type  string `json:"type"`
}

// Global state
type TerraState struct {
	Version   int    `json:"version"`
	TFVersion string `json:"terraform_version"`
	Serial    int64  `json:"serial"`
	Lineage   string `json:"lineage"`

	Outputs map[string]ValType `json:"outputs"`
	// currently checking
	Resources []*map[string]ResourceState `json:"resources"`
}

type ResourceState struct {
	Name      string       `json:"name"`
	Type      string       `json:"type"`
	Mode      ResourceMode `json:"mode"`           // by now all are managed
	Each      EachMode     `json:"each,omitempty"` // to do the list
	Provider  string       `json:"provider"`
	Instances []InstanceState
}

type InstanceState struct {
	SchemaVersion int      `json:"schema_version"`
	IndexKey      int      `json:"index_key, omitempty"`
	Dependencies  []string `json:"depends_on,omitempty"`
	// ToDo to match with the fields in ptkill
	Attributes map[string]interface{} `json:"attributes"`
}
