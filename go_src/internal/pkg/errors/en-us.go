package errors

var enUSText = map[int]string{
	SUCCESS:            "OK",
	FAILURE:            "FAIL",
	NotFound:           "resources not found",
	ServerError:        "Internal server error",
	TooManyRequests:    "Too many requests",
	InvalidParameter:   "Parameter error",
	UserDoesNotExist:   "user does not exist",
	AuthorizationError: "Authorization error",
	RBACError:          "No access",
	UserISExist:        "User is Exits",
	BLOGIDNOTFOUND:     "Blog id not found",
	NOTOKEN:            "No token. you don't have permission!",
	TOKENERROR:         "Token error",
}
