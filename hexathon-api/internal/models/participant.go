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
	Team      Team `gorm:"foreignKey:TeamID;references:ID;constraint:OnDelete:SET NULL;"`
	CheckedIn bool `gorm:"default:false"`
}

// CreateParticipant creates a new participant
func (participant *Participant) CreateParticipant() error {
	return database.DB.Create(participant).Error
}

// UpdateParticipant updates a participant
func (participant *Participant) UpdateParticipant() error {
	return database.DB.Save(participant).Error
}

// DeleteParticipant deletes a participant
func (participant *Participant) DeleteParticipant() error {
	return database.DB.Delete(participant).Error
}

// CheckParticipantExists checks if a participant exists
func CheckParticipantExists(regNo string) bool {
	var participant Participant
	database.DB.Where("reg_no = ?", regNo).First(&participant)
	return participant.ID != uuid.Nil
}

// GetParticipants returns a list of all participants
func GetParticipants() ([]Participant, error) {
	var participants []Participant
	err := database.DB.Find(&participants).Error
	return participants, err
}

// GetParticipantsNotCheckedIn returns a list of all participants who have not checked in
func GetParticipantsNotCheckedIn() ([]Participant, error) {
	var participants []Participant
	err := database.DB.Where("checked_in = ?", false).Find(&participants).Error
	return participants, err
}

// GetParticipantByID returns a participant by id
func GetParticipantByID(id string) (*Participant, error) {
	var participant Participant
	err := database.DB.Where("id = ?", id).First(&participant).Error
	return &participant, err
}

// GetParticipantByEmail returns a participant by email
func GetParticipantByEmail(email string) (*Participant, error) {
	var participant Participant
	err := database.DB.Where("email = ?", email).First(&participant).Error
	return &participant, err
}

// GetParticipantByRegNo returns a participant by reg no
func GetParticipantByRegNo(regNo string) (*Participant, error) {
	var participant Participant
	err := database.DB.Where("reg_no = ?", regNo).First(&participant).Error
	return &participant, err
}

// SearchParticipantByRegNo returns a participant by reg no
func SearchParticipantByRegNo(regNo string) ([]Participant, error) {
	var participant []Participant
	err := database.DB.Where("reg_no LIKE ?", regNo+"%").Find(&participant).Error
	return participant, err
}

// SearchParticipantByRegNo who have not checked in
func SearchParticipantByRegNoNotCheckedIn(regNo string) ([]Participant, error) {
	var participant []Participant
	err := database.DB.Where("reg_no LIKE ? AND checked_in = ?", regNo+"%", false).Find(&participant).Error
	return participant, err
}

// SearchParticipantByName returns a participant by name
func SearchParticipantByName(name string) ([]Participant, error) {
	var participant []Participant
	err := database.DB.Where("name LIKE ?", name+"%").Find(&participant).Error
	return participant, err
}

// SearchParticipantByName who have not checked in
func SearchParticipantByNameNotCheckedIn(name string) ([]Participant, error) {
	var participant []Participant
	err := database.DB.Where("name LIKE ? AND checked_in = ?", name+"%", false).Find(&participant).Error
	return participant, err
}

// Searches for a participant by name or reg no who have not checked in
func SearchParticipantNotCheckedIn(query string) ([]Participant, error) {
	var participant []Participant
	err := database.DB.Where("name LIKE ? OR reg_no LIKE ? AND checked_in = ?", query+"%", query+"%", false).Find(&participant).Error
	return participant, err
}

// GetParticipantsCheckedIn returns a list of all participants who have checked in
func GetParticipantsCheckedIn() ([]Participant, error) {
	var participants []Participant
	err := database.DB.Where("checked_in = ?", true).Find(&participants).Error
	return participants, err
}

// Searches for a participant by name or reg no who have checked in
func SearchParticipantCheckedIn(name string) ([]Participant, error) {
	var participant []Participant
	err := database.DB.Where("name LIKE ? OR reg_no LIKE ? AND checked_in = ?", name+"%", name+"%", true).Find(&participant).Error
	return participant, err
}
