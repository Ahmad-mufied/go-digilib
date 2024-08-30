package handler

import (
	"github.com/Ahmad-mufied/go-digilib/data"
	"github.com/go-playground/validator/v10"
)

var repo *data.Models
var validate *validator.Validate

func InitHandler(m *data.Models, v *validator.Validate) {
	repo = m
	validate = v
}
