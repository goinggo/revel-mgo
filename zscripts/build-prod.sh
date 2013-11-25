# Build Code For Production

rm -f $GOPATH/revel/revel-mgo
revel build github.com/goinggo/revel-mgo $GOPATH/revel/revel-mgo
date
ls -l $GOPATH/bin/revel-mgo