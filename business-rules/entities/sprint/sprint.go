package sprint

import (
	"time"

	"github.com/moul/go-clean-architecture/business-rules/entities/issues"
)

// Sprint represents an sprint.
type Sprint struct {
	id                int
	status            string
	createdAt         time.Time
	expectedClosedAt  time.Time
	effectiveClosedAt time.Time
	issues            []*issue.Issue
}

// New returns an instanciated instance of Sprint.
func New() *Sprint {
	return &Sprint{
		issues:    make([]*issue.Issue, 0),
		status:    Open,
		createdAt: time.Now(),
	}
}

/* Setters */

// SetID sets the ID of the sprint.
func (i *Sprint) SetID(val int) { i.id = val }

// SetStatus sets the Status of the sprint.
func (i *Sprint) SetStatus(val string) { i.status = val }

// SetCreatedAt sets the CreatedAt of the sprint.
func (i *Sprint) SetCreatedAt(val time.Time) { i.createdAt = val }

// SetExpectedClosedAt sets the ExpectedClosedAt of the sprint.
func (i *Sprint) SetExpectedClosedAt(val time.Time) { i.expectedClosedAt = val }

// SetEffectiveClosedAt sets the EffectiveClosedAt of the sprint.
func (i *Sprint) SetEffectiveClosedAt(val time.Time) { i.effectiveClosedAt = val }

/* Getters */

// GetID returns the ID of the sprint.
func (i *Sprint) GetID() int { return i.id }

// GetStatus returns the status of the sprint.
func (i *Sprint) GetStatus() string { return i.status }

// GetCreatedAt returns the creation date of the sprint.
func (i *Sprint) GetCreatedAt() time.Time { return i.createdAt }

// GetExpectedClosedAt returns the finish date of the sprint.
func (i *Sprint) GetExpectedClosedAt() time.Time { return i.expectedClosedAt }

// GetEffectiveClosedAt returns the finish date of the sprint.
func (i *Sprint) GetEffectiveClosedAt() time.Time { return i.effectiveClosedAt }

/* ---- */

// AddIssue adds an issue to the sprint.
func (i *Sprint) AddIssue(issue *issue.Issue) {
	i.issues = append(i.issues, issue)
}

// GetIssues returns the issues of the sprint.
func (i *Sprint) GetIssues() []*issue.Issue {
	return i.issues
}

// GetIssuesCount returns the count of issues in the sprint.
func (i *Sprint) GetIssuesCount() int {
	return len(i.issues)
}

// IsClosed returns true if the sprint status is "CLOSED".
func (i *Sprint) IsClosed() bool { return i.GetStatus() == Closed }

// Close closes an open sprint
func (i *Sprint) Close() error {
	if i.IsClosed() {
		return &AlreadyClosedError{}
	}

	for idx := len(i.issues) - 1; idx >= 0; idx-- {
		issue := i.issues[idx]

		if issue.IsDone() {
			if err := issue.Close(); err != nil {
				return err
			}
		} else {
			i.issues = append(i.issues[:idx], i.issues[idx+1:]...)
		}
	}

	i.effectiveClosedAt = time.Now()
	i.status = Closed
	return nil
}
