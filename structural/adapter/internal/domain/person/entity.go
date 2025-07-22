package person

import (
	"time"

	"gorm.io/gorm"
)

type Person struct {
	ID        uint   `gorm:"primaryKey;autoIncrement"`
	FirstName string `gorm:"type:varchar(30);not null"`
	LastName  string `gorm:"type:varchar(30);not null"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
