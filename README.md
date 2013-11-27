# Revel Mgo Example

Copyright 2013 Ardan Studios. All rights reserved.  
Use of this source code is governed by a BSD-style license that can be found in the LICENSE handle.

This application provides a sample to use the revel web framework and the Go MongoDB driver mgo. This program connects to a public MongoDB at MongoLab. A single collection is available for testing. The configuration can be found in the app.conf file.

The project also includes several shell scripts to make building and running the web application easier.

Ardan Studios  
12973 SW 112 ST, Suite 153  
Miami, FL 33186  
bill@ardanstudios.com

### Installation

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

### LiteIDE
If you use LiteIDE, add this to your gosrc.xml file under Preferences/LiteBuild. You must have the
run.go file open in the editor to use these commands from within LiteIDE.
	
<action id="Run Revel" img="blue/run.png" cmd="sh" args="$(EDITOR_DIR)/zscripts/run.sh" output="true" codec="utf-8" readline="true"/>

<action id="Install Debug" menu="Run Revel" img="blue/install.png" cmd="sh" args="$(EDITOR_DIR)/zscripts/build-debug.sh" output="true" codec="utf-8" readline="true"/>

<action id="R-Test" menu="Run Revel" img="blue/test.png" cmd="sh" args="$(EDITOR_DIR)/zscripts/run-tests.sh" output="true" codec="utf-8" readline="true"/>


Then in your Build Configuration add your project import path to the TARGETARGS

	TARGETARGS: github.com/goinggo/revel-mgo
	
This will allow you to run revel from inside of LiteIDE

### Notes About Architecture

I have been asked why I have organized the code in this way?

For me the controller should do nothing more than call into the service layer. The service layer contains the business logic for processing the request.

The controller methods just exist to receive the request and send the response. The more that can be abstracted into the base controller the better. This way, adding a new controller methods is simple and you don't need to worry about forgetting to do something important. Authentication always comes to mind.

The interceptor is being used to perform operations before and after the controller is called. Mongo related stuff is done there for now. Exception handling should be done with an interceptor as well.

The models folder contains the data structures for the individual services. Each service places their models in a separate folder.

The services folder contains the base service code and then an individual folder for each service.

The utilities folder is just that, support for the web application, mostly used by the services. You have exception handling support, extended logging support and the mongo support.

Init should be self explanatory. Anything that needs to be initialized before you handle your first request should be done there. Revel does not have an application end event or construct, so there is no closing of those resources.

The abstraction layer for executing MongoDB queries and commands help hide the boilerplate code away into the base service and mongo utility code.
