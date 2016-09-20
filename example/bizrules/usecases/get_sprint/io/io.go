package getsprintio

import (
	"time"

	"github.com/moul/cleanarch"
	"github.com/moul/cleanarch/example/bizrules/entities"
)

type Response interface {
	cleanarch.UseCaseResponse

	GetCreatedAt() time.Time
	GetEffectiveClosedAt() time.Time
	GetExpectedClosedAt() time.Time
	GetID() int
	GetStatus() string
}

type ResponseAssembler interface {
	Write(*entities.Sprint) (Response, error)
}

type Request interface {
	cleanarch.UseCaseRequest

	GetID() int
}

type RequestBuilder interface {
	cleanarch.UseCaseRequestBuilder
	Create() RequestBuilder
	WithSprintID(int) RequestBuilder
	Build() Request
}
