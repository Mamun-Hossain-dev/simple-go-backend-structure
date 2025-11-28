package user

type User struct {
	ID        int    `db:"id" json:"id"`
	FirstName string `db:"first_name" json:"firstName"`
	LastName  string `db:"last_name" json:"lastName"`
	Email     string `db:"email" json:"email"`
	Password  string `db:"password" json:"password"`
	IsAdmin   bool   `db:"is_admin" json:"isAdmin"`
}
type CreateUser struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type UserResponse struct {
	Message string `json:"message"`
	Data    User   `json:"data"`
}

type LoggedUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
