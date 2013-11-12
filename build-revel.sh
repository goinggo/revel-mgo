# Build revel tool

rm -f $GOPATH/bin/revel
cd github.com/robfig/revel/revel
go build -o $GOPATH/bin/revel
date
ls -l $GOPATH/bin/revel