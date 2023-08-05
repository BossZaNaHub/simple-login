package errors

import "github.com/gofiber/fiber/v2"

type Error interface {
	Code() ErrorCode
	SetMessage(msg string)
	Message() string
	Error() string
	FiberMap() fiber.Map
}
