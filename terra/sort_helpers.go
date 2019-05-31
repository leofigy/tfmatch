package terra

// helpers to sort the items in the file
// based on terraform no credits for this 

type SortableResourcesV4 []ResourceStateV4
type SortableInstancesV4 []InstanceObjectStateV4

func (list SortableResourcesV4) Len() int {
	return len(list)
}

func (list SortableResourcesV4) Swap(i, j int) {
	list[i], list[j] = list[j], list[i]
}

func (list SortableResourcesV4) Less(i, j int) bool {
	switch {
	// Sort by data/managed
	case list[i].Mode != list[j].Mode:
		return list[i].Mode < list[j].Mode
	case list[i].Type != list[j].Type:
		return list[i].Type < list[j].Type
	case list[i].Name < list[j].Name:
		return list[i].Name < list[j].Name
	default:
		return false
	}
}

///////////////////////////////////////////////////////////////////

func (list SortableInstancesV4) Len() int {
	return len(list)
}

func (list SortableInstancesV4) Swap(i, j int) {
	list[i], list[j] = list[j], list[i]
}

func (list SortableInstancesV4) Less(i, j int) bool {
	// instances indexes 
	Ai, Bi := list[i].IndexKey, list[j].IndexKey
	if Ai != Bi {
		if (Ai == nil) != (Bi == nil){
			return Ai == nil
		}
		// nums 
		if AiInt , isInt := Ai.(int); isInt {
			if BiInt, isInt := Bi.(int); isInt{
				return AiInt < BiInt
			}
			return true
		}
		// strs
		if AiStr, isStr := Ai.(string); isStr {
			if BiStr, isStr := Bi.(string); isStr {
				return AiStr < BiStr
			}
			return true
		}
	}

	if list[i].Deposed != list[j].Deposed {
		return list[i].Deposed < list[j].Deposed
	}

	return false 
}