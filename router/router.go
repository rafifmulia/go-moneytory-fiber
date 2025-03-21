package router

import (
	"restfulapi/conf"
	"restfulapi/handler"
	"restfulapi/helper"
	mw "restfulapi/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/pprof"
)

const (
	basePath string = "/v1"
)

var (
	httpPprof bool
	app       *fiber.App
	trx       handler.TransactionHandler
)

func setFlags() {
	httpPprof = conf.GetHttpPprof()
}

func initResources() {
	setFlags()
	app = conf.InitFiber()
	mw.InitResources()
	initHandler()
}

func initHandler() {
	trx = handler.NewTransactionHandler()
}

func panicRoute(c *fiber.Ctx) error {
	panic("panic route")
}

func notFoundRoute(c *fiber.Ctx) error {
	return helper.RespNotFound(c, "Route not found")
}

func InitRouter() *fiber.App {
	initResources()
	app.Use(mw.RootMiddleware)
	app.Get("/panic", panicRoute)
	if httpPprof {
		app.Use(pprof.New())
	}
	au := app.Group(basePath, mw.AuthMiddleware)
	au.Get("/transaction", trx.ListTransaction)
	au.Post("/transaction", trx.CreateTransaction)
	au.Get("/transaction/:trxId", trx.GetTransaction)
	au.Put("/transaction/:trxId", trx.UpdateTransaction)
	au.Delete("/transaction/:trxId", trx.DeleteTransaction)
	app.Use(notFoundRoute)
	return app
}
