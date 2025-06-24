package services

type UserService struct {
	Repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{
		Repo: repo,
	}
}

func (s *UserService) GetUsers() ([]*model.User, error) {
	users, err := s.Repo.GetAllUsers()
	if err != nil {
		return nil, err
	}
	return users, nil
}
