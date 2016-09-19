package sprintuc

import (
	"time"

	"github.com/moul/go-clean-architecture/business-rules/entities/sprint"
)
import "github.com/moul/go-clean-architecture/business-rules/entities/issues"

var sprintStub1 = sprint.New()
var sprintStub2 = sprint.New()
var issueStub1 = issue.New()
var issueStub2 = issue.New()

func init() {
	issueStub1.SetID(10)
	issueStub1.SetTitle("Issue 1")
	issueStub1.SetDescription("Description of the first issue")

	issueStub2.SetID(20)
	issueStub2.SetDone()
	issueStub2.SetDoneAt(time.Unix(1234567890, 0))
	issueStub1.SetTitle("Issue 2")
	issueStub1.SetDescription("Description of the second issue")

	sprintStub1.SetID(42)
	sprintStub1.Close()
	sprintStub1.AddIssue(issueStub1)
	sprintStub1.AddIssue(issueStub2)

	sprintStub2.SetID(43)
	sprintStub2.AddIssue(issueStub1)
	sprintStub2.AddIssue(issueStub2)
}
