package types

import (
	"github.com/sev-2/raiden"
)

type TransactionStatusEnum struct {
	raiden.TypeBase
}

func (t *TransactionStatusEnum) Name() string {
	return "transaction_status_enum"
}

func (r *TransactionStatusEnum) Format() string {
	return "transaction_status_enum"
}

func (r *TransactionStatusEnum) Enums() []string {
	return []string{"Pending","Paid","Failed","Refunded"}
}

func (r *TransactionStatusEnum) Attributes() []string {
	return []string{}
}

func (r *TransactionStatusEnum) Comment() *string {
	return nil
}

