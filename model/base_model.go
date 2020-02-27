package model

import (
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"time"
)

type BaseModelOnlyId struct {
	ID uint `gorm:"primary_key"`
}

type BaseModel struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (BaseModel) create() {

}
