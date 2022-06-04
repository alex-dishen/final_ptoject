package movie

type Service interface {
	FindAll() ([]Movie, error)
	FindById(id int64) (*Movie, error)
	CreateMovie(m *Movie) (*Movie, error)
	UpdateMovie(m *Movie) (*Movie, error)
	Delete(id int64) error
}

type service struct {
	repo *Repository
}

func NewService(r *Repository) Service {
	return &service{
		repo: r,
	}
}

func (s *service) FindAll() ([]Movie, error) {
	return (*s.repo).FindAll()
}

func (s *service) FindById(id int64) (*Movie, error) {
	return (*s.repo).FindById(id)
}

func (s *service) CreateMovie(m *Movie) (*Movie, error) {
	return (*s.repo).CreateMovie(m)
}

func (s *service) UpdateMovie(m *Movie) (*Movie, error) {
	return (*s.repo).UpdateMovie(m)
}

func (s *service) Delete(id int64) error {
	return (*s.repo).Delete(id)
}
