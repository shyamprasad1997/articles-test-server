package utils

import (
	"fmt"
	"runtime"
	"strings"

	"github.com/pkg/errors"
)

var goPath string

// ErrorsWrap add caller/line and set errors.Wrap.
func ErrorsWrap(err error, s string) error {
	path := strings.Trim(getCallerString(), goPath)
	return errors.Wrap(err, path+s)
}

// GetCallerString get caller string.
func getCallerString() string {
	_, f, l, _ := runtime.Caller(2)
	return fmt.Sprintf("%v:%v:", f, l)
}
