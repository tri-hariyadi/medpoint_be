package types

import (
	"github.com/sev-2/raiden"
)

type FaskesAddress struct {
	raiden.TypeBase
}

func (t *FaskesAddress) Name() string {
	return "_faskes_address"
}

func (r *FaskesAddress) Format() string {
	return "faskes_address[]"
}

func (r *FaskesAddress) Enums() []string {
	return []string{}
}

func (r *FaskesAddress) Attributes() []string {
	return []string{}
}

func (r *FaskesAddress) Comment() *string {
	return nil
}

