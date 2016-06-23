package issue

import "fmt"

// IssueAlreadyClosedError is raised when trying to close an already closed issue.
type AlreadyClosedError struct{}

func (f AlreadyClosedError) Error() string {
	return fmt.Sprintf("Issue already closed")
}

// IssueAlreadyDoneError is raised when trying to close an already done issue.
type AlreadyDoneError struct{}

func (f AlreadyDoneError) Error() string {
	return fmt.Sprintf("Issue already done")
}
