package models

import (
	"github.com/sev-2/raiden/pkg/db"
	"time"
)

type Faskes struct {
	db.ModelBase
	Id        int32      `json:"id,omitempty" column:"name:id;type:integer;primaryKey;autoIncrement;nullable:false"`
	Name      string     `json:"name,omitempty" column:"name:name;type:varchar;nullable:false"`
	Type      string     `json:"type,omitempty" column:"name:type;type:varchar;nullable:false"`
	AddressId int32      `json:"address_id,omitempty" column:"name:address_id;type:integer;nullable:false;unique"`
	CreatedAt *time.Time `json:"created_at,omitempty" column:"name:created_at;type:timestampz;nullable;default:now()"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" column:"name:updated_at;type:timestampz;nullable;default:now()"`

	// Table information
	Metadata string `json:"-" schema:"public" tableName:"faskes" rlsEnable:"false" rlsForced:"false"`

	// Access control
	Acl string `json:"-" read:"" write:""`

	// Relations
	DoctorScheduleFaskes                   []*DoctorSchedules `json:"doctor_schedule_faskes,omitempty" onUpdate:"no action" onDelete:"set null" join:"joinType:hasMany;primaryKey:id;foreignKey:faskes_id"`
	DoctorSchedulesThroughReservationFaske []*DoctorSchedules `json:"doctor_schedules_through_reservation_faske,omitempty" join:"joinType:manyToMany;through:reservation;sourcePrimaryKey:id;sourceForeignKey:faskes_id;targetPrimaryKey:id;targetForeign:faskes_id"`
	DoctorsThroughDoctorSchedulesFaske     []*Doctor          `json:"doctors_through_doctor_schedules_faske,omitempty" join:"joinType:manyToMany;through:doctor_schedules;sourcePrimaryKey:id;sourceForeignKey:faskes_id;targetPrimaryKey:id;targetForeign:faskes_id"`
	DoctorsThroughReservationFaske         []*Doctor          `json:"doctors_through_reservation_faske,omitempty" join:"joinType:manyToMany;through:reservation;sourcePrimaryKey:id;sourceForeignKey:faskes_id;targetPrimaryKey:id;targetForeign:faskes_id"`
	FaskesAddressAddress                   *FaskesAddress     `json:"faskes_address_address,omitempty" onUpdate:"no action" onDelete:"cascade" join:"joinType:hasOne;primaryKey:id;foreignKey:address_id"`
	PoliFaskes                             []*Poli            `json:"poli_faskes,omitempty" onUpdate:"no action" onDelete:"cascade" join:"joinType:hasMany;primaryKey:id;foreignKey:faskes_id"`
	ReservationFaskes                      []*Reservation     `json:"reservation_faskes,omitempty" onUpdate:"no action" onDelete:"no action" join:"joinType:hasMany;primaryKey:id;foreignKey:faskes_id"`
	ServicesThroughReservationFaske        []*Services        `json:"services_through_reservation_faske,omitempty" join:"joinType:manyToMany;through:reservation;sourcePrimaryKey:id;sourceForeignKey:faskes_id;targetPrimaryKey:id;targetForeign:faskes_id"`
	UsersThroughReservationFaske           []*Users           `json:"users_through_reservation_faske,omitempty" join:"joinType:manyToMany;through:reservation;sourcePrimaryKey:id;sourceForeignKey:faskes_id;targetPrimaryKey:id;targetForeign:faskes_id"`
}
