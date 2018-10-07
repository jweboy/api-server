package errno

var (
	// Common errors
	OK                  = &Errno{Code: 0, Message: "ok"}
	InternalServerError = &Errno{Code: 10001, Message: "Internal server error."}
	ErrBind             = &Errno{Code: 10002, Message: "Error occurred while binding the request body to the struct."}

	// User Errors
	ErrUserNotFound = &Errno{Code: 20102, Message: "The user was not found."}
)
