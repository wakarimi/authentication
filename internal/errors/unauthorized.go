package errors

import "fmt"

type Unauthorized struct {
	Message string
}

func (e Unauthorized) Error() string {
	return fmt.Sprintf("%s", e.Message)
}
