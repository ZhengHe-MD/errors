package errors

import (
	"runtime"
	"strings"
)

func FuncName() string {
	return FuncNameWithSkip(2)
}

func FuncNameWithSkip(skip int) string {
	pc, _, _, _ := runtime.Caller(skip)
	f := runtime.FuncForPC(pc)
	parts := strings.Split(f.Name(), ".")
	return parts[len(parts)-1]
}