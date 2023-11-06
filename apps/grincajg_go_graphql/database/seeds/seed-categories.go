package seeds

import (
	"grincajg/database"
	"grincajg/graph/model"
	"log"
)

func seedCategories() {
	var count int64
	database.DB.Model(&model.Category{}).Count(&count)

	if count > 0 {
		log.Println("Seeding not required")
		return
	}

	// Create the parent category
	parentCategory := model.Category{Name: "Jaja"}

	result := database.DB.Create(&parentCategory)

	if result.Error != nil {
		log.Fatalf("Failed to seed parent category: %v", result.Error)
	}

	// Retrieve the ID of the newly created parent category
	parentID := parentCategory.ID

	// Define child categories with the parent ID
	childCategories := []model.Category{
		{Name: "Guščja jaja", CategoryID: &parentID},
		{Name: "Kokošja jaja", CategoryID: &parentID},
		{Name: "Nojeva jaja", CategoryID: &parentID},
		{Name: "Pačja jaja", CategoryID: &parentID},
		{Name: "Prepreličja jaja", CategoryID: &parentID},
	}

	// Insert child categories
	for _, category := range childCategories {
		result := database.DB.Create(&category)
		if result.Error != nil {
			log.Fatalf("Failed to seed child categories: %v", result.Error)
		}
	}

	log.Println("Category seeding completed.")
}
