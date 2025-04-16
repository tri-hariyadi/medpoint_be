package models

import (
	"github.com/google/uuid"
	"github.com/sev-2/raiden/pkg/db"
	"medpoint/internal/types"
	"time"
)

type Users struct {
	db.ModelBase
	Id        int32          `json:"id,omitempty" column:"name:id;type:integer;primaryKey;autoIncrement;nullable:false"`
	Role      types.RoleEnum `json:"role,omitempty" column:"name:role;type:role_enum;nullable:false"`
	FullName  string         `json:"full_name,omitempty" column:"name:full_name;type:varchar;nullable:false"`
	Email     string         `json:"email,omitempty" column:"name:email;type:varchar;nullable:false;unique"`
	Password  string         `json:"password,omitempty" column:"name:password;type:varchar;nullable:false"`
	Phone     *string        `json:"phone,omitempty" column:"name:phone;type:varchar;nullable"`
	Avatar    *string        `json:"avatar,omitempty" column:"name:avatar;type:varchar;nullable"`
	CreatedAt *time.Time     `json:"created_at,omitempty" column:"name:created_at;type:timestampz;nullable;default:now()"`
	UpdatedAt *time.Time     `json:"updated_at,omitempty" column:"name:updated_at;type:timestampz;nullable;default:now()"`
	AuthId    *uuid.UUID     `json:"auth_id,omitempty" column:"name:auth_id;type:uuid;nullable;unique"`

	// Table information
	Metadata string `json:"-" schema:"public" tableName:"users" rlsEnable:"true" rlsForced:"false"`

	// Access control
	Acl string `json:"-" read:"" write:""`

	// Relations
	DoctorUsers                     []*Doctors     `json:"doctor_users,omitempty" onUpdate:"no action" onDelete:"cascade" join:"joinType:hasMany;primaryKey:id;foreignKey:user_id"`
	DoctorsThroughReservationUser   []*Doctors     `json:"doctors_through_reservation_user,omitempty" join:"joinType:manyToMany;through:reservation;sourcePrimaryKey:id;sourceForeignKey:user_id;targetPrimaryKey:id;targetForeign:user_id"`
	FaskesThroughReservationUser    []*Faskes      `json:"faskes_through_reservation_user,omitempty" join:"joinType:manyToMany;through:reservation;sourcePrimaryKey:id;sourceForeignKey:user_id;targetPrimaryKey:id;targetForeign:user_id"`
	ReservationUsers                []*Reservation `json:"reservation_users,omitempty" onUpdate:"no action" onDelete:"no action" join:"joinType:hasMany;primaryKey:id;foreignKey:user_id"`
	SchedulesThroughReservationUser []*Schedules   `json:"schedules_through_reservation_user,omitempty" join:"joinType:manyToMany;through:reservation;sourcePrimaryKey:id;sourceForeignKey:user_id;targetPrimaryKey:id;targetForeign:user_id"`
	ServicesThroughReservationUser  []*Services    `json:"services_through_reservation_user,omitempty" join:"joinType:manyToMany;through:reservation;sourcePrimaryKey:id;sourceForeignKey:user_id;targetPrimaryKey:id;targetForeign:user_id"`
	UserAuth                        *Users         `json:"user_auth,omitempty" onUpdate:"no action" onDelete:"cascade" join:"joinType:hasOne;primaryKey:id;foreignKey:auth_id"`
}
