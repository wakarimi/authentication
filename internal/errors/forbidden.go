package errors

import "fmt"

type Forbidden struct {
	Message string
}

func (e Forbidden) Error() string {
	return fmt.Sprintf("%s", e.Message)
}
