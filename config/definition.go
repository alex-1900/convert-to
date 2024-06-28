package config

import "fmt"

type ProID string

type Resolver interface {
	Resolve([]byte) error
	ToClash() (Profile, error)
	ToSingBox() (Profile, error)
}

type Profile interface {
	fmt.Stringer
	ProID() ProID
}
