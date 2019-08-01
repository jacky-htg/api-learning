package response

import (
	"github.com/jacky-htg/go-services/models"
)

//UserResponse : format json response for user
type UserResponse struct {
	ID       uint64 `json:"id"`
	Username string `json:"username"`
	IsActive bool   `json:"is_active"`
}

//Transform from User model to User response
func (u *UserResponse) Transform(user models.User) {
	u.ID = user.ID
	u.Username = user.Username
	u.IsActive = user.IsActive
}
