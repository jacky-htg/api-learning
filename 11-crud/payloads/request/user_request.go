package request

import (
	"github.com/jacky-htg/go-services/models"
)

//NewUserRequest : format json request for new user
type NewUserRequest struct {
	Username   string `json:"username"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	RePassword string `json:"re_password"`
	IsActive   bool   `json:"is_active"`
}

//Transform NewUserRequest to User
func (u *NewUserRequest) Transform() *models.User {
	var user models.User
	user.Username = u.Username
	user.Email = u.Email
	user.Password = u.Password
	user.IsActive = u.IsActive

	return &user
}

//UserRequest : format json request for user
type UserRequest struct {
	ID         uint64 `json:"id,omitempty"`
	Username   string `json:"username,omitempty"`
	Email      string `json:"email,omitempty"`
	Password   string `json:"password,omitempty"`
	RePassword string `json:"re_password,omitempty"`
	IsActive   bool   `json:"is_active,omitempty"`
}

//Transform NewUserRequest to User
func (u *UserRequest) Transform(user *models.User) *models.User {
	if u.ID == user.ID {
		if len(u.Username) > 0 {
			user.Username = u.Username
		}

		if len(u.Email) > 0 {
			user.Email = u.Email
		}

		if len(u.Password) > 0 {
			user.Password = u.Password
		}

		user.IsActive = u.IsActive
	}
	return user
}
