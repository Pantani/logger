package logger

import (
	"github.com/Pantani/errors"
	log "github.com/sirupsen/logrus"
)

type errMessage struct {
	*message
	err error
}

// Error log an error message
func Error(args ...interface{}) {
	if len(args) == 0 {
		Panic("call to logger.Error with no arguments")
	}
	e := getError(args...)
	log.WithFields(e.params).Error(e.err)
}

// Fatal log a fatal message
func Fatal(args ...interface{}) {
	if len(args) == 0 {
		Panic("call to logger.Fatal with no arguments")
	}
	e := getError(args...)
	log.WithFields(e.params).Fatal(e.err)
}

// Panic trigger a app panic with message
func Panic(args ...interface{}) {
	if len(args) == 0 {
		Panic("call to logger.Panic with no arguments")
	}
	e := getError(args...)
	log.WithFields(e.params).Panic(e.err)
}

// getError parse arguments and creates the error message object
func getError(args ...interface{}) *errMessage {
	msg := getMessage(args...)
	err := &errMessage{message: msg}
	for _, arg := range args {
		switch arg := arg.(type) {
		case *errors.Error:
			err.err = arg
		case error:
			err.err = errors.E(arg)
		case nil:
			continue
		default:
			continue
		}
	}
	if err.err == nil {
		err.err = errors.E(msg.message, msg.params)
	}
	return err
}
