package event

type Service interface {
	FindAll() ([]Movie, error)
	FindOne(id int64) (*Movie, error)
	CreateMovie(name string, director string, year int64) (*Movie, error)
	UpdateName(id int64, name string, director string, year int64) (*Movie, error)
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

func (s *service) FindOne(id int64) (*Movie, error) {
	return (*s.repo).FindOne(id)
}

func (s *service) CreateMovie(name string, director string, year int64) (*Movie, error) {
	return (*s.repo).CreateMovie(name, director, year)
}

func (s *service) UpdateName(id int64, name string, director string, year int64) (*Movie, error) {
	return (*s.repo).UpdateName(id, name, director, year)
}
