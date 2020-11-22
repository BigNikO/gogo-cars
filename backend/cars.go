package app

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

type car struct {
	gorm.Model
	Year      int
	Make      string
	ModelName string
	DriverID  int
}

func (app *App) getCars(w http.ResponseWriter, r *http.Request) {
	var cars []car
	app.db.Find(&cars)
	json.NewEncoder(w).Encode(&cars)
}

func (app *App) getCar(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var car car
	app.db.First(&car, params["id"])
	json.NewEncoder(w).Encode(&car)
}

func (app *App) deleteCar(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var c car
	app.db.First(&c, params["id"])
	app.db.Delete(&c)
	var cars []car
	app.db.Find(&cars)
	json.NewEncoder(w).Encode(&cars)
}

type myData struct {
	Name string
	name string
}

func (app *App) addCar(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var car car
	err := decoder.Decode(&car)
	if err != nil {
		panic(err)
	}
	println(car.DriverID)
	println(car.ModelName)
	println(car.Year)
	println(car.Make)

	if app.db.Create(&car).Error != nil {
		println("Cannot create a new car")
	}

}
