package models

import (
	"github.com/sev-2/raiden/pkg/db"
	"time"
)

type EPrescriptions struct {
	db.ModelBase
	Id            int32      `json:"id,omitempty" column:"name:id;type:integer;primaryKey;autoIncrement;nullable:false"`
	ReservationId int32      `json:"reservation_id,omitempty" column:"name:reservation_id;type:integer;nullable:false"`
	DoctorId      int32      `json:"doctor_id,omitempty" column:"name:doctor_id;type:integer;nullable:false"`
	Medications   string     `json:"medications,omitempty" column:"name:medications;type:text;nullable:false"`
	CreatedAt     *time.Time `json:"created_at,omitempty" column:"name:created_at;type:timestampz;nullable;default:now()"`
	UpdatedAt     *time.Time `json:"updated_at,omitempty" column:"name:updated_at;type:timestampz;nullable;default:now()"`

	// Table information
	Metadata string `json:"-" schema:"public" tableName:"e_prescriptions" rlsEnable:"false" rlsForced:"false"`

	// Access control
	Acl string `json:"-" read:"" write:""`

	// Relations
	Doctor      *Doctors     `json:"doctor,omitempty" onUpdate:"no action" onDelete:"cascade" join:"joinType:hasOne;primaryKey:id;foreignKey:doctor_id"`
	Reservation *Reservation `json:"reservation,omitempty" onUpdate:"no action" onDelete:"cascade" join:"joinType:hasOne;primaryKey:id;foreignKey:reservation_id"`
}
