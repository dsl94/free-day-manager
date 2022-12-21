package entity

import (
	"github.com/jinzhu/gorm"
	"time"
)

type FreeDay struct {
	gorm.Model
	Reason    string
	StartDate time.Time
	EndDate   time.Time
	Status    string
	User      User `gorm:"foreignKey:UserId"`
}
