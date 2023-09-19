package models

import (
	"log"

	"github.com/GDGVIT/hexathon23-backend/hexathon-api/internal/database"
)

// InitializeModels creates or migrates all the models
func InitializeModels() {
	MODELS := map[string]interface{}{
		"Team":     &Team{},
		"Item":     &Item{},
		"Category": &Category{},
	}

	for name, model := range MODELS {
		err := database.DB.AutoMigrate(model)
		if err != nil {
			log.Fatal("Failed to initialize model: ", name)
		} else {
			log.Println("Successfully initialized model: ", name)
		}
	}
}
