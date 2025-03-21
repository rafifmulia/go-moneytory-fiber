// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.4.1 DO NOT EDIT.
package api

import (
	"time"

	uuid "github.com/google/uuid"
)

const (
	UserAuthScopes = "UserAuth.Scopes"
)

// Defines values for EnumCalendarFilter.
const (
	Custom EnumCalendarFilter = "custom"
	Month  EnumCalendarFilter = "month"
	Today  EnumCalendarFilter = "today"
	Week   EnumCalendarFilter = "week"
	Year   EnumCalendarFilter = "year"
)

// EnumCalendarFilter defines model for EnumCalendarFilter.
type EnumCalendarFilter string

// Meta defines model for Meta.
type Meta struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

// ReqCreateTransaction defines model for ReqCreateTransaction.
type ReqCreateTransaction struct {
	Amount      *float64   `form:"amount,omitempty" json:"amount" validate:"required,number,gt=0"`
	Date        *time.Time `form:"date,omitempty" json:"date" validate:"required"`
	Description *string    `form:"description,omitempty" json:"description" validate:"required,printascii,lte=125"`
}

// ReqUpdateTransaction defines model for ReqUpdateTransaction.
type ReqUpdateTransaction struct {
	Amount      *float64   `form:"amount,omitempty" json:"amount,omitempty" validate:"omitempty,number,gt=0"`
	Date        *time.Time `form:"date,omitempty" json:"date,omitempty"`
	Description *string    `form:"description,omitempty" json:"description,omitempty" validate:"omitempty,printascii,lte=125"`
}

// RespDetailTransaction defines model for RespDetailTransaction.
type RespDetailTransaction struct {
	Data *Transaction `json:"data,omitempty"`
	Meta *Meta        `json:"meta,omitempty"`
}

// RespListTransactions defines model for RespListTransactions.
type RespListTransactions struct {
	Data Transactions `json:"data,omitempty"`
	Meta *Meta        `json:"meta,omitempty"`
}

// Transaction defines model for Transaction.
type Transaction struct {
	Amount      *float64   `json:"amount,omitempty"`
	Date        *time.Time `json:"date,omitempty"`
	Description *string    `json:"description,omitempty"`
	Uuid        *string    `json:"uuid,omitempty"`
}

// Transactions defines model for Transactions.
type Transactions = []*Transaction

// TypeDate defines model for TypeDate.
type TypeDate = string

// ParamTrxFilter defines model for paramTrxFilter.
type ParamTrxFilter = EnumCalendarFilter

// ParamTrxId defines model for paramTrxId.
type ParamTrxId = uuid.UUID

// ParamTrxRangeEnd defines model for paramTrxRangeEnd.
type ParamTrxRangeEnd = TypeDate

// ParamTrxRangeStart defines model for paramTrxRangeStart.
type ParamTrxRangeStart = TypeDate

// RespBadRequest defines model for RespBadRequest.
type RespBadRequest struct {
	Meta *Meta `json:"meta,omitempty"`
}

// RespInternalServerError defines model for RespInternalServerError.
type RespInternalServerError struct {
	Meta *Meta `json:"meta,omitempty"`
}

// RespNotFound defines model for RespNotFound.
type RespNotFound struct {
	Meta *Meta `json:"meta,omitempty"`
}

// RespUnauthorized defines model for RespUnauthorized.
type RespUnauthorized struct {
	Meta *Meta `json:"meta,omitempty"`
}

// RespUnprocessableEntity defines model for RespUnprocessableEntity.
type RespUnprocessableEntity struct {
	Meta *Meta `json:"meta,omitempty"`
}

// GetTransactionParams defines parameters for GetTransaction.
type GetTransactionParams struct {
	Filter     ParamTrxFilter     `form:"filter,omitempty" json:"filter,omitempty" query:"filter,omitempty"`
	RangeStart ParamTrxRangeStart `form:"range_start,omitempty" json:"range_start,omitempty" query:"range_start,omitempty"`
	RangeEnd   ParamTrxRangeEnd   `form:"range_end,omitempty" json:"range_end,omitempty" query:"range_end,omitempty"`
}

// PostTransactionFormdataRequestBody defines body for PostTransaction for application/x-www-form-urlencoded ContentType.
type PostTransactionFormdataRequestBody = ReqCreateTransaction

// PutTransactionParamTrxIdFormdataRequestBody defines body for PutTransactionParamTrxId for application/x-www-form-urlencoded ContentType.
type PutTransactionParamTrxIdFormdataRequestBody = ReqUpdateTransaction
