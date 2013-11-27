// Copyright 2013 Ardan Studios. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE handle.

/*
	BuoyController is implementing the buoy specific api

	http://localhost:9000/region/Gulf%20Of%20Mexico
	http://localhost:9000/station/42002
*/
package controllers

import (
	cb "github.com/goinggo/revel-mgo/app/controllers/base"
	"github.com/goinggo/revel-mgo/app/services/buoy"
	"github.com/robfig/revel"
)

//** TYPES

type (
	// Buoy controller the buoy api
	Buoy struct {
		cb.BaseController
	}
)

//** INIT FUNCTION

// init is called when the first request into the controller is made
func init() {
	revel.InterceptMethod((*Buoy).Before, revel.BEFORE)
	revel.InterceptMethod((*Buoy).After, revel.AFTER)
	revel.InterceptMethod((*Buoy).Panic, revel.PANIC)
}

//** CONTROLLER FUNCTIONS

// Stations returns the specified station
func (this *Buoy) Station(stationId string) revel.Result {
	buoyStation, err := buoy.FindStation(this.Services(), stationId)
	if err != nil {
		return this.RenderText(err.Error())
	}

	return this.RenderJson(buoyStation)
}

// Stations returns the specified region
func (this *Buoy) Region(region string) revel.Result {
	buoyStations, err := buoy.FindRegion(this.Services(), region)
	if err != nil {
		return this.RenderText(err.Error())
	}

	return this.RenderJson(buoyStations)
}
