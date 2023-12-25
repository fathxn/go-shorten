package dto

const (
	MsgOk                  = "OK"
	MsgInternalServerError = "Internal Server Error"
	MsgBadRequest          = "Bad Request"
	MsgNotFound            = "Not Found"
	MsgCreated             = "Created"
	MsgUnauthorized        = "Incorrect email or password"
)

type APIResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
