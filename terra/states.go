package terra

import "encoding/json"

//This structs are an excerpt of the V4
// state file generation from Terraform

const stateVersion = 4
const TFVersion = "0.12.0"

// V4 state file
type StateV4 struct {
	Version          int                      `json:"version"`
	TerraformVersion string                   `json:"terraform_version"`
	Serial           uint64                   `json:"serial"`
	Lineage          string                   `json:"lineage"`
	RootOutputs      map[string]OutputStateV4 `json:"outputs"`
	Resources        []ResourceStateV4        `json:"resources"`
}

type OutputStateV4 struct {
	ValueRaw     json.RawMessage `json:"value"`
	ValueTypeRaw json.RawMessage `json:"type"`
	Sensitive    bool            `json:"sensitive,omitempty"`
}

type ResourceStateV4 struct {
	Module         string                  `json:"module,omitempty"`
	Mode           string                  `json:"mode"`
	Type           string                  `json:"type"`
	Name           string                  `json:"name"`
	EachMode       string                  `json:"each,omitempty"`
	ProviderConfig string                  `json:"provider"`
	Instances      []InstanceObjectStateV4 `json:"instances"`
}

type InstanceObjectStateV4 struct {
	IndexKey interface{} `json:"index_key,omitempty"`
	Status   string      `json:"status,omitempty"`
	Deposed  string      `json:"deposed,omitempty"`

	SchemaVersion  uint64            `json:"schema_version"`
	AttributesRaw  json.RawMessage   `json:"attributes,omitempty"`
	AttributesFlat map[string]string `json:"attributes_flat,omitempty"`

	PrivateRaw []byte `json:"private,omitempty"`

	Dependencies []string `json:"depends_on,omitempty"`
}
