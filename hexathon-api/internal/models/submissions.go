package models

import "github.com/google/uuid"

type Submission struct {
	ID                 uuid.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	Team               Team      `gorm:"foreignKey:TeamID"`
	TeamID             uuid.UUID
	ProblemStatement   ProblemStatement `gorm:"foreignKey:ProblemStatementID"`
	ProblemStatementID uuid.UUID
	FigmaURL           string
	VideoURL           string
	Description        string
}
