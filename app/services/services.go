// Copyright 2013 Ardan Studios. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE handle.

/*
	Services provides boilerplate functionality for all services.
	Any state required by all the services is maintained here.
*/
package services

import (
	"github.com/goinggo/revel-mgo/utilities/helper"
	"github.com/goinggo/revel-mgo/utilities/mongo"
	"labix.org/v2/mgo"
)

//** TYPES

type (
	// Services contains common fields and behavior for all services
	Service struct {
		MongoSession *mgo.Session
		UserId       string
	}
)

//** PUBLIC FUNCTIONS

// DBAction executes queries and commands against MongoDB
func (this *Service) DBAction(collectionName string, mongoCall mongo.MongoCall) (err error) {
	return mongo.Execute(this.UserId, this.MongoSession, helper.MONGO_DATABASE, collectionName, mongoCall)
}
