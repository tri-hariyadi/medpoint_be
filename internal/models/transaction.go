package models

import (
	"github.com/sev-2/raiden/pkg/db"
	"medpoint/internal/types"
	"time"
)

type Transaction struct {
	db.ModelBase
	Id            int32                       `json:"id,omitempty" column:"name:id;type:integer;primaryKey;autoIncrement;nullable:false"`
	ReservationId int32                       `json:"reservation_id,omitempty" column:"name:reservation_id;type:integer;nullable:false"`
	Amount        int32                       `json:"amount,omitempty" column:"name:amount;type:integer;nullable:false"`
	Status        types.TransactionStatusEnum `json:"status,omitempty" column:"name:status;type:transaction_status_enum;nullable:false"`
	CreatedAt     *time.Time                  `json:"created_at,omitempty" column:"name:created_at;type:timestampz;nullable;default:now()"`
	UpdatedAt     *time.Time                  `json:"updated_at,omitempty" column:"name:updated_at;type:timestampz;nullable;default:now()"`

	// Table information
	Metadata string `json:"-" schema:"public" tableName:"transaction" rlsEnable:"false" rlsForced:"false"`

	// Access control
	Acl string `json:"-" read:"" write:""`

	// Relations
	Reservation *Reservation `json:"reservation,omitempty" onUpdate:"no action" onDelete:"no action" join:"joinType:hasOne;primaryKey:id;foreignKey:reservation_id"`
}
