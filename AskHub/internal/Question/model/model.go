package model

import (
	"time"

	"github.com/LashkaPashka/AskHub/internal/Answer/model"
)

type Question struct {
	ID        uint           `gorm:"primaryKey"`
	Text      string         `gorm:"type:text;not null"`
	Answers   []model.Answer `gorm:"constraint:OnDelete:CASCADE;"`
	CreatedAt time.Time
	UpdatedAt *time.Time
}
