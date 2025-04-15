package types

import (
	"github.com/sev-2/raiden"
)

type DoctorSchedules struct {
	raiden.TypeBase
}

func (t *DoctorSchedules) Name() string {
	return "_doctor_schedules"
}

func (r *DoctorSchedules) Format() string {
	return "doctor_schedules[]"
}

func (r *DoctorSchedules) Enums() []string {
	return []string{}
}

func (r *DoctorSchedules) Attributes() []string {
	return []string{}
}

func (r *DoctorSchedules) Comment() *string {
	return nil
}

