package utils

import (
	"fmt"
	"time"

	"github.com/go-playground/validator"
)

// parse validation errors while saving a risk,
// these validation rules are defined in models,
// currently there are 3 types of rules:
//  1. required
//  2. oneof
//  3. max
func ParseValidationErr(err *error) *[]string {
	errors := (*err).(validator.ValidationErrors)
	var errMsgs = make([]string, len(errors))

	for i, error := range errors {
		switch error.Tag() {
		case "required":
			errMsgs[i] = fmt.Sprintf("%s is required", error.Namespace())
		case "oneof":
			errMsgs[i] = fmt.Sprintf("%s value must be one of [%s]", error.Namespace(), error.Param())
		case "max":
			errMsgs[i] = fmt.Sprintf("%s exceeds the max length of %s", error.Namespace(), error.Param())
		}
	}

	return &errMsgs
}

// simplify format of log output, should use other logger libraries that provide
// more features in the future, e.g. zap, logrus
func FormatLog(methodName string, messages interface{}) string {
	return fmt.Sprintf("%v %s - %v", time.Now(), methodName, messages)
}
