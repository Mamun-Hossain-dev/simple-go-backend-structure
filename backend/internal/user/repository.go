package user

type UserRepository interface {
	StoreUser(u User) User
	Find(email, pass string) *User
}

type userRepo struct{}

func NewUserRepository() UserRepository {
	return &userRepo{}
}

func (r *userRepo) StoreUser(u User) User {
	if u.ID != 0 {
		return u
	}

	u.ID = len(Users) + 1
	Users = append(Users, u)
	return u
}

func (r *userRepo) Find(email, pass string) *User {
	for i := range Users {
		if Users[i].Email == email && Users[i].Password == pass {
			return &Users[i]
		}
	}
	return nil
}
