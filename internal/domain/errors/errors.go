package errors

type ErrorCode string

const (
	MsgOk                  = "OK"
	MsgInternalServerError = "Internal Server Error"
	MsgBadRequest          = "Bad Request"
	MsgNotFound            = "Not Found"
	MsgCreated             = "Created"
	MsgUnauthorized        = "Incorrect email or password"
)

var (
	ErrNotFound     ErrorCode = "NOT_FOUND"
	ErrInvalidInput ErrorCode = "INVALID_INPUT"
	ErrInternal     ErrorCode = "INTERNAL_ERROR"
	ErrUnauthorized ErrorCode = "UNAUTHORIZED"
)

type DomainError struct {
	Code    ErrorCode
	Message string
	Err     error
}

func (e *DomainError) Error() string {
	return e.Message
}
