package issueimpl

import (
	"github.com/moul/go-clean-architecture/business-rules/entities/issues"
	"github.com/moul/go-clean-architecture/business-rules/entities/sprint"
)

type IssueImpl struct {
	issue.Issue

	sprint sprint.Sprint
}
