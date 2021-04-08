package model

import "gin-g/global"

type Book struct {
	global.MODEL
	StageId   uint           	  `json:"stage_id" binding:"required,numeric"`
	Stage     Stage			      `gorm:"foreignKey:StageId"`
	Name      string              `json:"name" binding:"required"`
	Cover     string              `json:"cover" binding:"required,gt=1"`
	Content   string              `json:"content"`
	Link      string              `json:"link"`
	Status    uint8           	  `json:"status" binding:"numeric"`
	Sort      uint16           	  `json:"sort" binding:"numeric"`
}
