package models

import (
	"github.com/sev-2/raiden/pkg/db"
	"time"
)

type Schedules struct {
	db.ModelBase
	Id        int32      `json:"id,omitempty" column:"name:id;type:integer;primaryKey;autoIncrement;nullable:false"`
	DoctorId  int32      `json:"doctor_id,omitempty" column:"name:doctor_id;type:integer;nullable:false"`
	FaskesId  *int32     `json:"faskes_id,omitempty" column:"name:faskes_id;type:integer;nullable;unique"`
	Date      time.Time  `json:"date,omitempty" column:"name:date;type:date;nullable:false"`
	StartTime time.Time  `json:"start_time,omitempty" column:"name:start_time;type:time;nullable:false"`
	EndTime   time.Time  `json:"end_time,omitempty" column:"name:end_time;type:time;nullable:false"`
	CreatedAt *time.Time `json:"created_at,omitempty" column:"name:created_at;type:timestampz;nullable;default:now()"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" column:"name:updated_at;type:timestampz;nullable;default:now()"`

	// Table information
	Metadata string `json:"-" schema:"public" tableName:"schedules" rlsEnable:"true" rlsForced:"false"`

	// Access control
	Acl string `json:"-" read:"" write:""`

	// Relations
	Doctor                             *Doctors       `json:"doctor,omitempty" onUpdate:"no action" onDelete:"cascade" join:"joinType:hasOne;primaryKey:id;foreignKey:doctor_id"`
	DoctorsThroughReservationSchedule  []*Doctors     `json:"doctors_through_reservation_schedule,omitempty" join:"joinType:manyToMany;through:reservation;sourcePrimaryKey:id;sourceForeignKey:schedule_id;targetPrimaryKey:id;targetForeign:schedule_id"`
	Faske                              *Faskes        `json:"faske,omitempty" onUpdate:"no action" onDelete:"set null" join:"joinType:hasOne;primaryKey:id;foreignKey:faskes_id"`
	FaskesThroughReservationSchedule   []*Faskes      `json:"faskes_through_reservation_schedule,omitempty" join:"joinType:manyToMany;through:reservation;sourcePrimaryKey:id;sourceForeignKey:schedule_id;targetPrimaryKey:id;targetForeign:schedule_id"`
	ReservationSchedules               []*Reservation `json:"reservation_schedules,omitempty" onUpdate:"no action" onDelete:"no action" join:"joinType:hasMany;primaryKey:id;foreignKey:schedule_id"`
	ServicesThroughReservationSchedule []*Services    `json:"services_through_reservation_schedule,omitempty" join:"joinType:manyToMany;through:reservation;sourcePrimaryKey:id;sourceForeignKey:schedule_id;targetPrimaryKey:id;targetForeign:schedule_id"`
	UsersThroughReservationSchedule    []*Users       `json:"users_through_reservation_schedule,omitempty" join:"joinType:manyToMany;through:reservation;sourcePrimaryKey:id;sourceForeignKey:schedule_id;targetPrimaryKey:id;targetForeign:schedule_id"`
}
