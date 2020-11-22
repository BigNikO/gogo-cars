package app

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

type driver struct {
	gorm.Model
	Name    string
	License string
	Cars    []car
}

func (app *App) getDriver(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var driver driver
	var cars []car
	app.db.First(&driver, params["id"])
	app.db.Model(&driver).Related(&cars)
	driver.Cars = cars
	json.NewEncoder(w).Encode(&driver)
}
