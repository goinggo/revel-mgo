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
	tracelog.TRACE(this.Session.Id(), "Before", this.Request.URL.Path)

	var err error
	this.MongoSession, err = mongo.CopyMonotonicSession(this.Session.Id())
	if err != nil {
		tracelog.ERRORf(err, this.Session.Id(), "Before", this.Request.URL.Path)
		return this.RenderError(err)
	}

	return nil
}

// After is called once the controller method completes
func (this *BaseController) After() revel.Result {
	defer func() {
		if this.MongoSession != nil {
			mongo.CloseSession(this.Session.Id(), this.MongoSession)
			this.MongoSession = nil
		}
	}()

	tracelog.TRACE(this.Session.Id(), "After", this.Request.URL.Path)
	return nil
}

// After is called once the controller method completes
func (this *BaseController) Panic() revel.Result {
	defer func() {
		mongo.CloseSession(this.Session.Id(), this.MongoSession)
		this.MongoSession = nil
	}()

	tracelog.TRACE(this.Session.Id(), "Panic", this.Request.URL.Path)
	return nil
}

// Base returns a pointer of the BaseController type
func (this *BaseController) Base() *BaseController {
	return this
}
