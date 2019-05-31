package terra

type EachMode rune

//go:generate stringer -type=EachMode -output=each_mode_string.go each_mode.go
const (
	NoEach   EachMode = 0
	EachList EachMode = 'L'
	EachMap  EachMode = 'M'
)
