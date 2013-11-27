// Copyright 2013 Ardan Studios. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE handle.

/*
	AppController is implementing the home page
*/
package controllers

import (
	cb "github.com/goinggo/revel-mgo/app/controllers/base"
	"github.com/robfig/revel"
)

//** TYPES

type (
	// App controller controlls home page
	App struct {
		cb.BaseController
	}
)

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
