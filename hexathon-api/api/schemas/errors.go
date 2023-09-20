package schemas

// Error is the schema for error response

var InvalidBody = map[string]interface{}{
	"detail": "Invalid request body",
}

var InvalidParams = map[string]interface{}{
	"detail": "Invalid request params",
}

var InternalServerError = map[string]interface{}{
	"detail": "Internal Server Error",
}

var NotFound = map[string]interface{}{
	"detail": "Not Found",
}
