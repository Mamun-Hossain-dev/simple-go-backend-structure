package user

import "fmt"

type UserService interface {
	RegisterUser(cu CreateUser) (User, error)
	LoginUser(lu LoggedUser) (*User, error)
}

type userService struct {
	repo UserRepository
}

func NewUserService(r UserRepository) UserService {
	return &userService{
		repo: r,
	}
}

func (s *userService) RegisterUser(cu CreateUser) (User, error) {
	u := User{
		FirstName: cu.FirstName,
		LastName:  cu.LastName,
		Email:     cu.Email,
		Password:  cu.Password,
	}

	createdUser, err := s.repo.StoreUser(u)
	if err != nil {
		return User{}, err
	}

	return createdUser, nil
}

func (s *userService) LoginUser(lu LoggedUser) (*User, error) {
	user, err := s.repo.Find(lu.Email, lu.Password)
	if err != nil {
		return nil, fmt.Errorf("invalid email or password")
	}
	return user, nil
}
