package types

import (
	"github.com/sev-2/raiden"
)

type Transaction struct {
	raiden.TypeBase
}

func (t *Transaction) Name() string {
	return "_transaction"
}

func (r *Transaction) Format() string {
	return "transaction[]"
}

func (r *Transaction) Enums() []string {
	return []string{}
}

func (r *Transaction) Attributes() []string {
	return []string{}
}

func (r *Transaction) Comment() *string {
	return nil
}

