package app

type ProductService struct {
	Persistence ProductPersistenceInterface
}

func NewProductService(p ProductPersistenceInterface) *ProductService {
	return &ProductService{Persistence: p}
}

func (s *ProductService) Get(id string) (ProductInterface, error) {
	product, err := s.Persistence.Get(id)

	if err != nil {
		return nil, err
	}

	return product, nil
}
