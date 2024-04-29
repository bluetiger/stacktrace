package stacktrace

import (
	"errors"
	"testing"
)

func Test_main(t *testing.T) {
	Trace(errors.New("test_err"))
}
