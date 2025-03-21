package handler

import (
	"context"
	"net/http"
	"restfulapi/api"
	"restfulapi/exception"
	"restfulapi/helper"
	"restfulapi/libs"
	"restfulapi/service"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type TransactionHandlerImpl struct {
	svc      service.TransactionService
	validate *validator.Validate
}

func NewTransactionHandler() TransactionHandler {
	var (
		validate *validator.Validate        = libs.ExportValidator()
		svc      service.TransactionService = service.NewTransactionServiceImpl()
	)
	return &TransactionHandlerImpl{svc: svc, validate: validate}
}

func (impl *TransactionHandlerImpl) ListTransaction(c *fiber.Ctx) error {
	var (
		now       time.Time       = time.Now()
		rctx      context.Context = c.Context()
		err       error
		trxParams *api.GetTransactionParams = &api.GetTransactionParams{}
		trxResp   []*api.Transaction
	)
	if err := c.QueryParser(trxParams); err != nil {
		panic(exception.NewBadRequestException(err.Error()))
	}
	switch trxParams.Filter {
	case "today":
		trxParams.RangeStart = helper.StartOfDay(&now).Format(time.RFC3339)
		trxParams.RangeEnd = helper.EndOfDay(&now).Format(time.RFC3339)
	case "week":
		trxParams.RangeStart = helper.StartOfWeek(&now).Format(time.RFC3339)
		trxParams.RangeEnd = helper.EndOfWeek(&now).Format(time.RFC3339)
	case "year":
		trxParams.RangeStart = helper.StartOfYear(&now).Format(time.RFC3339)
		trxParams.RangeEnd = helper.EndOfYear(&now).Format(time.RFC3339)
	case "custom":
		trxParams.RangeStart = helper.StartOfDay(helper.StrDateToTime(trxParams.RangeStart)).Format(time.RFC3339)
		trxParams.RangeEnd = helper.EndOfDay(helper.StrDateToTime(trxParams.RangeEnd)).Format(time.RFC3339)
	default: // case "month":
		trxParams.RangeStart = helper.StartOfMonth(&now).Format(time.RFC3339)
		trxParams.RangeEnd = helper.EndOfMonth(&now).Format(time.RFC3339)
	}
	ctx, cancel := context.WithCancel(rctx)
	defer cancel()
	trxResp, err = impl.svc.FindAll(ctx, trxParams)
	if err != nil {
		panic(err)
	}
	err = c.Status(http.StatusOK).JSON(api.RespListTransactions{
		Meta: &api.Meta{
			Code:    200,
			Message: "Success list transactions",
		},
		Data: trxResp,
	})
	if err != nil {
		panic(err)
	}
	return nil
}

func (impl *TransactionHandlerImpl) GetTransaction(c *fiber.Ctx) error {
	var (
		err     error
		rctx    context.Context = c.Context()
		trxId   string          = c.Params("trxId")
		trxResp *api.Transaction
	)
	if len(trxId) != 36 {
		panic(exception.NewUnprocessableEntityException("Transaction id is not valid"))
	}
	ctx, cancel := context.WithCancel(rctx)
	defer cancel()
	trxResp, err = impl.svc.FindById(ctx, &trxId)
	if err != nil {
		panic(err)
	}
	err = c.Status(http.StatusOK).JSON(api.RespDetailTransaction{
		Meta: &api.Meta{
			Code:    200,
			Message: "Success get detail transaction",
		},
		Data: trxResp,
	})
	if err != nil {
		panic(err)
	}
	return nil
}

func (impl *TransactionHandlerImpl) CreateTransaction(c *fiber.Ctx) error {
	var (
		err     error
		rctx    context.Context           = c.Context()
		trxReq  *api.ReqCreateTransaction = &api.ReqCreateTransaction{}
		trxResp *api.Transaction
	)
	err = c.BodyParser(trxReq)
	if err != nil {
		panic(exception.NewBadRequestException(err.Error()))
	}
	err = impl.validate.Struct(trxReq)
	if err != nil {
		panic(err)
	}
	ctx, cancel := context.WithCancel(rctx)
	defer cancel()
	trxResp = impl.svc.Create(ctx, trxReq)
	err = c.Status(http.StatusCreated).JSON(api.RespDetailTransaction{
		Meta: &api.Meta{
			Code:    201,
			Message: "Success create transaction",
		},
		Data: trxResp,
	})
	if err != nil {
		panic(err)
	}
	return nil
}

func (impl *TransactionHandlerImpl) UpdateTransaction(c *fiber.Ctx) error {
	var (
		err     error
		rctx    context.Context           = c.Context()
		trxId   string                    = c.Params("trxId")
		trxReq  *api.ReqUpdateTransaction = &api.ReqUpdateTransaction{}
		trxResp *api.Transaction
	)
	if len(trxId) != 36 {
		panic(exception.NewUnprocessableEntityException("Transaction id is not valid"))
	}
	err = c.BodyParser(trxReq)
	if err != nil {
		panic(exception.NewBadRequestException(err.Error()))
	}
	err = impl.validate.Struct(trxReq)
	if err != nil {
		panic(err)
	}
	ctx, cancel := context.WithCancel(rctx)
	defer cancel()
	trxResp, err = impl.svc.Update(ctx, trxReq, &trxId)
	if err != nil {
		panic(err)
	}
	err = c.Status(http.StatusOK).JSON(api.RespDetailTransaction{
		Meta: &api.Meta{
			Code:    200,
			Message: "Success update transaction",
		},
		Data: trxResp,
	})
	if err != nil {
		panic(err)
	}
	return nil
}

func (impl *TransactionHandlerImpl) DeleteTransaction(c *fiber.Ctx) error {
	var (
		err   error
		rctx  context.Context = c.Context()
		trxId string          = c.Params("trxId")
	)
	if len(trxId) != 36 {
		panic(exception.NewUnprocessableEntityException("Transaction id is not valid"))
	}
	ctx, cancel := context.WithCancel(rctx)
	defer cancel()
	err = impl.svc.Delete(ctx, &trxId)
	if err != nil {
		panic(err)
	}
	err = c.Status(http.StatusOK).JSON(api.RespDetailTransaction{
		Meta: &api.Meta{
			Code:    200,
			Message: "Success delete transaction",
		},
	})
	if err != nil {
		panic(err)
	}
	return nil
}
