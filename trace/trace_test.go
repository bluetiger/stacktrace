package stacktrace

import (
	"errors"
	"fmt"
	"testing"
)

func Test_main(t *testing.T) {
	fmt.Println(StackTrace(errors.New("test_err")).Error())
}
