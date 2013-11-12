# Build Code
# Use to test builds before running

rm -f $GOPATH/bin/revel-mgo
cd app/tmp
go build -o $GOPATH/bin/revel-mgo
date
ls -l $GOPATH/bin/revel-mgo