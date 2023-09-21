package models

import (
	"strings"

	"github.com/GDGVIT/hexathon23-backend/hexathon-api/internal/database"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// Team is the db model for teams table
type Team struct {
	ID                   uuid.UUID `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	Name                 string    `gorm:"unique;not null"`
	Password             string    `gorm:"not null"`
	Logo                 string
	Members              string
	Role                 string           `gorm:"default:participant"`
	Amount               int              `gorm:"default:0"`
	ProblemStatement     ProblemStatement `gorm:"foreignKey:ProblemStatementID;references:ID"`
	ProblemStatementID   *uuid.UUID
	StatementGenerations int    `gorm:"default:3"`
	ItemsPurchased       []Item `gorm:"many2many:team_items;"`
	Submitted            bool	`gorm:"default:false"`
}

// SetMembers sets the members of a team
func (team *Team) SetMembers(members []string) {
	team.Members = strings.Join(members, ",")
}

// GetMembers returns a list of members
func (team *Team) GetMembers() []string {
	return strings.Split(team.Members, ",")
}

// CreateTeam creates a new team
func (team *Team) CreateTeam() error {
	team.Amount = DEFAULT_AMOUNT
	return database.DB.Create(team).Error
}

// UpdateTeam updates a team
func (team *Team) UpdateTeam() error {
	// Save with associations
	return database.DB.Session(&gorm.Session{FullSaveAssociations: true}).Save(team).Error
}

// DeleteTeam deletes a team
func (team *Team) DeleteTeam() error {
	return database.DB.Delete(team).Error
}

// GetTeamByID returns a team by id
func GetTeamByID(id string) (*Team, error) {
	var team Team
	err := database.DB.Where("id = ?", id).First(&team).Error
	return &team, err
}

// CheckTeamNameExists checks if a team name exists
func CheckTeamNameExists(name string) bool {
	var count int64
	database.DB.Model(&Team{}).Where("name = ?", name).Count(&count)
	return count > 0
}

// ValidateTeamPassword validates a team password
func ValidateTeamPassword(password string) bool {
	// No spaces, min 8 chars, max 32 chars
	// At least one uppercase, one lowercase, one digit, one special char
	if len(password) < 8 || len(password) > 32 {
		return false
	}
	var lowercase, uppercase, digit, special bool
	for _, char := range password {
		if char == ' ' {
			return false
		}
		if char >= 'a' && char <= 'z' {
			lowercase = true
		}
		if char >= 'A' && char <= 'Z' {
			uppercase = true
		}
		if char >= '0' && char <= '9' {
			digit = true
		}
		if char == '!' ||
			char == '@' ||
			char == '#' ||
			char == '$' ||
			char == '%' ||
			char == '^' ||
			char == '&' ||
			char == '*' {
			special = true
		}
	}
	if !lowercase || !uppercase || !digit || !special {
		return false
	}
	return true
}

// ValidateTeamName validates a team name
func ValidateTeamName(name string) bool {
	// Only lowercase alphanumeric and underscore allowed
	for _, char := range name {
		if char == ' ' {
			return false
		}
		if char < 'a' || char > 'z' {
			if char < '0' || char > '9' {
				if char != '_' {
					return false
				}
			}
		}
	}
	return true
}

// GetTeamByName returns a team by name
func GetTeamByName(name string) (*Team, error) {
	var team Team
	// Preload all clause associations
	err := database.DB.Preload(clause.Associations).Where("name = ?", name).First(&team).Error
	return &team, err
}

// GetTeams returns all teams
func GetTeams() ([]Team, error) {
	var teams []Team
	// Preload all clause associations
	err := database.DB.Preload(clause.Associations).Find(&teams).Error
	return teams, err
}
