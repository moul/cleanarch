package pingdto

import "github.com/moul/cleanarch/example/bizrules/usecases/ping/io"

type RequestBuilder struct {
	pingio.RequestBuilder
	request *Request
}

type Request struct{ pingio.Request }

type ResponseAssembler struct{ pingio.ResponseAssembler }
type Response struct{ pong string }

func (b RequestBuilder) Create() pingio.RequestBuilder {
	b.request = &Request{}
	return b
}

func (b RequestBuilder) Build() pingio.Request { return b.request }

func (r Response) GetPong() string { return r.pong }

func (a ResponseAssembler) Write(pong string) (pingio.Response, error) {
	return Response{pong: pong}, nil
}
