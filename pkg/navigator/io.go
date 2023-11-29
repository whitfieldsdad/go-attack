package navigator

import (
	"encoding/json"
	"io"
	"os"
)

func ReadLayer(path string) (*Layer, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	blob, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}
	m := make(map[string]interface{})
	err = json.Unmarshal(blob, &m)
	if err != nil {
		return nil, err
	}
	return ParseLayer(m)
}
