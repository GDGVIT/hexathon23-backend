package models

import "github.com/google/uuid"

type Submission struct {
	ID                 uuid.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	Team               Team      `gorm:"foreignKey:TeamID"`
	TeamID             string
	ProblemStatement   ProblemStatement `gorm:"foreignKey:ProblemStatementID"`
	ProblemStatementID string
	FigmaURL           string
	VideoURL           string
	Description        string
}
