package navigator

type Technique struct {
	Tactic            string     `json:"tactic,omitempty" yaml:"tactic,omitempty"`
	Comment           string     `json:"comment,omitempty" yaml:"comment,omitempty"`
	TechniqueID       string     `json:"techniqueID" yaml:"techniqueID"`
	Enabled           bool       `json:"enabled" yaml:"enabled"`
	Color             string     `json:"color,omitempty" yaml:"color,omitempty"`
	Score             *int       `json:"score,omitempty" yaml:"score,omitempty"`
	Metadata          []Metadata `json:"metadata,omitempty" yaml:"metadata,omitempty"`
	Links             []Link     `json:"links,omitempty" yaml:"links,omitempty"`
	ShowSubtechniques bool       `json:"showSubtechniques,omitempty" yaml:"showSubtechniques,omitempty"`
}

func NewTechnique(techniqueID string) (*Technique, error) {
	technique := &Technique{
		TechniqueID: techniqueID,
		Enabled:     true,
	}
	return technique, nil
}
