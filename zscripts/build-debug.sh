# Build Code
# Use to create a build for debugging

rm -f $GOPATH/bin/revel-mgo
cd $GOPATH/src/github.com/goinggo/revel-mgo/app/tmp
go build -x -gcflags "-N -l" -o $ROOT/bin/revel-mgo
date
ls -l $GOPATH/bin/revel-mgo