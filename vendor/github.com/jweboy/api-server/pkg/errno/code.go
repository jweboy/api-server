package errno

var (
	// Common errors
	OK                  = &Errno{Code: 0, Message: "ok"}
	InternalServerError = &Errno{Code: 10001, Message: "Internal server error."}
	ErrBind             = &Errno{Code: 10002, Message: "Error occurred while binding the request body to the struct."}

	ErrValidation = &Errno{Code: 20001, Message: "Validation failed."}
	ErrDatabase   = &Errno{Code: 20002, Message: "Database error."}
	ErrToken      = &Errno{Code: 20003, Message: "Error occurred while signing the JSON web token."}

	// User Errors
	ErrEncrypt           = &Errno{Code: 20101, Message: "Error occurred while encrypting the user password."}
	ErrUserNotFound      = &Errno{Code: 20102, Message: "The user was not found."}
	ErrTokenInvalid      = &Errno{Code: 20103, Message: "The token was invalid."}
	ErrPasswordIncorrect = &Errno{Code: 20104, Message: "The password was incorrect."}

	// Bucket errors
	ErrListBucketError = &Errno{Code: 20130, Message: "List Bucket error."}
	// File errors
	ErrFileUpload = &Errno{Code: 20150, Message: "File upload error"}
	ErrFileDelete = &Errno{Code: 20151, Message: "File delete error"}
)
