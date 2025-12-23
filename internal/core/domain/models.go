package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BaseModel struct {
	ID        uuid.UUID      `gorm:"type:uuid;primary_key;defualt;gen_random_uuid()" json:"id"`
	CreatedAt time.Time      `json:"createdat"`
	UpdatedAt time.Time      `json:"updatedat"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedat,omitempty"`
}

type Role struct {
	BaseModel
	Name string
}
