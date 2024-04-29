package stacktrace

import (
	"runtime"
	"strconv"
)

type StackErrors struct {
	frame     *runtime.Frames
	baseError error
}

func (s *StackErrors) Error() string {
	msg := s.baseError.Error()
	fr, ok := s.frame.Next()
	for ok {
		fr, ok = s.frame.Next()
		msg += fr.Function + "\n"
		msg += "\t" + fr.File + ":" + strconv.Itoa(fr.Line) + "\n"
	}
	return msg
}

func StackTrace(err error) error {
	var (
		pc [100]uintptr
	)

	return &StackErrors{
		frame:     runtime.CallersFrames(pc[:runtime.Callers(1, pc[:])]),
		baseError: err,
	}
}
