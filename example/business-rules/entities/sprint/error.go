package sprint

import "fmt"

// SprintAlreadyClosedError is raised when trying to close an already closed sprint.
type AlreadyClosedError struct{}

func (f AlreadyClosedError) Error() string {
	return fmt.Sprintf("Sprint already closed")
}

// SprintNotFoundError is raised when trying to close an not found sprint.
type NotFoundError struct{}

func (f NotFoundError) Error() string {
	return fmt.Sprintf("Sprint not found")
}
