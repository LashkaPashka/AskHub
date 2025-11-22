package model

import (
	"gorm.io/gorm"
)

type Answer struct {
	gorm.Model
	QuestionID uint   `gorm:"not null" json:"question_id"`
	UserID     string `gorm:"type:uuid;not null" json:"user_id"`
	Text       string `gorm:"type:text;not null" json:"text"`
}
