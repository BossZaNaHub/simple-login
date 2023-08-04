package errors

var (
	ErrClientMemberNotFound    = NewError(ErrCodeClientMemberNotFound, "member not found.")
	ErrClientPasswordMismatch  = NewError(ErrCodeClientPasswordMismatch, "password incorrect.")
	ErrClientMemberNotVerified = NewError(ErrCodeClientMemberNotVerified, "password incorrect.")
	ErrClientTokenInvalid      = NewError(ErrCodeClientTokenInvalid, "token invalid")
)
