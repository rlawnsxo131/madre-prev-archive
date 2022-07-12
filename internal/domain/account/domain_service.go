package account

type (
	DomainService interface{}
	domainService struct{}
)

func NewDomainService() DomainService {
	return &domainService{}
}
