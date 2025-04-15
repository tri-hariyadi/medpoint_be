package types

import (
	"github.com/sev-2/raiden"
)

type Faskes struct {
	raiden.TypeBase
}

func (t *Faskes) Name() string {
	return "_faskes"
}

func (r *Faskes) Format() string {
	return "faskes[]"
}

func (r *Faskes) Enums() []string {
	return []string{}
}

func (r *Faskes) Attributes() []string {
	return []string{}
}

func (r *Faskes) Comment() *string {
	return nil
}

