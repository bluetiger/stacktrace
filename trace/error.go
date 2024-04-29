package stacktrace

import "runtime/debug"

func Trace(err error) {
	debug.PrintStack()
}
