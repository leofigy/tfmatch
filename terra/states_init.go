package terra

import (
	"sort"

	uuid "github.com/satori/go.uuid"
)

//NewTerraState
func NewStateV4() *StateV4 {
	state := &StateV4{}
	state.init()
	return state
}

//init functions
func (s *StateV4) init() {
	if s.Version == 0 {
		s.Version = stateVersion
	}

	if len(s.TerraformVersion) == 0 {
		s.TerraformVersion = TFVersion
	}

	if len(s.Lineage) == 0 {
		s.Lineage = uuid.NewV4().String()
	}
}

// Grab from terraform state
func (s *StateV4) normalize() {
	sort.Stable(SortableResourcesV4(s.Resources))
	for x := range s.Resources {
		sort.Stable(SortableInstancesV4(s.Resources[x].Instances))
	}
}
