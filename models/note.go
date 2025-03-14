package models

import (
	"time"

	"github.com/google/uuid"
)

// Modèle de paste sécurisé
type Note struct {
	Id        uuid.UUID `gorm:"type:char(36);primaryKey" json:"id"`
	Content   string    `gorm:"type:text;not null"`
	CreatedAt time.Time `json:"created_at"`
	ExpiresAt time.Time `json:"expires_at"`
}
