package navigator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewLayer(t *testing.T) {
	layer, err := NewLayer("test", EnterpriseAttack)
	assert.Nil(t, err)
	assert.Equal(t, "test", layer.Name)

	layer.AddTechniques(Technique{TechniqueID: "T1234"})
	assert.Equal(t, 1, len(layer.Techniques))
}

func TestNewLayerWithColors(t *testing.T) {
	layer, err := NewLayer("test", EnterpriseAttack)
	assert.Nil(t, err)
	assert.Equal(t, "test", layer.Name)

	layer.AddTechniques(Technique{TechniqueID: "T1234", Color: "blue"})
	layer.AddTechniques(Technique{TechniqueID: "T1235", Color: "red"})
	assert.Equal(t, 2, len(layer.Techniques))
}
