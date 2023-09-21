package models

import (
	"log"

	"github.com/GDGVIT/hexathon23-backend/hexathon-api/internal/database"
)

var MODELS map[string]interface{} = map[string]interface{}{
	"Team":             &Team{},
	"Item":             &Item{},
	"Category":         &Category{},
	"Submission":       &Submission{},
	"ProblemStatement": &ProblemStatement{},
	"Cart":             &Cart{},
}

var DEFAULT_AMOUNT int

// InitializeModels creates or migrates all the models
func InitializeModels(defaultAmount int) {
	DEFAULT_AMOUNT = defaultAmount
	for name, model := range MODELS {
		err := database.DB.AutoMigrate(model)
		if err != nil {
			log.Fatal("Failed to initialize model: ", name)
		} else {
			log.Println("Successfully initialized model: ", name)
		}
	}
}
