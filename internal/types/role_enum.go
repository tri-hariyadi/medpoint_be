package types

import (
	"github.com/sev-2/raiden"
)

type RoleEnum struct {
	raiden.TypeBase
}

func (t *RoleEnum) Name() string {
	return "role_enum"
}

func (r *RoleEnum) Format() string {
	return "role_enum"
}

func (r *RoleEnum) Enums() []string {
	return []string{"Super Admin","Admin","Doctor","Patient"}
}

func (r *RoleEnum) Attributes() []string {
	return []string{}
}

func (r *RoleEnum) Comment() *string {
	return nil
}

