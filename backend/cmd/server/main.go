package main

import (
	"fmt"
	"log"

	app "gogo.cars"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // postgres
)

func main() {
	dbName := app.GetEnvData("DATABASE_NAME")
	dbPassword := app.GetEnvData("DATABASE_PASSWORD")
	host := app.GetEnvData("BASE_URL")

	db, err := gorm.Open("postgres", fmt.Sprintf("host=%s port=5432 user=postgres dbname=%s sslmode=disable password=%s", host, dbName, dbPassword))

	if err != nil {
		log.Fatalf("Failed to connect database")
	}

	defer db.Close()

	app := app.New(db)
	app.Serve()
}
