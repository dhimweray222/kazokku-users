package web

import "github.com/google/uuid"

type UserRepsonse struct {
	UserId uuid.UUID `json:"user_id"`
}

type UserRequest struct {
	Name       string   `json:"name"`
	Address    string   `json:"address"`
	Email      string   `json:"email"`
	Password   string   `json:"password"`
	CreditType string   `form:"credit_type"`
	CCNumber   string   `form:"number"`
	CCName     string   `form:"cc_name"`
	CCExpired  string   `form:"expired"`
	CCV        string   `form:"ccv"`
	Photos     []string `json:"photos"`
}
