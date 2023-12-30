package domain

import (
	"github.com/dhimweray222/users/model/web"
	"github.com/google/uuid"
)

type User struct {
	ID         uuid.UUID `json:"id"`
	Name       string    `json:"name"`
	Address    string    `json:"address"`
	Email      string    `json:"email"`
	Password   string    `json:"password"`
	CreditType string    `json:"credit_type"`
	CCNumber   string    `json:"number"`
	CCName     string    `json:"cc_name"`
	CCExpired  string    `json:"expired"`
	CCV        string    `json:"ccv"`
	Photos     []string  `json:"photos"`
}

func (user *User) GenerateIdKey() {
	id := uuid.New()
	user.ID = id
}
func ToUserResponse(user User) web.UserRepsonse {
	return web.UserRepsonse{
		UserId: user.ID,
	}
}
