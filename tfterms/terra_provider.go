package tfterms

// ResourceProvider
type TerraProvider interface {
	Resources() []ResourceType
}

// ToDo fetch resource type from the list
// terraform aws -> cloudformation + pkill available resources structs
type ResourceType interface {
	ProviderCanonicalResourceName() string
	HCL() string
}
