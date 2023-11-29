package navigator

import (
	"github.com/mitchellh/mapstructure"
)

func ParseLayer(data map[string]interface{}) (*Layer, error) {
	var layer Layer
	techniques, _ := data["techniques"].([]interface{})
	for i, technique := range techniques {
		m, _ := technique.(map[string]interface{})

		// Set `enabled` to `true` by default.
		if _, ok := m["enabled"]; !ok {
			m["enabled"] = true
			techniques[i] = m
		}
	}
	err := mapstructure.Decode(data, &layer)
	if err != nil {
		return nil, err
	}
	return &layer, nil
}

