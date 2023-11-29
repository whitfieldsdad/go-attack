package navigator

const (
	NavigatorVersion = "4.9.0"
	LayerVersion     = "4.5"
)

type Version struct {
	Attack    string `json:"attack" yaml:"attack"`
	Navigator string `json:"navigator" yaml:"navigator"`
	Layer     string `json:"layer" yaml:"layer"`
}

func NewVersion(attackVersion, navigatorVersion, layerVersion string) (*Version, error) {
	if navigatorVersion == "" {
		navigatorVersion = NavigatorVersion
	}
	if layerVersion == "" {
		layerVersion = LayerVersion
	}
	return &Version{
		Attack:    attackVersion,
		Navigator: NavigatorVersion,
		Layer:     LayerVersion,
	}, nil
}
