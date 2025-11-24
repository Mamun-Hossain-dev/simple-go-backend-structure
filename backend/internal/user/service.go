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
		IsAdmin:   false,
	}

	createdUser := s.repo.StoreUser(u)

	return createdUser, nil
}

func (s *userService) LoginUser(lu LoggedUser) (*User, error) {
	user := s.repo.Find(lu.Email, lu.Password)
	if user == nil {
		return nil, fmt.Errorf("invalid email or password")
	}
	return user, nil
}
