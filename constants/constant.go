package constants

import (
	"github.com/Ahmad-mufied/go-digilib/utils"
	"net/http"
)

const (
	ResponseStatusFailed  = "Failed"
	ResponseStatusSuccess = "Success"
)

var (
	ErrNotFound            = utils.NewAPIError(http.StatusNotFound, "Resource not found", nil)
	ErrBadRequest          = utils.NewAPIError(http.StatusBadRequest, "Invalid request data", nil)
	ErrInternalServerError = utils.NewAPIError(http.StatusInternalServerError, "Internal Server Error", nil)
	ErrUnauthorized        = utils.NewAPIError(http.StatusUnauthorized, "Unauthorized access", nil)
	ErrConflict            = utils.NewAPIError(http.StatusConflict, "Resource already exists", nil)
	ErrForbidden           = utils.NewAPIError(http.StatusForbidden, "Forbidden access", nil)
)

type UsersStatusEnum string
type BorrowStatusEnum string
type DurationType string
type UserRoleEnum string
type PaymentStatusEnum string
type BookStatusEnum string

const (
	UserStatusActive   UsersStatusEnum = "active"
	UserStatusInactive UsersStatusEnum = "inactive"
	UserStatusBanned   UsersStatusEnum = "banned"

	BorrowStatusReturned BorrowStatusEnum = "returned"
	BorrowStatusSuccess  BorrowStatusEnum = "success"
	BorrowStatusPending  BorrowStatusEnum = "pending"
	BorrowStatusCancel   BorrowStatusEnum = "cancel"

	DurationTypeDaily   DurationType = "daily"
	DurationTypeWeekly  DurationType = "weekly"
	DurationTypeMonthly DurationType = "monthly"

	UserRoleReader UserRoleEnum = "reader"
	UserRoleAdmin  UserRoleEnum = "admin"

	PaymentStatusPending   PaymentStatusEnum = "pending"
	PaymentStatusConfirmed PaymentStatusEnum = "confirmed"
	PaymentStatusRefunded  PaymentStatusEnum = "refunded"

	BookStatusAvailable    BookStatusEnum = "available"
	BookStatusNotAvailable BookStatusEnum = "not available"
)
