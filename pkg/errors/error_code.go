package errors

type ErrorCode int

const (
	Success ErrorCode = 0

	ErrCodeInternalServer ErrorCode = 500

	ErrCodeClientMemberNotFound    ErrorCode = 100
	ErrCodeClientPasswordMismatch  ErrorCode = 101
	ErrCodeClientMemberNotVerified ErrorCode = 102

	ErrCodeClientTokenInvalid ErrorCode = 200
	ErrCodeClientUnauthorized ErrorCode = 201

	ErrCodeBadParameter ErrorCode = 400
)
