package terra

import (
	"encoding/json"
	"io"
)

// all operations related to the file system
func (s *StateV4) WriteV4(w io.Writer) error {
	s.normalize()
	src, err := json.MarshalIndent(s, "", " ")
	if err != nil {
		return err
	}

	src = append(src, '\n')
	_, err = w.Write(src)
	return err
}
