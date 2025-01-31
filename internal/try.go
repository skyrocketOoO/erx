package Error

/*
Extend from go built-in error
Support
- Logging
- Call trace
- Only Log once
- Custom errorcode
- Trace ID to search log
*/

// var logInstance logger

// type logger interface {
// 	Log(content string) error
// }

// // add logger
// func EmbeddedLogger(logger logger) {
// 	logInstance = logger
// }

// type CustomError struct {
// 	Code     uint
// 	LogMsg   LogMessage
// 	ErrorMsg ErrorMessage
// }

// type ErrorMessage struct {
// 	LogId uuid.UUID
// 	Msg   string
// }

// type LogMessage struct {
// 	LogId uuid.UUID
// 	Trace string
// 	Msg   string
// }

// func (e *CustomError) Error() string {
// 	jsonstr, _ := json.Marshal(e.ErrorMsg)
// 	return string(jsonstr)
// }

// func New(msg string) (out error) {
// 	return Wrap(errors.New(msg))
// }

// func Wrap(in error) (out error) {
// 	if in == nil {
// 		return nil
// 	}

// 	// if error == customError means already logged
// 	if _, ok := in.(*CustomError); ok {
// 		return in
// 	}

// 	logId := uuid.New().String()
// 	trace := getCallStack(3)
// 	msg := fmt.Sprintf("uuid: %s \nerror: %s\ntrace: %s", logId, in.Error(), trace)
// 	logInstance.Log(msg)

// 	return &CustomError{
// 		ErrorMsg: msg,
// 	}
// }

// func getCallStack(callerSkip ...int) (stkMsg string) {
// 	pc := make([]uintptr, 10)
// 	skip := 2
// 	if len(callerSkip) > 0 {
// 		skip = callerSkip[0]
// 	}
// 	n := runtime.Callers(skip, pc)

// 	frames := runtime.CallersFrames(pc[:n])

// 	stkMsg = "Call stack:"

// 	for {
// 		frame, more := frames.Next()
// 		stkMsg += fmt.Sprintf("%s\n\t%s:%d\n", frame.Function, frame.File, frame.Line)
// 		if !more {
// 			break
// 		}
// 	}
// 	return stkMsg
// }
