package models

import (
	"github.com/GDGVIT/hexathon23-backend/hexathon-api/internal/database"
	"github.com/google/uuid"
)

type Participant struct {
	ID        uuid.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	Name      string
	RegNo     string
	Email     string
	TeamID    *uuid.UUID
	Team      Team
	CheckedIn bool `gorm:"default:false"`
}

// CreateParticipant creates a new participant
func (participant *Participant) CreateParticipant() error {
	return database.DB.Create(participant).Error
}
