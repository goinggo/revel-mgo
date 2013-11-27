// Copyright 2013 Ardan Studios. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE handle.

/*
	Buoy implements the service for the buoy functionality
*/
package buoy

import (
	"github.com/goinggo/revel-mgo/app/models/buoy"
	"github.com/goinggo/revel-mgo/app/services"
	"github.com/goinggo/revel-mgo/utilities/helper"
	"github.com/goinggo/revel-mgo/utilities/mongo"
	"github.com/goinggo/revel-mgo/utilities/tracelog"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

//** PUBLIC FUNCTIONS

// FindStation retrieves the specified station
func FindStation(service *services.Service, stationId string) (buoyStation *buoyModel.BuoyStation, err error) {
	defer helper.CatchPanic(&err, service.UserId, "FindStation")

	tracelog.STARTED(service.UserId, "FindStation")

	// Find the specified station id
	queryMap := bson.M{"station_id": stationId}
	tracelog.TRACE(helper.MAIN_GO_ROUTINE, "FindStation", "Query : %s", mongo.ToString(queryMap))

	// Execute the query
	buoyStation = &buoyModel.BuoyStation{}
	err = service.DBAction("buoy_stations",
		func(collection *mgo.Collection) error {
			return collection.Find(queryMap).One(buoyStation)
		})

	if err != nil {
		tracelog.COMPLETED_ERROR(err, helper.MAIN_GO_ROUTINE, "FindStation")
		return buoyStation, err
	}

	tracelog.COMPLETED(service.UserId, "FindStation")
	return buoyStation, err
}

// FindRegion retrieves the stations for the specified region
func FindRegion(service *services.Service, region string) (buoyStations []*buoyModel.BuoyStation, err error) {
	defer helper.CatchPanic(&err, service.UserId, "FindRegion")

	tracelog.STARTED(service.UserId, "FindRegion")

	// Find the specified region
	queryMap := bson.M{"region": region}
	tracelog.TRACE(helper.MAIN_GO_ROUTINE, "FindRegion", "Query : %s", mongo.ToString(queryMap))

	// Capture the specified buoy
	buoyStations = []*buoyModel.BuoyStation{}
	err = service.DBAction("buoy_stations",
		func(collection *mgo.Collection) error {
			return collection.Find(queryMap).All(&buoyStations)
		})

	if err != nil {
		tracelog.COMPLETED_ERROR(err, helper.MAIN_GO_ROUTINE, "FindRegion")
		return buoyStations, err
	}

	tracelog.COMPLETED(service.UserId, "FindRegion")
	return buoyStations, err
}
