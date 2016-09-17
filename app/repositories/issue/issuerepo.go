package issuerepo

import "github.com/moul/go-clean-architecture/business-rules/gateways/issue"

type IssueRepository struct {
	issuegw.IssueGateway
}

/*func (r *IssueRepository) Find(int) (*issue.Issue, error) {
	return nil, fmt.Errorf("Not implemented")
}

func (r *IssueRepository) FindIssueToClose() (*issue.Issue, error) {
	return nil, fmt.Errorf("Not implemented")
}

func (r *IssueRepository) FindAverageClosedIssues() int {
	return 0
}

func (r *IssueRepository) Update(*issue.Issue) error {
	return fmt.Errorf("Not implemented")
}*/
