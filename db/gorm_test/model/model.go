package model

import "time"

type Test001 struct {
	ConstructionID uint64 `gorm:"primary_key"`
	Value          int    `gorm:"type:integer;"`
	Name           string `gorm:"type:varchar(45);"`
}

type Test002 struct {
	ConstructionID uint64    `gorm:"primary_key;column:constructionid;"`
	Value          int       `gorm:"column:value" sql:"type:integer;"`
	Name           string    `gorm:"column:name" sql:"type:varchar(45);"`
	CreatedBy      int64     `gorm:"column:createdby;" sql:"not null;type:bigint"`
	CreatedAt      time.Time `gorm:"column:createdat;" sql:"not null;type:timestamp;"`
	UpdatedBy      int64     `gorm:"column:updatedby;" sql:"not null;type:bigint"`
	UpdatedAt      time.Time `gorm:"column:updatedat;" sql:"not null;type:timestamp;"`
}

//type Test002 struct {
//	ConstructionID    uint64 `gorm:"primary_key"`
//	Value int `gorm:"type:integer;"`
//	Name  string `gorm:"type:varchar(45);"`
//	CreatedBy uint16 `gorm:"type:bigint; column:CreatedBy"`
//	RegisteredDate   time.Time `sql:"not null;type:timestamp;" gorm:"column:RegisteredDate"`
//	Updatedby uint16 `gorm:"type:bigint;"`
//	Updatedat   time.Time `sql:"not null;type:timestamp;"`
//}
