package types

import (
	"github.com/sev-2/raiden"
)

type ReservationStatusEnum struct {
	raiden.TypeBase
}

func (t *ReservationStatusEnum) Name() string {
	return "reservation_status_enum"
}

func (r *ReservationStatusEnum) Format() string {
	return "reservation_status_enum"
}

func (r *ReservationStatusEnum) Enums() []string {
	return []string{"Booked","Rescheduled","Cancelled","Completed"}
}

func (r *ReservationStatusEnum) Attributes() []string {
	return []string{}
}

func (r *ReservationStatusEnum) Comment() *string {
	return nil
}

