package database

type User struct {
	ID        int    `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	IsAdmin   bool   `json:"isAdmin"`
}

type UserResponse struct {
	Message string `json:"message"`
	Data    User   `json:"data"`
}

type LoggedUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

var users []User

func (u User) StoreUser() User {
	if u.ID != 0 {
		return u
	}

	u.ID = len(users) + 1
	users = append(users, u)
	return u
}

func Find(email, pass string) *User {
	for i := range users {
		if users[i].Email == email && users[i].Password == pass {
			return &users[i]
		}
	}
	return nil
}
