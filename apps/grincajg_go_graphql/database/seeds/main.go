package main

import (
	"fmt"
	"grincajg/database"
	"grincajg/env"
	"grincajg/graph/model"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	env.LoadEnv()
	database.Connect()

	seedCategories()
	seedUsersAndOrganizations()
	seedProducts()

}

func seedCategories() {
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

func seedUsersAndOrganizations() {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("1"), bcrypt.DefaultCost)

	if err != nil {
		log.Fatalf("Failed to hash password")
	}

	users := []model.User{
		{Name: "User 1", Email: "user1@gmail.com", Password: string(hashedPassword)},
		{Name: "User 2", Email: "user2@gmail.com", Password: string(hashedPassword)},
		{Name: "User 3", Email: "user3@gmail.com", Password: string(hashedPassword)},
		{Name: "User 4", Email: "user4@gmail.com", Password: string(hashedPassword)},
		{Name: "User 5", Email: "user5@gmail.com", Password: string(hashedPassword)},
		{Name: "User 6", Email: "user6@gmail.com", Password: string(hashedPassword)},
		{Name: "User 7", Email: "user7@gmail.com", Password: string(hashedPassword)},
	}

	for _, user := range users {

		result := database.DB.Create(&user)
		if result.Error != nil {
			log.Fatalf("Failed to seed:  %v", result.Error)
		}

	}

	allUsers := []model.User{}

	database.DB.Find(&allUsers)

	for i, findUser := range allUsers {
		organization := model.Organization{
			Name:        fmt.Sprintf("Organization %d", i+1),
			AdminUserID: findUser.ID,
			Latitude:    10.0,
			Longitude:   10.0,
		}
		database.DB.Create(&organization)
	}

}

func seedProducts() {
	allOrgniazations := []model.Organization{}
	database.DB.Find(&allOrgniazations)

	for _, organization := range allOrgniazations {
		products := []model.Product{
			{Name: "Product 1", CategoryID: 1},
			{Name: "Product 2", CategoryID: 2},
			{Name: "Product 3", CategoryID: 3},
			{Name: "Product 4", CategoryID: 4},
			{Name: "Product 5123", CategoryID: 5},
			{Name: "Product 2131", CategoryID: 6},
			{Name: "Product 62", CategoryID: 7},
			{Name: "Product 51", CategoryID: 8},
			{Name: "Product 41", CategoryID: 9},
			{Name: "Product 31", CategoryID: 10},
			{Name: "Product 12", CategoryID: 11},
			{Name: "Product 16", CategoryID: 12},
			{Name: "Product 3", CategoryID: 13},
		}
		for _, product := range products {
			createProduct := model.Product{
				Name:           product.Name,
				CategoryID:     product.CategoryID,
				OrganizationID: organization.ID,
			}
			database.DB.Create(&createProduct)
		}
	}

}
