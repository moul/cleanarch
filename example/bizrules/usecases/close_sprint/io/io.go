package closesprintio

type Response interface {
	// cleanarch.UseCaseResponse

	GetAverageClosedIssues() float64
	GetClosedIssuesCount() int
	GetSprintID() int
}

type ResponseBuilder interface {
	Create() ResponseBuilder
	WithAverageClosedIssues(float64) ResponseBuilder
	WithClosedIssuesCount(int) ResponseBuilder
	WithSprintID(int) ResponseBuilder
	Build() (Response, error)
}

type Request interface {
	// cleanarch.UseCaseRequest
	GetSprintID() int
}

type RequestBuilder interface {
	// cleanarch.UseCaseRequestBuilder

	Create() RequestBuilder
	WithSprintID(int) RequestBuilder
	Build() Request
}
