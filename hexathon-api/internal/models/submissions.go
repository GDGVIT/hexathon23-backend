package models

import (
	"time"

	"github.com/GDGVIT/hexathon23-backend/hexathon-api/internal/database"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Submission struct {
	ID                 uuid.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	CreatedAt          time.Time `gorm:"autoCreateTime"`
	Team               Team      `gorm:"foreignKey:TeamID"`
	TeamID             uuid.UUID
	ProblemStatement   ProblemStatement `gorm:"foreignKey:ProblemStatementID;references:ID"`
	ProblemStatementID uuid.UUID
	FigmaURL           string
	DocURL             string
}

// Creates a submission
func (submission *Submission) CreateSubmission() error {
	return database.DB.Session(&gorm.Session{FullSaveAssociations: true}).Create(submission).Error
}

// Updates a submission
func (submission *Submission) UpdateSubmission() error {
	return database.DB.Session(&gorm.Session{FullSaveAssociations: true}).Save(submission).Error
}

// Deletes a submission
func (submission *Submission) DeleteSubmission() error {
	return database.DB.Delete(submission).Error
}

// Retrieves submissions by Team ID
func GetSubmissionByTeamID(id string) (*Submission, error) {
	var submission Submission
	err := database.DB.Preload(clause.Associations).Where("team_id = ?", id).First(&submission).Error
	return &submission, err
}

// Retrieves submission by submission ID
func GetSubmissionByID(id string) (*Submission, error) {
	var submission Submission
	err := database.DB.Preload(clause.Associations).Where("id = ?", id).First(&submission).Error
	return &submission, err
}

// GetSubmissions returns all the submissions
func GetSubmissions() ([]Submission, error) {
	var submissions []Submission
	// Preload all clause associations
	err := database.DB.Preload(clause.Associations).Find(&submissions).Error
	return submissions, err
}
