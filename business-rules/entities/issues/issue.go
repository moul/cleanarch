package issue

import "time"

// Issue represents an issue.
type Issue struct {
	id          int
	status      string
	title       string
	description string
	createdAt   time.Time
	doneAt      time.Time
	closedAt    time.Time
}

// New returns an instanciated instance of Issue.
func New() *Issue {
	return &Issue{
		createdAt: time.Now(),
	}
}

// GetID returns the ID of the issue.
func (i *Issue) GetID() int { return i.id }

// GetStatus returns the status of the issue.
func (i *Issue) GetStatus() string { return i.status }

// GetStatus returns the title of the issue.
func (i *Issue) GetTitle() string { return i.title }

// GetStatus returns the description of the issue.
func (i *Issue) GetDescription() string { return i.description }

// GetStatus returns the creation date of the issue.
func (i *Issue) GetCreatedAt() time.Time { return i.createdAt }

// GetStatus returns the finish date of the issue.
func (i *Issue) GetDoneAt() time.Time { return i.doneAt }

// GetStatus returns the closing date of the issue.
func (i *Issue) GetClosedAt() time.Time { return i.closedAt }

// IsDone returns true if the issue status is "DONE".
func (i *Issue) IsDone() bool { return i.GetStatus() == Done }

// IsClosed returns true if the issue status is "CLOSED".
func (i *Issue) IsClosed() bool { return i.GetStatus() == Closed }

// Close closes an open issue
func (i *Issue) Close() error {
	if i.IsClosed() {
		return &AlreadyClosedError{}
	}

	i.closedAt = time.Now()
	i.status = Closed
	return nil
}
