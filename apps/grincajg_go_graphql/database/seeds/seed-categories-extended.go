package seeds

import (
	"grincajg/database"
	"grincajg/graph/model"
	"log"
)

func seedEggs() {
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

func seedMushrooms() {
	var count int64
	database.DB.Model(&model.Category{}).Where("name = ?", "Gljive").Count(&count)

	if count > 0 {
		log.Println("Mushrooms seeding not required")
		return
	}

	// Create the parent category for mushrooms
	parentMushroomCategory := model.Category{Name: "Gljive"}
	result := database.DB.Create(&parentMushroomCategory)
	if result.Error != nil {
		log.Fatalf("Failed to seed parent mushroom category: %v", result.Error)
	}

	// Retrieve the ID of the newly created parent mushroom category
	parentMushroomID := parentMushroomCategory.ID

	// Define child mushroom categories with the parent ID
	childMushroomCategories := []model.Category{
		{Name: "Bukovače", CategoryID: &parentMushroomID},
		{Name: "Gotovi proizvodi", CategoryID: &parentMushroomID},
		{Name: "Lisičarke", CategoryID: &parentMushroomID},
		{Name: "Ostalo", CategoryID: &parentMushroomID},
		{Name: "Šampinjoni", CategoryID: &parentMushroomID},
		{Name: "Shiitake gljive", CategoryID: &parentMushroomID},
		{Name: "Sunčanice", CategoryID: &parentMushroomID},
		{Name: "Tartufi", CategoryID: &parentMushroomID},
		{Name: "Vrganji", CategoryID: &parentMushroomID},
	}

	// Insert child mushroom categories
	for _, mushroomCategory := range childMushroomCategories {
		result := database.DB.Create(&mushroomCategory)
		if result.Error != nil {
			log.Fatalf("Failed to seed child mushroom categories: %v", result.Error)
		}
	}

	log.Println("Mushroom category seeding completed.")
}

func seedHerbsAndSpices() {
	var count int64
	database.DB.Model(&model.Category{}).Where("name = ?", "Bilje i začini").Count(&count)

	if count > 0 {
		log.Println("Herbs and spices seeding not required")
		return
	}

	// Create the parent category for herbs and spices
	parentHerbsSpicesCategory := model.Category{Name: "Bilje i začini"}
	result := database.DB.Create(&parentHerbsSpicesCategory)
	if result.Error != nil {
		log.Fatalf("Failed to seed parent herbs and spices category: %v", result.Error)
	}

	// Retrieve the ID of the newly created parent herbs and spices category
	parentHerbsSpicesID := parentHerbsSpicesCategory.ID

	// Define child herbs and spices categories with the parent ID
	childHerbsSpicesCategories := []model.Category{
		{Name: "Ostalo", CategoryID: &parentHerbsSpicesID},
		{Name: "Začini", CategoryID: &parentHerbsSpicesID},
	}

	// Insert child herbs and spices categories
	for _, herbsSpicesCategory := range childHerbsSpicesCategories {
		result := database.DB.Create(&herbsSpicesCategory)
		if result.Error != nil {
			log.Fatalf("Failed to seed child herbs and spices categories: %v", result.Error)
		}
	}

	log.Println("Herbs and spices category seeding completed.")
}

func seedFatsAndOils() {
	var count int64
	// Check if the parent category already exists
	database.DB.Model(&model.Category{}).Where("name = ?", "Masti i ulja").Count(&count)

	if count > 0 {
		log.Println("Fats and oils category seeding not required")
		return
	}

	// Create the parent category for fats and oils
	parentFatsOilsCategory := model.Category{Name: "Masti i ulja"}
	result := database.DB.Create(&parentFatsOilsCategory)
	if result.Error != nil {
		log.Fatalf("Failed to seed parent fats and oils category: %v", result.Error)
	}

	// Retrieve the ID of the newly created parent fats and oils category
	parentFatsOilsID := parentFatsOilsCategory.ID

	// Define child fats and oils categories with the parent ID
	childFatsOilsCategories := []model.Category{
		{Name: "Biljne masti", CategoryID: &parentFatsOilsID},
		{Name: "Ostalo", CategoryID: &parentFatsOilsID},
		{Name: "Ulja", CategoryID: &parentFatsOilsID},
	}

	// Insert child fats and oils categories
	for _, category := range childFatsOilsCategories {
		result := database.DB.Create(&category)
		if result.Error != nil {
			log.Fatalf("Failed to seed child fats and oils categories: %v", result.Error)
		}
	}

	log.Println("Fats and oils category seeding completed.")
}


func seedDairyProducts() {
	var count int64
	// Check if the parent category already exists
	database.DB.Model(&model.Category{}).Where("name = ?", "Mliječni proizvodi").Count(&count)

	if count > 0 {
		log.Println("Dairy products category seeding not required")
		return
	}

	// Create the parent category for dairy products
	parentDairyCategory := model.Category{Name: "Mliječni proizvodi"}
	result := database.DB.Create(&parentDairyCategory)
	if result.Error != nil {
		log.Fatalf("Failed to seed parent dairy category: %v", result.Error)
	}

	// Retrieve the ID of the newly created parent dairy category
	parentDairyID := parentDairyCategory.ID

	// Define child dairy categories with the parent ID
	childDairyCategories := []model.Category{
		{Name: "Jogurt", CategoryID: &parentDairyID},
		{Name: "Kiselo mlijeko", CategoryID: &parentDairyID},
		{Name: "Kobilje mlijeko", CategoryID: &parentDairyID},
		{Name: "Maslac", CategoryID: &parentDairyID},
		{Name: "Mlijeko", CategoryID: &parentDairyID},
		{Name: "Ostalo", CategoryID: &parentDairyID},
		{Name: "Sir", CategoryID: &parentDairyID},
		{Name: "Vrhnje", CategoryID: &parentDairyID},
	}

	// Insert child dairy categories
	for _, dairyCategory := range childDairyCategories {
		result := database.DB.Create(&dairyCategory)
		if result.Error != nil {
			log.Fatalf("Failed to seed child dairy categories: %v", result.Error)
		}
	}
	log.Println("Dairy products category seeding completed.")
}
