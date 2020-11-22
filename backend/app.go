package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

// App type
type App struct {
	db *gorm.DB
}

// Serve function
func (app *App) Serve() {
	controller := cors.AllowAll()
	router := mux.NewRouter()

	app.db.AutoMigrate(&driver{})
	app.db.AutoMigrate(&car{})

	router.HandleFunc("/cars", app.getCars).Methods("GET")
	router.HandleFunc("/cars/{id}", app.getCar).Methods("GET")
	router.HandleFunc("/cars", app.addCar).Methods("POST")
	router.HandleFunc("/drivers/{id}", app.getDriver).Methods("GET")
	router.HandleFunc("/cars/{id}", app.deleteCar).Methods("DELETE")

	handler := controller.Handler(router)
	port := GetEnvData("SERVER_PORT")

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), handler))
}

// New application
func New(db *gorm.DB) *App {
	return &App{
		db: db,
	}
}

// GetEnvData ...
func GetEnvData(key string) string {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}
