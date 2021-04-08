package model

import "gin-g/global"

type Stage struct {
	global.MODEL
	ParentId   uint           	  `json:"parent_id" binding:"required,numeric"`
	Parent    *Stage 			  `gorm:"foreignKey:ParentId"`
	Name      string              `json:"name" binding:"required"`
	Status    uint8           	  `json:"status" binding:"numeric"`
	Sort      uint16           	  `json:"sort" binding:"numeric"`
}
