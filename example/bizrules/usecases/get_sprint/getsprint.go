package getsprint

import (
	"github.com/moul/cleanarch/example/bizrules/gateways"
	"github.com/moul/cleanarch/example/bizrules/usecases/get_sprint/io"
)

type UseCase struct {
	// cleanarch.UseCase

	gw   gateways.Sprints
	resp getsprintio.ResponseAssembler
}

func New(gw gateways.Sprints, resp getsprintio.ResponseAssembler) UseCase {
	return UseCase{
		gw:   gw,
		resp: resp,
	}
}

func (uc *UseCase) Execute(req getsprintio.Request) (getsprintio.Response, error) {
	sprint, err := uc.gw.Find(req.GetID())
	if err != nil {
		return nil, err
	}

	return uc.resp.Write(sprint)
}
