package global

import (
	"gorm.io/gorm"
)

type MODEL struct {
	ID        uint `json:"id" gorm:"primarykey"`
	CreatedAt int64 `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt int64  `json:"updated_at" gorm:"autoCreateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
