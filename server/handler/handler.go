package handler

import "github.com/Ahmad-mufied/go-digilib/data"

var entity *data.Models

func InitHandler(m *data.Models) {
	entity = m
}
