package loader

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"os"
)

var database *gorm.DB

func init() {
	env := godotenv.Load()

	if env != nil {
		fmt.Println(env)
	}

	host := os.Getenv("database_host")
	username := os.Getenv("database_username")
	password := os.Getenv("database_password")
	name := os.Getenv("database_name")

	params := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", host, username, name, password)

	connection, err := gorm.Open("postgres", params)
	if err != nil {
		fmt.Print(err)
	}

	database = connection
}

func GetDatabase() *gorm.DB {
	return database
}
