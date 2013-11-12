package tracelog

import (
	"fmt"
	"github.com/robfig/revel"
	"sync"
)

var (
	Serialize sync.Mutex
)

//** Started and Completed

func STARTED(routineName string, functionName string) {
	Serialize.Lock()
	defer Serialize.Unlock()
	revel.TRACE.Output(2, fmt.Sprintf("%s : %s : Started\n", routineName, functionName))
}

func STARTEDf(routineName string, functionName string, format string, a ...interface{}) {
	Serialize.Lock()
	defer Serialize.Unlock()
	revel.TRACE.Output(2, fmt.Sprintf("%s : %s : Started : %s\n", routineName, functionName, fmt.Sprintf(format, a...)))
}

func COMPLETED(routineName string, functionName string) {
	Serialize.Lock()
	defer Serialize.Unlock()
	revel.TRACE.Output(2, fmt.Sprintf("%s : %s : Completed\n", routineName, functionName))
}

func COMPLETEDf(routineName string, functionName string, format string, a ...interface{}) {
	Serialize.Lock()
	defer Serialize.Unlock()
	revel.TRACE.Output(2, fmt.Sprintf("%s : %s : Completed : %s\n", routineName, functionName, fmt.Sprintf(format, a...)))
}

func COMPLETED_ERROR(err error, routineName string, functionName string) {
	Serialize.Lock()
	defer Serialize.Unlock()
	revel.ERROR.Output(2, fmt.Sprintf("%s : %s : Completed : %s\n", routineName, functionName, err))
}

func COMPLETED_ERRORf(err error, routineName string, functionName string, format string, a ...interface{}) {
	Serialize.Lock()
	defer Serialize.Unlock()
	revel.ERROR.Output(2, fmt.Sprintf("%s : %s : Completed : %s : %s\n", routineName, functionName, fmt.Sprintf(format, a...), err))
}

//** TRACE

func TRACE(routineName string, functionName string, format string, a ...interface{}) {
	Serialize.Lock()
	defer Serialize.Unlock()
	revel.TRACE.Output(2, fmt.Sprintf("%s : %s : Info : %s\n", routineName, functionName, fmt.Sprintf(format, a...)))
}

//** INFO

func INFO(routineName string, functionName string, format string, a ...interface{}) {
	Serialize.Lock()
	defer Serialize.Unlock()
	revel.INFO.Output(2, fmt.Sprintf("%s : %s : Info : %s\n", routineName, functionName, fmt.Sprintf(format, a...)))
}

//** WARN

func WARN(routineName string, functionName string, format string, a ...interface{}) {
	Serialize.Lock()
	defer Serialize.Unlock()
	revel.WARN.Output(2, fmt.Sprintf("%s : %s : Info : %s\n", routineName, functionName, fmt.Sprintf(format, a...)))
}

//** ERROR

func ERROR(err error, routineName string, functionName string) {
	Serialize.Lock()
	defer Serialize.Unlock()
	revel.ERROR.Output(2, fmt.Sprintf("%s : %s : Info : %s\n", err))
	revel.ERROR.Output(2, fmt.Sprintf("%s : %s : Info : %s\n", routineName, functionName, err))
}

func ERRORf(err error, routineName string, functionName string, format string, a ...interface{}) {
	Serialize.Lock()
	defer Serialize.Unlock()
	revel.ERROR.Output(2, fmt.Sprintf("%s : %s : Info : %s : %s\n", routineName, functionName, fmt.Sprintf(format, a...), err))
}

//** ALERT

// ALERT write to the ERROR destination and sends email alert
func ALERT(subject string, routineName string, functionName string, format string, a ...interface{}) {
	message := fmt.Sprintf("%s : %s : Info : %s\n", routineName, functionName, fmt.Sprintf(format, a...))

	Serialize.Lock()
	defer Serialize.Unlock()
	revel.ERROR.Output(2, message)

	SendEmail(routineName, subject, message)
}

// COMPLETED_ALERT write to the ERROR destination, writes a Completed tag to the log line and sends email alert
func COMPLETED_ALERT(subject string, routineName string, functionName string, format string, a ...interface{}) {
	message := fmt.Sprintf("%s : %s : Completed : %s\n", routineName, functionName, fmt.Sprintf(format, a...))

	Serialize.Lock()
	defer Serialize.Unlock()
	revel.ERROR.Output(2, message)

	SendEmail(routineName, subject, message)
}
