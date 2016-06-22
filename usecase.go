package cleanarch

type UseCase interface {
	Execute(UseCaseRequest) (UseCaseResponse, error)
}
