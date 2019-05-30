package tfterms

// using the same terms as terraform
// to create the status

// "mode": "data",
// "mode": "managed",

//go:generate stringer -type=ResourceMode -output=resource_mode_string.go resource_mode.go
type ResourceMode int

const(
     ManagedResourceMode ResourceMode = iota
     DataResourceMode
)
