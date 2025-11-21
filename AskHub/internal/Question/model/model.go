package model

import (
	"gorm.io/gorm"
)

type Question struct {
	*gorm.Model
	Text string `gorm:"type:text;not null" json:"text"`
}
