package models

import (
	"github.com/sev-2/raiden/pkg/db"
	"time"
)

type Doctors struct {
	db.ModelBase
	Id             int32      `json:"id,omitempty" column:"name:id;type:integer;primaryKey;autoIncrement;nullable:false"`
	UserId         int32      `json:"user_id,omitempty" column:"name:user_id;type:integer;nullable:false;unique"`
	Specialization string     `json:"specialization,omitempty" column:"name:specialization;type:varchar;nullable:false"`
	Experience     int32      `json:"experience,omitempty" column:"name:experience;type:integer;nullable:false"`
	CreatedAt      *time.Time `json:"created_at,omitempty" column:"name:created_at;type:timestampz;nullable;default:now()"`
	UpdatedAt      *time.Time `json:"updated_at,omitempty" column:"name:updated_at;type:timestampz;nullable;default:now()"`

	// Table information
	Metadata string `json:"-" schema:"public" tableName:"doctors" rlsEnable:"true" rlsForced:"false"`

	// Access control
	Acl string `json:"-" read:"" write:""`

	// Relations
	EPrescriptionDoctors                    []*EPrescriptions `json:"e_prescription_doctors,omitempty" onUpdate:"no action" onDelete:"cascade" join:"joinType:hasMany;primaryKey:id;foreignKey:doctor_id"`
	FaskesThroughReservationDoctor          []*Faskes         `json:"faskes_through_reservation_doctor,omitempty" join:"joinType:manyToMany;through:reservation;sourcePrimaryKey:id;sourceForeignKey:doctor_id;targetPrimaryKey:id;targetForeign:doctor_id"`
	FaskesThroughSchedulesDoctor            []*Faskes         `json:"faskes_through_schedules_doctor,omitempty" join:"joinType:manyToMany;through:schedules;sourcePrimaryKey:id;sourceForeignKey:doctor_id;targetPrimaryKey:id;targetForeign:doctor_id"`
	ReservationDoctors                      []*Reservation    `json:"reservation_doctors,omitempty" onUpdate:"no action" onDelete:"no action" join:"joinType:hasMany;primaryKey:id;foreignKey:doctor_id"`
	ReservationsThroughEPrescriptionsDoctor []*Reservation    `json:"reservations_through_e_prescriptions_doctor,omitempty" join:"joinType:manyToMany;through:e_prescriptions;sourcePrimaryKey:id;sourceForeignKey:doctor_id;targetPrimaryKey:id;targetForeign:doctor_id"`
	ScheduleDoctors                         []*Schedules      `json:"schedule_doctors,omitempty" onUpdate:"no action" onDelete:"cascade" join:"joinType:hasMany;primaryKey:id;foreignKey:doctor_id"`
	SchedulesThroughReservationDoctor       []*Schedules      `json:"schedules_through_reservation_doctor,omitempty" join:"joinType:manyToMany;through:reservation;sourcePrimaryKey:id;sourceForeignKey:doctor_id;targetPrimaryKey:id;targetForeign:doctor_id"`
	ServicesThroughReservationDoctor        []*Services       `json:"services_through_reservation_doctor,omitempty" join:"joinType:manyToMany;through:reservation;sourcePrimaryKey:id;sourceForeignKey:doctor_id;targetPrimaryKey:id;targetForeign:doctor_id"`
	User                                    *Users            `json:"user,omitempty" onUpdate:"no action" onDelete:"cascade" join:"joinType:hasOne;primaryKey:id;foreignKey:user_id"`
	UsersThroughReservationDoctor           []*Users          `json:"users_through_reservation_doctor,omitempty" join:"joinType:manyToMany;through:reservation;sourcePrimaryKey:id;sourceForeignKey:doctor_id;targetPrimaryKey:id;targetForeign:doctor_id"`
}
