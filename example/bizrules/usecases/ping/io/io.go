package pingio

type Response interface {
	// cleanarch.UseCaseResponse
	GetPong() string
}

type ResponseAssembler interface {
	Write(string) (Response, error)
}

type Request interface {
	// cleanarch.UseCaseRequest
}

type RequestBuilder interface {
	// cleanarch.UseCaseRequestBuilder
	Create() RequestBuilder
	Build() Request
}
