package attack

import (
	"github.com/TcM1911/stix2"
	"github.com/pkg/errors"
	"github.com/whitfieldsdad/go-attack/pkg/util"
)

const (
	DefaultAttackEnterprisePath = "data/stix2.1/enterprise-attack.json"
)

type AttackClient struct {
	Collection *stix2.Collection
}

func NewAttackClient(path string) (*AttackClient, error) {
	if path == "" {
		path = DefaultAttackEnterprisePath
	}
	blob, err := util.ReadJsonFile(path)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to read STIX 2 JSON file")
	}
	collection, err := stix2.FromJSON(blob)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to parse STIX 2 content")
	}
	c := &AttackClient{
		Collection: collection,
	}
	return c, nil
}

func NewBuiltinEnterpriseAttackClient() (*AttackClient, error) {
	return NewAttackClient(DefaultAttackEnterprisePath)
}

func (c AttackClient) ListTactics() ([]stix2.STIXObject, error) {
	return c.ListObjectsByType("x-mitre-tactic")
}

func (c AttackClient) ListTechniques() ([]stix2.STIXObject, error) {
	return c.ListObjectsByType("attack-pattern")
}

func (c AttackClient) ListMalware() ([]stix2.STIXObject, error) {
	return c.ListObjectsByType("malware")
}

func (c AttackClient) ListTools() ([]stix2.STIXObject, error) {
	return c.ListObjectsByType("tool")
}

func (c AttackClient) ListSoftware() ([]stix2.STIXObject, error) {
	malware, err := c.ListMalware()
	if err != nil {
		return nil, err
	}
	tools, err := c.ListTools()
	if err != nil {
		return nil, err
	}
	software := append(malware, tools...)
	return software, nil
}

func (c AttackClient) ListMitigations() ([]stix2.STIXObject, error) {
	return c.ListObjectsByType("course-of-action")
}

func (c AttackClient) ListGroups() ([]stix2.STIXObject, error) {
	return c.ListObjectsByType("intrusion-set")
}

func (c AttackClient) ListRelationships() ([]stix2.STIXObject, error) {
	return c.ListObjectsByType("relationship")
}

func (c AttackClient) ListDataSources() ([]stix2.STIXObject, error) {
	return c.ListObjectsByType("x-mitre-data-source")
}

func (c AttackClient) ListDataComponents() ([]stix2.STIXObject, error) {
	return c.ListObjectsByType("x-mitre-data-component")
}

func (c AttackClient) ListCampaigns() ([]stix2.STIXObject, error) {
	return c.ListObjectsByType("x-mitre-data-component")
}

func (c AttackClient) ListObjectsByType(t string) ([]stix2.STIXObject, error) {
	var objects []stix2.STIXObject
	stixType := stix2.STIXType(t)
	for _, o := range c.Collection.GetAll(stixType) {
		props := o.GetExtendedTopLevelProperties()
		if props == nil {
			continue
		}
		if t == string(o.GetType()) {
			objects = append(objects, o)
		}
	}
	if len(objects) == 0 {
		return nil, errors.Errorf("No objects of type %s found", t)
	}
	return objects, nil
}
