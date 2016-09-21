package addsprintio

import (
	"time"

	"github.com/moul/cleanarch/example/bizrules/entities"
)

type Response interface {
	// cleanarch.UseCaseResponse

	GetCreatedAt() time.Time
	GetID() int
}
type ResponseAssembler interface {
	Write(*entities.Sprint) (Response, error)
}

type Request interface {
	// cleanarch.UseCaseRequest
}
type RequestBuilder interface {
	// cleanarch.UseCaseRequestBuilder

	Create() RequestBuilder
	Build() Request
}
