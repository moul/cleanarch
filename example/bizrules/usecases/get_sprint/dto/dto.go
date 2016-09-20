package getsprintdto

import (
	"time"

	"github.com/moul/cleanarch/example/bizrules/entities"
	"github.com/moul/cleanarch/example/bizrules/usecases/get_sprint/io"
)

/* Request */

type Request struct {
	getsprintio.Request

	id int
}

func (r Request) GetID() int { return r.id }

/* RequestBuilder */

type RequestBuilder struct {
	getsprintio.RequestBuilder

	request *Request
}

func (b RequestBuilder) Create() getsprintio.RequestBuilder {
	b.request = &Request{}
	return b
}

func (b RequestBuilder) WithSprintID(id int) getsprintio.RequestBuilder {
	b.request.id = id
	return b
}

func (b RequestBuilder) Build() getsprintio.Request { return b.request }

/* Response */

type Response struct {
	getsprintio.Response

	createdAt         time.Time
	effectiveClosedAt time.Time
	expectedClosedAt  time.Time
	id                int
	status            string
}

func (r Response) GetCreatedAt() time.Time         { return r.createdAt }
func (r Response) GetEffectiveClosedAt() time.Time { return r.effectiveClosedAt }
func (r Response) GetExpectedClosedAt() time.Time  { return r.expectedClosedAt }
func (r Response) GetID() int                      { return r.id }
func (r Response) GetStatus() string               { return r.status }

/* ResponseAssembler */

type ResponseAssembler struct {
	getsprintio.ResponseAssembler
}

func (a ResponseAssembler) Write(sprint *entities.Sprint) (getsprintio.Response, error) {
	resp := Response{
		createdAt:         sprint.GetCreatedAt(),
		effectiveClosedAt: sprint.GetEffectiveClosedAt(),
		expectedClosedAt:  sprint.GetExpectedClosedAt(),
		id:                sprint.GetID(),
		status:            sprint.GetStatus(),
	}
	return resp, nil
}
