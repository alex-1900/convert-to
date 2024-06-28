package resolver

import (
	"github.com/alex-1900/convert-to/config"
	"github.com/alex-1900/convert-to/structs/profile/clash"
	"gopkg.in/yaml.v3"
)

type Clash struct {
	entity *clash.Clash
}

func NewClash(entity *clash.Clash) *Clash {
	return &Clash{entity: entity}
}

func (c Clash) Resolve(bs []byte) error {
	return yaml.Unmarshal(bs, c.entity)
}

func (c Clash) ToClash() (config.Profile, error) {
	return c.entity, nil
}

func (c Clash) ToSingBox() (config.Profile, error) {
	return c.entity, nil
}
