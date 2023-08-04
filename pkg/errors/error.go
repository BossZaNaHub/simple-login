package errors

type Error interface {
	Code() ErrorCode
	SetMessage(msg string)
	Message() string
	Error() string
}
