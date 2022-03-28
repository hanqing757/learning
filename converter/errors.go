package converter

import (
	"fmt"
	"strings"
)

type Error struct {
	Errors []string
}

func (e *Error) Error() string {
	points := make([]string, len(e.Errors))
	for i, s := range e.Errors {
		points[i] = fmt.Sprintf("* %s", s)
	}
	return fmt.Sprintf("%d errors in process:\n%s\n",
		len(points),
		strings.Join(points, "\n"),
	)
}

func (e *Error) AppendError(err error) {
	switch realErr := err.(type) {
	case *Error:
		e.Errors = append(e.Errors, realErr.Errors...)
	default:
		e.Errors = append(e.Errors, err.Error())
	}
}

func (e *Error) IsNil() bool {
	return e.Errors == nil || len(e.Errors) == 0
}
