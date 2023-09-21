package models

import (
	"time"

	"github.com/GDGVIT/hexathon23-backend/hexathon-api/internal/database"
	"github.com/google/uuid"
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

// func ValidateLink(link string) bool {
// 	return true
// }

func (submission *Submission) CreateSubmission() error {
	return database.DB.Create(submission).Error
}
