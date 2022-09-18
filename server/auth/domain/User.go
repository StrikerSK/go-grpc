package domain

import "github.com/StrikerSK/go-grpc/proto/auth"

type User struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func NewUser() *User {
	return &User{}
}

func (r *User) FromRegisterRequest(user *auth.RegisterRequest) *User {
	r.Email = user.Email
	r.Username = user.Username
	r.Password = user.Password
	return r
}
