package usecases

import (
	"time"

	"github.com/moul/cleanarch/example/bizrules/entities"
)

var SprintStub1 = entities.NewSprint()
var SprintStub2 = entities.NewSprint()
var IssueStub1 = entities.NewIssue()
var IssueStub2 = entities.NewIssue()

func init() {
	IssueStub1.SetID(10)
	IssueStub1.SetTitle("Issue 1")
	IssueStub1.SetDescription("Description of the first issue")

	IssueStub2.SetID(20)
	if err := IssueStub2.SetDone(); err != nil {
		panic(err)
	}
	IssueStub2.SetDoneAt(time.Unix(1234567890, 0))
	IssueStub1.SetTitle("Issue 2")
	IssueStub1.SetDescription("Description of the second issue")

	SprintStub1.SetID(42)
	if err := SprintStub1.Close(); err != nil {
		panic(err)
	}
	SprintStub1.AddIssue(IssueStub1)
	SprintStub1.AddIssue(IssueStub2)

	SprintStub2.SetID(43)
	SprintStub2.AddIssue(IssueStub1)
	SprintStub2.AddIssue(IssueStub2)
}
