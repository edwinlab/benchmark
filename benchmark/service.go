package benchmark

type Service interface {
	UpdateRandom()
}

type service struct {
	repo Repository
}

func (s *service) UpdateRandom() {
	s.repo.UpdateRow(s.repo.GetRandomId())
}

func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}
