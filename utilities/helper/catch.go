package helper

import (
	"fmt"
	"github.com/goinggo/revel-mgo/utilities/tracelog"
	"runtime"
)

// CatchPanic is used to catch any Panic and log exceptions to Stdout. It will also write the stack trace
func CatchPanic(err *error, sessionId string, functionName string) {
	if r := recover(); r != nil {
		buf := make([]byte, 10000)
		runtime.Stack(buf, false)

		tracelog.ALERT(tracelog.EmailAlertSubject, sessionId, functionName, "PANIC Defered [%v] : Stack Trace : %v", r, string(buf))

		if err != nil {
			*err = fmt.Errorf("%v", r)
		}
	}
}
