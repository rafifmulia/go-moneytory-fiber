package helper

import (
	"net/http"
	"restfulapi/api"

	"github.com/gofiber/fiber/v2"
)

func RespBadRequest(c *fiber.Ctx, msg string) error {
	var err error
	if msg == "" {
		msg = "Problem parsing request data."
	}
	err = c.Status(http.StatusBadRequest).JSON(api.RespBadRequest{
		Meta: &api.Meta{
			Code:    400,
			Message: msg,
		},
	})
	if err != nil {
		return err
	}
	return nil
}

func RespUnauthorized(c *fiber.Ctx, msg string) error {
	var err error
	if msg == "" {
		msg = "Unauthorized."
	}
	err = c.Status(http.StatusUnauthorized).JSON(api.RespUnauthorized{
		Meta: &api.Meta{
			Code:    401,
			Message: msg,
		},
	})
	if err != nil {
		return err
	}
	return nil
}

func RespNotFound(c *fiber.Ctx, msg string) error {
	var err error
	if msg == "" {
		msg = "Data is empty or not found."
	}
	err = c.Status(http.StatusNotFound).JSON(api.RespNotFound{
		Meta: &api.Meta{
			Code:    404,
			Message: msg,
		},
	})
	if err != nil {
		return err
	}
	return nil
}

func RespUnprocessableEntity(c *fiber.Ctx, msg string) error {
	var err error
	if msg == "" {
		msg = "Fields is unprocessable."
	}
	err = c.Status(http.StatusUnprocessableEntity).JSON(api.RespUnprocessableEntity{
		Meta: &api.Meta{
			Code:    422,
			Message: msg,
		},
	})
	if err != nil {
		return err
	}
	return nil
}

func RespInternalServerError(c *fiber.Ctx, msg string) error {
	var err error
	if msg == "" {
		msg = "Internal server error."
	}
	err = c.Status(http.StatusInternalServerError).JSON(api.RespInternalServerError{
		Meta: &api.Meta{
			Code:    500,
			Message: msg,
		},
	})
	if err != nil {
		return err
	}
	return nil
}
