package navigator

const (
	EnterpriseAttack = "enterprise-attack"
	MobileAttack     = "mobile-attack"
	ICSAttack        = "ics-attack"
)

func NewLayer(name, domain string) (*Layer, error) {
	layer := &Layer{
		Name:   name,
		Domain: domain,
	}
	return layer, nil
}

func NewEnterpriseLayer(name string) (*Layer, error) {
	return NewLayer(name, EnterpriseAttack)
}

func NewMobileLayer(name string) (*Layer, error) {
	return NewLayer(name, MobileAttack)
}

func NewICSLayer(name string) (*Layer, error) {
	return NewLayer(name, ICSAttack)
}

type Layer struct {
	Name                          string       `json:"name" yaml:"name"`
	Description                   string       `json:"description" yaml:"description"`
	Domain                        string       `json:"domain" yaml:"domain"`
	Techniques                    []Technique  `json:"techniques" yaml:"techniques"`
	Versions                      *Version     `json:"versions" yaml:"versions"`
	CustomDataURL                 string       `json:"customDataURL" yaml:"customDataURL"`
	Filters                       *Filter      `json:"filters" yaml:"filters"`
	Sorting                       int          `json:"sorting" yaml:"sorting"`
	Layout                        *Layout      `json:"layout" yaml:"layout"`
	HideDisabled                  bool         `json:"hideDisabled" yaml:"hideDisabled"`
	Gradient                      *Gradient    `json:"gradient" yaml:"gradient"`
	LegendItems                   []LegendItem `json:"legendItems" yaml:"legendItems"`
	ShowTacticRowBackground       bool         `json:"showTacticRowBackground" yaml:"showTacticRowBackground"`
	TacticRowBackground           string       `json:"tacticRowBackground" yaml:"tacticRowBackground"`
	SelectTechniquesAcrossTactics bool         `json:"selectTechniquesAcrossTactics" yaml:"selectTechniquesAcrossTactics"`
	SelectSubtechniquesWithParent bool         `json:"selectSubtechniquesWithParent" yaml:"selectSubtechniquesWithParent"`
	SelectVisibleTechniques       bool         `json:"selectVisibleTechniques" yaml:"selectVisibleTechniques"`
	Metadata                      []Metadata   `json:"metadata" yaml:"metadata"`
	Links                         []Link       `json:"links" yaml:"links"`
}

func (layer Layer) GetSelectedTechniques() []Technique {
	var results []Technique
	for _, technique := range layer.Techniques {
		if technique.TechniqueID == "" || technique.Color == "" {
			continue
		}
		if !technique.Enabled {
			continue
		}
		results = append(results, technique)
	}
	return results
}

func (layer Layer) GetSelectedTechniqueIds() []string {
	var ids []string
	for _, technique := range layer.GetSelectedTechniques() {
		ids = append(ids, technique.TechniqueID)
	}
	return ids
}
