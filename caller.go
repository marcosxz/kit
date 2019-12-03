package kit

import (
	"runtime"
	"strings"
)

var callerPkgName string

func init() {
	pcs := make([]uintptr, 2)
	runtime.Callers(0, pcs)
	callerPkgName = packageName(runtime.FuncForPC(pcs[1]).Name())
}

func packageName(fn string) string {
	if idx := strings.LastIndex(fn, "/"); idx > -1 {
		return fn[:idx]
	}
	return fn
}

func CallerFrame(skip int) (*runtime.Frame, bool) {
	pcs := make([]uintptr, 25)
	depth := runtime.Callers(skip, pcs)
	frames := runtime.CallersFrames(pcs[:depth])
	for frame, again := frames.Next(); again; frame, again = frames.Next() {
		if packageName(frame.Function) != callerPkgName {
			return &frame, true
		}
	}
	return nil, false
}

func Caller(skip int) (pc uintptr, file string, line int, ok bool) {
	return runtime.Caller(skip)
}
