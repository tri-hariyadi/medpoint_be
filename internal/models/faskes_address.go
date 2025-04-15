package models

import (
	"github.com/sev-2/raiden/pkg/db"
)

type FaskesAddress struct {
	db.ModelBase
	Id         int32  `json:"id,omitempty" column:"name:id;type:integer;primaryKey;autoIncrement;nullable:false"`
	Street     string `json:"street,omitempty" column:"name:street;type:varchar;nullable:false"`
	City       string `json:"city,omitempty" column:"name:city;type:varchar;nullable:false"`
	Province   string `json:"province,omitempty" column:"name:province;type:varchar;nullable:false"`
	PostalCode string `json:"postal_code,omitempty" column:"name:postal_code;type:varchar;nullable:false"`
	Country    string `json:"country,omitempty" column:"name:country;type:varchar;nullable:false"`

	// Table information
	Metadata string `json:"-" schema:"public" tableName:"faskes_address" rlsEnable:"false" rlsForced:"false"`

	// Access control
	Acl string `json:"-" read:"" write:""`

	// Relations
	FaskeAddresses []*Faskes `json:"faske_addresses,omitempty" onUpdate:"no action" onDelete:"cascade" join:"joinType:hasMany;primaryKey:id;foreignKey:address_id"`
}
