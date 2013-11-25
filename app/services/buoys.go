package services

import (
	cb "github.com/goinggo/revel-mgo/app/controllers/base"
	"github.com/goinggo/revel-mgo/utilities/helper"
	"github.com/goinggo/revel-mgo/utilities/mongo"
	"github.com/goinggo/revel-mgo/utilities/tracelog"
	"labix.org/v2/mgo/bson"
)

//** TYPES

type (
	// BuoyCondition contains information for an individual station
	BuoyCondition struct {
		WindSpeed     float64 `bson:"wind_speed_milehour"`
		WindDirection int     `bson:"wind_direction_degnorth"`
		WindGust      float64 `bson:"gust_wind_speed_milehour"`
	}

	// BuoyLocation contains the buoys location
	BuoyLocation struct {
		Type        string    `bson:"type"`
		Coordinates []float64 `bson:"coordinates"`
	}

	// BuoyStation contains information for an individual station
	BuoyStation struct {
		ID        bson.ObjectId `bson:"_id,omitempty"`
		StationId string        `bson:"station_id"`
		Name      string        `bson:"name"`
		LocDesc   string        `bson:"location_desc"`
		Condition BuoyCondition `bson:"condition"`
		Location  BuoyLocation  `bson:"location"`
	}
)

//** PUBLIC FUNCTIONS

// FindStation retrieves the specified station
func FindStation(controller *cb.BaseController, stationId string) (buoyStation *BuoyStation, err error) {
	defer helper.CatchPanic(&err, controller.Session.Id(), "FindStation")

	tracelog.STARTED(controller.Session.Id(), "FindStation")

	// Access the collection
	collection, err := mongo.GetCollection(controller.MongoSession, helper.MONGO_DATABASE, "buoy_stations")
	if err != nil {
		tracelog.COMPLETED_ERROR(err, helper.MAIN_GO_ROUTINE, "FindStation")
		return buoyStation, err
	}

	// Find all the specified stations
	query := collection.Find(bson.M{"station_id": stationId})

	// Capture the specified buoy
	buoyStation = &BuoyStation{}
	err = query.One(buoyStation)

	tracelog.COMPLETED(controller.Session.Id(), "FindStation")
	return buoyStation, err
}

// FindRegion retrieves the stations for the specified region
func FindRegion(controller *cb.BaseController, region string) (buoyStations []*BuoyStation, err error) {
	defer helper.CatchPanic(&err, controller.Session.Id(), "FindRegion")

	tracelog.STARTED(controller.Session.Id(), "FindRegion")

	// Access the collection
	collection, err := mongo.GetCollection(controller.MongoSession, helper.MONGO_DATABASE, "buoy_stations")
	if err != nil {
		tracelog.COMPLETED_ERROR(err, helper.MAIN_GO_ROUTINE, "FindRegion")
		return buoyStations, err
	}

	// Build the query and display it to the log
	queryMap := bson.M{"region": region}
	tracelog.TRACE(helper.MAIN_GO_ROUTINE, "FindRegion", "Query : %s", mongo.DisplayQuery(queryMap))

	// Find all the specified stations
	query := collection.Find(queryMap)

	// Capture the specified buoy
	buoyStations = []*BuoyStation{}
	err = query.All(&buoyStations)

	tracelog.COMPLETED(controller.Session.Id(), "FindRegion")
	return buoyStations, err
}
