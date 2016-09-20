package addsprint

import (
	"github.com/moul/cleanarch"
	"github.com/moul/cleanarch/example/bizrules/gateways"
	"github.com/moul/cleanarch/example/bizrules/usecases/add_sprint/io"
)

type UseCase struct {
	cleanarch.UseCase

	gw   gateways.Sprints
	resp addsprintio.ResponseAssembler
}

func New() UseCase {
	return UseCase{}
}

func (uc *UseCase) SetSprintsGateway(gw gateways.Sprints) {
	uc.gw = gw
}

func (uc *UseCase) SetResponseAssembler(resp addsprintio.ResponseAssembler) {
	uc.resp = resp
}

func (uc *UseCase) Execute(req addsprintio.Request) (addsprintio.Response, error) {
	newSprint, err := uc.gw.New()
	if err != nil {
		return nil, err
	}

	return uc.resp.Write(newSprint)
}
