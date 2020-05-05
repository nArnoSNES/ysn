package errors

import "fmt"

type NoMatch struct{}

func (e *NoMatch) Error() string {
	return fmt.Sprintf("No match found")
}

type EndOfSource struct{}

func (e *EndOfSource) Error() string {
	return fmt.Sprintf("End of source")
}
