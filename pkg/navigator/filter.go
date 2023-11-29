package navigator

import "errors"

var (
	EnterprisePlatforms = []string{"PRE", "Windows", "Linux", "macOS", "Network", "AWS", "GCP", "Azure", "Azure AD", "Office 365", "SaaS"}
	MobilePlatforms     = []string{"Android", "iOS"}
	ICSPlatforms        = []string{"Windows", "Control Server", "Data Historian", "Engineering Workstation", "Field Controller/RTU/PLC/IED", "Human-Machine Interface", "Input/Output Server", "Safety Instrumented System/Protection Relay"}
)

var (
	PlatformsByDomain = map[string][]string{
		EnterpriseAttack: EnterprisePlatforms,
		MobileAttack:     MobilePlatforms,
		ICSAttack:        ICSPlatforms,
	}
)

type Filter struct {
	Platforms []string `json:"platforms" yaml:"platforms"`
}

func NewFilter(domain string) (*Filter, error) {
	platforms, ok := PlatformsByDomain[domain]
	if !ok {
		return nil, errors.New("invalid domain")
	}
	return &Filter{
		Platforms: platforms,
	}, nil
}
