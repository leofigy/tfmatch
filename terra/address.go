package terra

type TerraAddress struct {
	Path   []string
	Index  int    `json:"index_key"`
	Name   string `json:"name"`
	Type   string `json:"type"`
	Module string `json:"module" optional:"yes"` // not all resources required module
	Mode   ResourceMode
}

func (addr *TerraAddress) GetModule() {

}
