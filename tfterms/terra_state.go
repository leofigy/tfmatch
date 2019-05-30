package tfterms

const version = 4
const TFVersion = "0.12"

// dummy struct xD we could use a map of maps
// but just to avoid confusion
type ValType struct {
	Value string `json:"value"`
	Type  string `json:"type"`
}

type TerraState struct {
	Version   int    `json:"version"`
	TFVersion string `json:"terraform_version, omitempty"`
	Serial    int64  `json:"serial"`
	Lineage   string `json:"lineage"`

	Outputs map[string]ValType `json:"outputs"`
	// currently checking
	Resources []*map[string]interface{} `json:"resources"`
}
