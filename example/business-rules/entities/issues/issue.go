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

/* generic setters */
// SetID sets the ID of the issue.
func (i *Issue) SetID(val int) { i.id = val }

// SetStatus sets the status of the issue.
func (i *Issue) SetStatus(val string) { i.status = val }

// SetTitle sets the title of the issue.
func (i *Issue) SetTitle(val string) { i.title = val }

// SetDescription sets the description of the issue.
func (i *Issue) SetDescription(val string) { i.description = val }

// SetCreatedAt sets the creation date of the issue.
func (i *Issue) SetCreatedAt(val time.Time) { i.createdAt = val }

// SetDoneAt sets the finish date of the issue.
func (i *Issue) SetDoneAt(val time.Time) { i.doneAt = val }

// SetClosedAt sets the closing date of the issue.
func (i *Issue) SetClosedAt(val time.Time) { i.closedAt = val }

/* generic getters */

// GetID returns the ID of the issue.
func (i *Issue) GetID() int { return i.id }

// GetStatus returns the status of the issue.
func (i *Issue) GetStatus() string { return i.status }

// GetTitle returns the title of the issue.
func (i *Issue) GetTitle() string { return i.title }

// GetDescription returns the description of the issue.
func (i *Issue) GetDescription() string { return i.description }

// GetCreatedAt returns the creation date of the issue.
func (i *Issue) GetCreatedAt() time.Time { return i.createdAt }

// GetDoneAt returns the finish date of the issue.
func (i *Issue) GetDoneAt() time.Time { return i.doneAt }

// GetClosedAt returns the closing date of the issue.
func (i *Issue) GetClosedAt() time.Time { return i.closedAt }

/* other methods */

// IsDone returns true if the issue status is "DONE".
func (i *Issue) IsDone() bool { return i.GetStatus() == Done }

// IsClosed returns true if the issue status is "CLOSED".
func (i *Issue) IsClosed() bool { return i.GetStatus() == Closed }

// SetDone sets the issue status to "DONE"
func (i *Issue) SetDone() error {
	if i.IsDone() {
		return &AlreadyDoneError{}
	}

	i.doneAt = time.Now()
	i.status = Done
	return nil
}

// Close closes an open issue
func (i *Issue) Close() error {
	if i.IsClosed() {
		return &AlreadyClosedError{}
	}

	i.closedAt = time.Now()
	i.status = Closed
	return nil
}
