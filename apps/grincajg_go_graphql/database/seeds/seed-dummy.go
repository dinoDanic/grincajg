package seeds

import (
	"grincajg/database"
	"grincajg/graph/model"
	"log"
)

func seedDummy() {

	childCategories := []model.Category{
		{Name: "Jaja"},
		{Name: "Gljive"},
		{Name: "Mlijeko"},
		{Name: "Sir"},
		{Name: "Vrnje"},
		{Name: "Med"},
		{Name: "Jogurt"},
		{Name: "Paradajz"},
		{Name: "Tikvice"},
		{Name: "Kruh"},
		{Name: "Ulje"},
		{Name: "Vino"},
		{Name: "Margo"},
	}

	// Insert child categories
	for _, category := range childCategories {
		result := database.DB.Create(&category)
		if result.Error != nil {
			log.Fatalf("Failed to seed child categories: %v", result.Error)
		}
	}
}
