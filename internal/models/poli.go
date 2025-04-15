package models

import (
	"github.com/sev-2/raiden/pkg/db"
	"time"
)

type Poli struct {
	db.ModelBase
	Id        int32      `json:"id,omitempty" column:"name:id;type:integer;primaryKey;autoIncrement;nullable:false"`
	Name      string     `json:"name,omitempty" column:"name:name;type:varchar;nullable:false"`
	FaskesId  int32      `json:"faskes_id,omitempty" column:"name:faskes_id;type:integer;nullable:false"`
	CreatedAt *time.Time `json:"created_at,omitempty" column:"name:created_at;type:timestampz;nullable;default:now()"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" column:"name:updated_at;type:timestampz;nullable;default:now()"`

	// Table information
	Metadata string `json:"-" schema:"public" tableName:"poli" rlsEnable:"false" rlsForced:"false"`

	// Access control
	Acl string `json:"-" read:"" write:""`

	// Relations
	Faske        *Faskes     `json:"faske,omitempty" onUpdate:"no action" onDelete:"cascade" join:"joinType:hasOne;primaryKey:id;foreignKey:faskes_id"`
	ServicePolis []*Services `json:"service_polis,omitempty" onUpdate:"no action" onDelete:"cascade" join:"joinType:hasMany;primaryKey:id;foreignKey:poli_id"`
}
