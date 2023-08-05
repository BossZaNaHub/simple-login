package errors

import "github.com/gofiber/fiber/v2"

type defaultError struct {
	code    ErrorCode
	message string
}

func (d *defaultError) SetMessage(msg string) {
	d.message = msg
}

func (d *defaultError) Message() string {
	return d.message
}

func (d *defaultError) Error() string {
	return d.message
}

func (d *defaultError) Code() ErrorCode {
	return d.code
}

func (d *defaultError) FiberMap() fiber.Map {
	return fiber.Map{"code": d.code, "message": d.message}
}

func NewError(code ErrorCode, message string) Error {
	return &defaultError{
		code:    code,
		message: message,
	}
}

func NewDefaultError(err error) Error {
	return &defaultError{
		code:    ErrCodeInternalServer,
		message: err.Error(),
	}
}
