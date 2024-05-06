package response

import "net/http"

type OkResponse struct {
	Status  int `json:"status"`
	Content any `json:"content,omitempty"`
}

type ErrorResponse struct {
	Status  int    `json:"status"`
	Error   string `json:"error,omitempty"`
	Message string `json:"message,omitempty"`
}

func Ok(content ...any) OkResponse {
	return OkResponse{
		Status:  http.StatusOK,
		Content: content,
	}
}

func BadRequest(msg string) ErrorResponse {
	return ErrorResponse{
		Status:  http.StatusBadRequest,
		Error:   "Incorrect request",
		Message: msg,
	}
}

func Unauthorized() ErrorResponse {
	return ErrorResponse{
		Status: http.StatusUnauthorized,
		Error:  "User not authorized",
	}
}

func Forbidden() ErrorResponse {
	return ErrorResponse{
		Status: http.StatusForbidden,
		Error:  "User forbidden",
	}
}

func NotFound() ErrorResponse {
	return ErrorResponse{
		Status: http.StatusNotFound,
		Error:  "Banner for user not found",
	}
}

func InternalServerError(msg string) ErrorResponse {
	return ErrorResponse{
		Status:  http.StatusInternalServerError,
		Error:   "Internal server error",
		Message: msg,
	}
}
