package common

import (
	"fmt"
	"time"
)

// NotImplementedException is not implemented exception
type NotImplementedException struct {
	method string
}

func (e *NotImplementedException) Error() string {
	return e.method + " is not implemented."
}

// RuntimeException is runtime exception
type RuntimeException struct {
	msg        string
	context    interface{}
	actionedAt time.Time
}

func (e *RuntimeException) Error() string {
	return fmt.Sprintf("%s %s context: %v", e.actionedAt.Format("2006-01-02 03:04:05"), e.msg, e.context)
}

// NewRuntimeException is return RuntimeException
func NewRuntimeException(msg string, context interface{}) *RuntimeException {
	return &RuntimeException{
		msg:        msg,
		context:    context,
		actionedAt: time.Now(),
	}
}

// InvalidParameterException is return InvalidParameterException
type InvalidParameterException struct {
	actual string
	expect string
}

func (e *InvalidParameterException) Error() string {
	return fmt.Sprintf("expect: [%s] actual: [%s]", e.expect, e.actual)
}
