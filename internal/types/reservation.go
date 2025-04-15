package types

import (
	"github.com/sev-2/raiden"
)

type Reservation struct {
	raiden.TypeBase
}

func (t *Reservation) Name() string {
	return "_reservation"
}

func (r *Reservation) Format() string {
	return "reservation[]"
}

func (r *Reservation) Enums() []string {
	return []string{}
}

func (r *Reservation) Attributes() []string {
	return []string{}
}

func (r *Reservation) Comment() *string {
	return nil
}

