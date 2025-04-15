package models

import (
	"github.com/sev-2/raiden/pkg/db"
	"medpoint/internal/types"
	"time"
)

type Reservation struct {
	db.ModelBase
	Id         int32                       `json:"id,omitempty" column:"name:id;type:integer;primaryKey;autoIncrement;nullable:false"`
	UserId     int32                       `json:"user_id,omitempty" column:"name:user_id;type:integer;nullable:false"`
	DoctorId   int32                       `json:"doctor_id,omitempty" column:"name:doctor_id;type:integer;nullable:false"`
	FaskesId   int32                       `json:"faskes_id,omitempty" column:"name:faskes_id;type:integer;nullable:false"`
	ServiceId  int32                       `json:"service_id,omitempty" column:"name:service_id;type:integer;nullable:false"`
	ScheduleId int32                       `json:"schedule_id,omitempty" column:"name:schedule_id;type:integer;nullable:false"`
	Status     types.ReservationStatusEnum `json:"status,omitempty" column:"name:status;type:reservation_status_enum;nullable:false"`
	CreatedAt  *time.Time                  `json:"created_at,omitempty" column:"name:created_at;type:timestampz;nullable;default:now()"`
	UpdatedAt  *time.Time                  `json:"updated_at,omitempty" column:"name:updated_at;type:timestampz;nullable;default:now()"`

	// Table information
	Metadata string `json:"-" schema:"public" tableName:"reservation" rlsEnable:"false" rlsForced:"false"`

	// Access control
	Acl string `json:"-" read:"" write:""`

	// Relations
	Doctor                                  *Doctor           `json:"doctor,omitempty" onUpdate:"no action" onDelete:"no action" join:"joinType:hasOne;primaryKey:id;foreignKey:doctor_id"`
	DoctorScheduleSchedule                  *DoctorSchedules  `json:"doctor_schedule_schedule,omitempty" onUpdate:"no action" onDelete:"no action" join:"joinType:hasOne;primaryKey:id;foreignKey:schedule_id"`
	DoctorsThroughEPrescriptionsReservation []*Doctor         `json:"doctors_through_e_prescriptions_reservation,omitempty" join:"joinType:manyToMany;through:e_prescriptions;sourcePrimaryKey:id;sourceForeignKey:reservation_id;targetPrimaryKey:id;targetForeign:reservation_id"`
	EPrescriptionReservations               []*EPrescriptions `json:"e_prescription_reservations,omitempty" onUpdate:"no action" onDelete:"cascade" join:"joinType:hasMany;primaryKey:id;foreignKey:reservation_id"`
	Faske                                   *Faskes           `json:"faske,omitempty" onUpdate:"no action" onDelete:"no action" join:"joinType:hasOne;primaryKey:id;foreignKey:faskes_id"`
	Service                                 *Services         `json:"service,omitempty" onUpdate:"no action" onDelete:"no action" join:"joinType:hasOne;primaryKey:id;foreignKey:service_id"`
	TransactionReservations                 []*Transaction    `json:"transaction_reservations,omitempty" onUpdate:"no action" onDelete:"no action" join:"joinType:hasMany;primaryKey:id;foreignKey:reservation_id"`
	User                                    *Users            `json:"user,omitempty" onUpdate:"no action" onDelete:"no action" join:"joinType:hasOne;primaryKey:id;foreignKey:user_id"`
}
