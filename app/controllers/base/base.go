package controllerbase

import (
	"github.com/goinggo/revel-mgo/utilities/mongo"
	"github.com/goinggo/revel-mgo/utilities/tracelog"
	"github.com/robfig/revel"
	"labix.org/v2/mgo"
)

//** TYPES

// BaseController contains common properties
type BaseController struct {
	*revel.Controller
	MongoSession *mgo.Session
}

//** INTERCEPT FUNCTIONS

// Before is called prior to the controller method
func (this *BaseController) Before() revel.Result {
	tracelog.TRACE(this.Session.Id(), "controllerbase", this.Request.URL.Path)
	this.MongoSession, _ = mongo.CopyMonotonicSession(this.Session.Id())
	return nil
}

// After is called once the controller method completes
func (this *BaseController) After() revel.Result {
	defer mongo.CloseSession(this.Session.Id(), this.MongoSession)
	tracelog.TRACE(this.Session.Id(), "controllerbase", this.Request.URL.Path)
	return nil
}

// Base returns a pointer of the BaseController type
func (this *BaseController) Base() *BaseController {
	return this
}
