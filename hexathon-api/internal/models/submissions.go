package models

import (
	"time"

	"github.com/GDGVIT/hexathon23-backend/hexathon-api/internal/database"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Submission struct {
	ID                 uuid.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	CreatedAt          time.Time
	Team               Team `gorm:"foreignKey:TeamID"`
	TeamID             uuid.UUID
	ProblemStatement   ProblemStatement `gorm:"foreignKey:ProblemStatementID"`
	ProblemStatementID uuid.UUID
	FigmaURL           string
	DocURL             string
}

// Creates a submission
func (submission *Submission) CreateSubmission() error {
	return database.DB.Session(&gorm.Session{FullSaveAssociations: true}).Create(submission).Error
}