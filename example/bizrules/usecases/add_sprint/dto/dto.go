package addsprintdto

import (
	"time"

	"github.com/moul/cleanarch/example/bizrules/entities"
	"github.com/moul/cleanarch/example/bizrules/usecases/add_sprint/io"
)

/* Requestbuilder */

type RequestBuilder struct {
	addsprintio.RequestBuilder

	request *Request
}

func (b RequestBuilder) Create() addsprintio.RequestBuilder {
	b.request = &Request{}
	return b
}

func (b RequestBuilder) Build() addsprintio.Request { return b.request }

/* Request */

type Request struct {
	addsprintio.Request

	id int
}

func (r Request) GetSprintID() int { return r.id }

/* Response */

type Response struct {
	addsprintio.Response

	id        int
	createdAt time.Time
}

func (r Response) GetCreatedAt() time.Time { return r.createdAt }
func (r Response) GetID() int              { return r.id }

/* ResponseAssembler */

type ResponseAssembler struct {
	addsprintio.ResponseAssembler
}

func (a ResponseAssembler) Write(sprint *entities.Sprint) (addsprintio.Response, error) {
	resp := Response{
		createdAt: sprint.GetCreatedAt(),
		id:        sprint.GetID(),
	}
	return resp, nil
}
