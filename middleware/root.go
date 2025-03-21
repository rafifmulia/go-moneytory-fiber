package middleware

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
	"restfulapi/conf"
	"restfulapi/exception"
	"restfulapi/helper"
	"runtime/debug"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var (
	debugMode bool
)

func setFlags() {
	debugMode = conf.GetDebugFlag()
}

func InitResources() {
	setFlags()
}

func catchPanic(c *fiber.Ctx, end func()) {
	msg := recover()
	if msg != nil {
		switch err := msg.(type) {
		case *exception.BadRequestException:
			helper.RespBadRequest(c, getErrMessage(err))
		case *exception.UnauthorizedException:
			helper.RespUnauthorized(c, getErrMessage(err))
		case *exception.NotFoundException:
			helper.RespNotFound(c, getErrMessage(err))
		case *exception.UnprocessableEntityException:
			helper.RespUnprocessableEntity(c, getErrMessage(err))
		case *validator.InvalidValidationError:
			helper.RespUnprocessableEntity(c, getErrMessage(err))
		case validator.ValidationErrors:
			helper.RespUnprocessableEntity(c, getErrMessage(err))
		case validator.FieldError:
			helper.RespUnprocessableEntity(c, getErrMessage(err))
		case error:
			helper.RespInternalServerError(c, getErrMessage(err))
		default:
			val := reflect.ValueOf(msg)
			tp := val.Type()
			log.Printf("panicHandler type:%s\npanicHandler val:%s\npanicHandler stack trace:%s\n", tp.Name(), msg, debug.Stack())
			if debugMode {
				helper.RespInternalServerError(c, fmt.Sprintf("%s", msg))
			} else {
				helper.RespInternalServerError(c, "")
			}
		}
	}
	end()
}

func getErrMessage(err error) string {
	if debugMode {
		return err.Error()
	}
	return ""
}

// Catch panic in handlers or middlewares and return http error response.
// Log incoming requests and its status code.
// CORS middleware.
func RootMiddleware(c *fiber.Ctx) error {
	log.Printf("Incoming request from %s %s\n", c.Method(), c.OriginalURL())
	defer catchPanic(c, func() {
		log.Printf("Request from %s %s has been responded %d\n", c.Method(), c.OriginalURL(), c.Response().StatusCode())
	})
	if c.Method() == "OPTIONS" {
		c.Set("Access-Control-Allow-Origin", "*")
		c.Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
		c.Set("Access-Control-Allow-Headers", "X-API-Key, Origin, Content-Type, Accept")
		c.Set("Access-Control-Max-Age", fmt.Sprintf("%d", 2*time.Hour))
		return c.SendStatus(http.StatusNoContent)
	}
	err := c.Next()
	if err != nil {
		panic(err)
	}
	return err
}
