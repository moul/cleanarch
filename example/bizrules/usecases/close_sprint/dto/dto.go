package closesprintdto

import "github.com/moul/cleanarch/example/bizrules/usecases/close_sprint/io"

type RequestBuilder struct {
	closesprintio.RequestBuilder

	request *Request
}

type Request struct {
	closesprintio.Request

	id int
}

func (r Request) GetSprintID() int { return r.id }

func (b RequestBuilder) Create() RequestBuilder {
	b.request = &Request{}
	return b
}

func (b RequestBuilder) WithSprintID(id int) RequestBuilder {
	b.request.id = id
	return b
}

func (b RequestBuilder) Build() *Request { return b.request }

type Response struct {
	closesprintio.Response

	averageClosedIssues float64
	closedIssuesCount   int
	sprintID            int
}

func (r Response) GetAverageClosedIssues() float64 { return r.averageClosedIssues }
func (r Response) GetClosedIssuesCount() int       { return r.closedIssuesCount }
func (r Response) GetSprintID() int                { return r.sprintID }

type ResponseBuilder struct {
	closesprintio.ResponseBuilder

	response *Response
}

func (a ResponseBuilder) Create() closesprintio.ResponseBuilder {
	a.response = &Response{}
	return a
}

func (a ResponseBuilder) WithAverageClosedIssues(val float64) closesprintio.ResponseBuilder {
	a.response.averageClosedIssues = val
	return a
}

func (a ResponseBuilder) WithClosedIssuesCount(val int) closesprintio.ResponseBuilder {
	a.response.closedIssuesCount = val
	return a
}

func (a ResponseBuilder) WithSprintID(val int) closesprintio.ResponseBuilder {
	a.response.sprintID = val
	return a
}

func (a ResponseBuilder) Build() (closesprintio.Response, error) {
	return a.response, nil
}
