package controllers

import (
	cb "github.com/goinggo/revel-mgo/app/controllers/base"
	"github.com/goinggo/revel-mgo/app/services"
	"github.com/robfig/revel"
)

//** TYPES

// App is the controller type
type App struct {
	cb.BaseController
}

//** INIT FUNCTION

// init is called when the first request into the controller is made
func init() {
	revel.InterceptMethod((*App).Before, revel.BEFORE)
	revel.InterceptMethod((*App).After, revel.AFTER)
	revel.InterceptMethod((*App).Panic, revel.PANIC)
}

//** CONTROLLER FUNCTIONS

// Index called to render the home page
func (this *App) Index() revel.Result {
	return this.Render()
}

// Stations returns the specified station
func (this *App) Station(stationId string) revel.Result {
	buoyStation, err := services.FindStation(this.Base(), stationId)
	if err != nil {
		return this.RenderText(err.Error())
	}

	return this.RenderJson(buoyStation)
}

// Stations returns the specified region
func (this *App) Region(region string) revel.Result {
	buoyStations, err := services.FindRegion(this.Base(), region)
	if err != nil {
		return this.RenderText(err.Error())
	}

	return this.RenderJson(buoyStations)
}
