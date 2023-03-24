package main

import (
	"fmt"
	"log"

	"github.com/ggichuru/the_go_blog/initializers"
	"github.com/ggichuru/the_go_blog/models"
)

// Load env variables into memory and create a connection pool to the postgreSQL db
func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}

	initializers.ConnectDB(&config)
}

// Evoke the GORM automigrate method to migrate the schema to the DB
func main() {
	initializers.DB.AutoMigrate(&models.User{}, &models.Post{})
	fmt.Println("? Migration Complete")
}
