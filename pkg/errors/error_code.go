package errors

type ErrorCode int

const (
	Success ErrorCode = 0

	ErrCodeInternalServer ErrorCode = 500

	ErrCodeClientMemberNotFound    = 100
	ErrCodeClientPasswordMismatch  = 101
	ErrCodeClientMemberNotVerified = 102

	ErrCodeClientTokenInvalid = 200

	ErrCodeBadParameter ErrorCode = 400
)
