package converter

import (
	"github.com/Ahmad-mufied/go-digilib/data"
	"github.com/Ahmad-mufied/go-digilib/model"
)

func UserToGetUserResponse(user *data.User, walletId uint) *model.GetUserResponse {
	return &model.GetUserResponse{
		ID:        user.ID,
		WalletID:  walletId,
		FullName:  user.FullName,
		Username:  user.Username,
		Email:     user.Email,
		Status:    user.Status,
		Role:      user.Role,
		BookCount: user.BookCount,
	}
}
