package movie

type Service interface {
	FindAll() ([]Movie, error)
	FindById(id int64) (*Movie, error)
	CreateMovie(name string, director string, year int64) (*Movie, error)
	UpdateMovie(id int64, name string, director string, year int64) (*Movie, error)
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

func (s *service) CreateMovie(name string, director string, year int64) (*Movie, error) {
	return (*s.repo).CreateMovie(name, director, year)
}

func (s *service) UpdateMovie(id int64, name string, director string, year int64) (*Movie, error) {
	return (*s.repo).UpdateMovie(id, name, director, year)
}
