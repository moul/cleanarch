package issue

import "fmt"

// IssueAlreadyClosedError is raised when trying to close an already closed issue.
type AlreadyClosedError struct{}

func (f AlreadyClosedError) Error() string {
	return fmt.Sprintf("Issue already closed")
}
