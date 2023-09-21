package models

import (
	"math/rand"
	"time"

	"github.com/GDGVIT/hexathon23-backend/hexathon-api/internal/database"
	"github.com/google/uuid"
)

type ProblemStatement struct {
	ID          uuid.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	Name        string    `gorm:"unique;not null"`
	Description string
}

// CreateProblemStatement creates a new problem statement
func (problemStatement *ProblemStatement) CreateProblemStatement() error {
	return database.DB.Create(problemStatement).Error
}

// UpdateProblemStatement updates an existing problem statement
func (problemStatement *ProblemStatement) UpdateProblemStatement() error {
	return database.DB.Save(problemStatement).Error
}

// DeleteProblemStatement deletes an existing problem statement
func (problemStatement *ProblemStatement) DeleteProblemStatement() error {
	return database.DB.Delete(problemStatement).Error
}

// GetProblemStatement returns a problem statement
func GetProblemStatement(id string) (ProblemStatement, error) {
	var problemStatement ProblemStatement
	err := database.DB.Where("id = ?", id).First(&problemStatement).Error
	return problemStatement, err
}

// GetProblemStatements returns a list of all problem statements
func GetProblemStatements() ([]ProblemStatement, error) {
	var problemStatements []ProblemStatement
	err := database.DB.Find(&problemStatements).Error
	return problemStatements, err
}

// ValidateProblemStatementName validates a problem statement name
func ValidateProblemStatementName(name string) bool {
	return len(name) > 0
}

// GenerateProblemStatementForTeam generates a problem statement for a team
func GenerateProblemStatementForTeam(team *Team) (*ProblemStatement, error) {
	if team.StatementGenerations <= 0 {
		return nil, nil
	}
	// Get a random problem statement
	problemStatements, err := GetProblemStatements()
	if err != nil {
		return nil, err
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	problemStatement := problemStatements[r.Intn(len(problemStatements))]
	team.StatementGenerations--
	team.ProblemStatementID = &problemStatement.ID
	team.ProblemStatement = problemStatement
	err = team.UpdateTeam()
	if err != nil {
		return nil, err
	}
	return &problemStatement, nil
}
