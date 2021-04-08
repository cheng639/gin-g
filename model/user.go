package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	UUID		uuid.UUID      `json:"-" gorm:"primaryKey"`
	RoleId      uint           `json:"role_id" `
	Openid      string         `json:"openid" `
	Unionid     string         `json:"unionid" `
	Nickname    string         `json:"nickname"`
	Avatar      string         `json:"avatar"`
	Gender      uint           `json:"gender"`
	Phone       string         `json:"phone" gorm:"unique"`
	Status      uint           `json:"status"`
	CreatedAt   int64          `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt   int64          `json:"updated_at" gorm:"autoCreateTime"`
	DeletedAt gorm.DeletedAt   `json:"-"`
}
