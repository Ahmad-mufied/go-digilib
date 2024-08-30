package model

import "github.com/Ahmad-mufied/go-digilib/constants"

// UserRegisterRequest POST /users/register
// request body -> { full_name, username, email, password}
type UserRegisterRequest struct {
	FullName string `json:"full_name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// GetUserResponse POST /users/register
// response body -> { id, full_name, username, email, status, role, book_count }
type GetUserResponse struct {
	ID        uint                      `json:"id"`
	WalletID  uint                      `json:"wallet_id"`
	FullName  string                    `json:"FullName"`
	Username  string                    `json:"username"`
	Email     string                    `json:"email"`
	Status    constants.UsersStatusEnum `json:"status"`
	Role      constants.UserRoleEnum    `json:"role"`
	BookCount int                       `json:"book_count"`
}

// UserLoginRequest POST /users/login
// request body -> { username, password }
type UserLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// UserLoginResponse POST /users/login
// response body -> { user_detail, token }
type UserLoginResponse struct {
	Token      string           `json:"token"`
	UserDetail *GetUserResponse `json:"user_detail"`
}
