package models

import (
	"github.com/sev-2/raiden/pkg/db"
	"time"
)

type Services struct {
	db.ModelBase
	Id          int32      `json:"id,omitempty" column:"name:id;type:integer;primaryKey;autoIncrement;nullable:false"`
	Name        string     `json:"name,omitempty" column:"name:name;type:varchar;nullable:false"`
	Description *string    `json:"description,omitempty" column:"name:description;type:text;nullable"`
	Price       int32      `json:"price,omitempty" column:"name:price;type:integer;nullable:false"`
	PoliId      int32      `json:"poli_id,omitempty" column:"name:poli_id;type:integer;nullable:false"`
	CreatedAt   *time.Time `json:"created_at,omitempty" column:"name:created_at;type:timestampz;nullable;default:now()"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty" column:"name:updated_at;type:timestampz;nullable;default:now()"`

	// Table information
	Metadata string `json:"-" schema:"public" tableName:"services" rlsEnable:"false" rlsForced:"false"`

	// Access control
	Acl string `json:"-" read:"" write:""`

	// Relations
	DoctorsThroughReservationService   []*Doctors     `json:"doctors_through_reservation_service,omitempty" join:"joinType:manyToMany;through:reservation;sourcePrimaryKey:id;sourceForeignKey:service_id;targetPrimaryKey:id;targetForeign:service_id"`
	FaskesThroughReservationService    []*Faskes      `json:"faskes_through_reservation_service,omitempty" join:"joinType:manyToMany;through:reservation;sourcePrimaryKey:id;sourceForeignKey:service_id;targetPrimaryKey:id;targetForeign:service_id"`
	Poli                               *Poli          `json:"poli,omitempty" onUpdate:"no action" onDelete:"cascade" join:"joinType:hasOne;primaryKey:id;foreignKey:poli_id"`
	ReservationServices                []*Reservation `json:"reservation_services,omitempty" onUpdate:"no action" onDelete:"no action" join:"joinType:hasMany;primaryKey:id;foreignKey:service_id"`
	SchedulesThroughReservationService []*Schedules   `json:"schedules_through_reservation_service,omitempty" join:"joinType:manyToMany;through:reservation;sourcePrimaryKey:id;sourceForeignKey:service_id;targetPrimaryKey:id;targetForeign:service_id"`
	UsersThroughReservationService     []*Users       `json:"users_through_reservation_service,omitempty" join:"joinType:manyToMany;through:reservation;sourcePrimaryKey:id;sourceForeignKey:service_id;targetPrimaryKey:id;targetForeign:service_id"`
}
