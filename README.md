# Revel Mgo Example

Copyright 2013 Ardan Studios. All rights reserved.  
Use of this source code is governed by a BSD-style license that can be found in the LICENSE handle.

This application provides a sample to use the revel web framework and the Go MongoDB driver mgo. This program connects to a public MongoDB at MongoLab. A single collection is available for testing. The configuration can be found in the app.conf file.

The project also includes several shell scripts to make building and running the web application easier.

Ardan Studios  
12973 SW 112 ST, Suite 153  
Miami, FL 33186  
bill@ardanstudios.com

	-- Get, build and install the code
	go get github.com/goinggo/revel-mgo
	
	-- Run the code
	cd $GOPATH/src/github.com/goinggo/revel-mgo
	./run.sh
	
	-- Test Web Service API's
	
	This will return a single station from Mongo
	http://localhost:9000/station/42002
	
	This will return a collection of stations for the region
	http://localhost:9000/region/Gulf%20Of%20Mexico
