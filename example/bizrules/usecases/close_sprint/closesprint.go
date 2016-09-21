package closesprint

import (
	"github.com/moul/cleanarch"
	"github.com/moul/cleanarch/example/bizrules/gateways"
	"github.com/moul/cleanarch/example/bizrules/usecases/close_sprint/io"
)

type UseCase struct {
	cleanarch.UseCase

	gw   gateways.Sprints
	resp closesprintio.ResponseBuilder
}

func New(gw gateways.Sprints, resp closesprintio.ResponseBuilder) UseCase {
	return UseCase{
		gw:   gw,
		resp: resp,
	}
}

func (uc *UseCase) Execute(req closesprintio.Request) (closesprintio.Response, error) {
	sprint, err := uc.gw.Find(req.GetSprintID())
	if err != nil {
		return nil, err
	}

	if err := sprint.Close(); err != nil {
		return nil, err
	}

	if err := uc.gw.Update(sprint); err != nil {
		return nil, err
	}

	return uc.resp.
		Create().
		WithAverageClosedIssues(uc.gw.FindAverageClosedIssues()).
		WithClosedIssuesCount(sprint.GetIssuesCount()).
		WithSprintID(sprint.GetID()).
		Build()
}
