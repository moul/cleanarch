package ping

import "github.com/moul/cleanarch/example/bizrules/usecases/ping/io"

type UseCase struct {
	// cleanarch.UseCase

	resp pingio.ResponseAssembler
}

func New(resp pingio.ResponseAssembler) UseCase {
	return UseCase{resp: resp}
}

func (uc *UseCase) Execute(req pingio.Request) (pingio.Response, error) {
	return uc.resp.Write("pong")
}
