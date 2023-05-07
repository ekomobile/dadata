package client

import "fmt"

// ResponseError indicates an HTTP non-200 response code.
type ResponseError struct {
	Status     string // e.g. "200 OK"
	StatusCode int    // e.g. 200
}

func (e *ResponseError) Error() string {
	return fmt.Sprintf("HTTP response: %s", e.Status)
}
