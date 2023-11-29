package navigator

const (
	FlatLayout = "flat"
	SideLayout = "side"
	MiniLayout = "mini"
)

var (
	DefaultLayoutType     = FlatLayout
	ShowId                = true
	ShowName              = true
	ShowAggregateScores   = false
	CountUnscored         = false
	AggregateFunction     = Avg
	ExpandedSubtechniques = "all"
)

const (
	Avg = "average"
	Sum = "sum"
	Max = "max"
	Min = "min"
)

func NewLayout() (*Layout, error) {
	layout := &Layout{
		Layout:                DefaultLayoutType,
		ShowId:                ShowId,
		ShowName:              ShowName,
		ShowAggregateScores:   ShowAggregateScores,
		CountUnscored:         CountUnscored,
		AggregateFunction:     Avg,
		ExpandedSubtechniques: ExpandedSubtechniques,
	}
	return layout, nil
}

type Layout struct {
	Layout                string `json:"layout" yaml:"layout"`
	ShowId                bool   `json:"showID" yaml:"showID"`
	ShowName              bool   `json:"showName" yaml:"showName"`
	ShowAggregateScores   bool   `json:"showAggregateScores" yaml:"showAggregateScores"`
	CountUnscored         bool   `json:"countUnscored" yaml:"countUnscored"`
	AggregateFunction     string `json:"aggregateFunction" yaml:"aggregateFunction"`
	ExpandedSubtechniques string `json:"expandedSubtechniques" yaml:"expandedSubtechniques"`
}
