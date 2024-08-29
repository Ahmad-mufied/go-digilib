package handler

import (
	"github.com/Ahmad-mufied/go-digilib/config"
	"github.com/Ahmad-mufied/go-digilib/constants"
	"github.com/Ahmad-mufied/go-digilib/data"
	"github.com/Ahmad-mufied/go-digilib/model"
	"github.com/Ahmad-mufied/go-digilib/model/converter"
	"github.com/Ahmad-mufied/go-digilib/server/middleware"
	"github.com/Ahmad-mufied/go-digilib/utils"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

func Register(c echo.Context) error {
	// Get User Payload
	userPayload := new(model.UserRegisterRequest)
	err := c.Bind(&userPayload)
	if err != nil {
		return utils.HandleError(c, constants.ErrBadRequest)
	}

	// Check if UserRegisterRequest payload the is empty
	if userPayload.FullName == "" || userPayload.Username == "" || userPayload.Email == "" || userPayload.Password == "" {
		return utils.HandleError(c, constants.ErrBadRequest)
	}

	// Check if email is already registered
	isExist, err := repo.User.CheckEmail(userPayload.Email)
	if err != nil {
		return utils.HandleError(c, constants.ErrInternalServerError, err.Error())
	}

	if isExist {
		return utils.HandleError(c, constants.ErrConflict, "Email already registered")
	}

	bcryptCost := config.Viper.GetInt("BCRYPT_SALT")
	if bcryptCost == 0 {
		bcryptCost = bcrypt.DefaultCost
	}

	// Hash Password
	generateFromPassword, err := bcrypt.GenerateFromPassword([]byte(userPayload.Password), bcryptCost)
	if err != nil {
		return utils.HandleError(c, constants.ErrInternalServerError)
	}

	user := new(data.User)
	user.FullName = userPayload.FullName
	user.Username = userPayload.Username
	user.Email = userPayload.Email
	user.Password = string(generateFromPassword)
	user.Status = constants.UserStatusActive
	user.Role = constants.UserRoleReader

	userId, err := repo.User.CreateUser(user)
	if err != nil {
		return utils.HandleError(c, constants.ErrInternalServerError, err.Error())
	}

	user.ID = userId

	userResponse := converter.UserToGetUserResponse(user)

	return c.JSON(http.StatusCreated, model.JSONResponse{
		Status:  constants.ResponseStatusSuccess,
		Message: "Success Register",
		Data:    userResponse,
	})
}

func Login(c echo.Context) error {
	// Get User Payload
	userPayload := new(model.UserLoginRequest)
	err := c.Bind(&userPayload)
	if err != nil {
		return utils.HandleError(c, constants.ErrBadRequest)
	}

	// Check if UserLoginRequest payload the is empty
	if userPayload.Email == "" || userPayload.Password == "" {
		return utils.HandleError(c, constants.ErrBadRequest, "Email or password is empty")
	}

	// Get password by email
	user, err := repo.User.GetUserByEmail(userPayload.Email)
	if err != nil {
		return utils.HandleError(c, constants.ErrInternalServerError, err.Error())
	}

	if user == nil {
		return utils.HandleError(c, constants.ErrUnauthorized, "User not found")
	}

	// Compare password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userPayload.Password))
	if err != nil {
		return utils.HandleError(c, constants.ErrUnauthorized, "Email or password is wrong")
	}

	token, err := middleware.GenerateToken(user.ID, user.Email, string(user.Role))
	if err != nil {
		return utils.HandleError(c, constants.ErrInternalServerError, "Failed to generate token")
	}

	userDetail := converter.UserToGetUserResponse(user)
	response := model.UserLoginResponse{
		Token:      token,
		UserDetail: userDetail,
	}

	return c.JSON(http.StatusOK, model.JSONResponse{
		Status:  constants.ResponseStatusSuccess,
		Message: "Success Login",
		Data:    response,
	})
}
