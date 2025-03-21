package handler

import (
	"github.com/gofiber/fiber/v2"
)

type TransactionHandler interface {
	ListTransaction(c *fiber.Ctx) error
	CreateTransaction(c *fiber.Ctx) error
	GetTransaction(c *fiber.Ctx) error
	UpdateTransaction(c *fiber.Ctx) error
	DeleteTransaction(c *fiber.Ctx) error
}
