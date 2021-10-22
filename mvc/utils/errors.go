package utils

type AppplicationError struct {
	Message    error  `json:"message"`
	StatusCode int    `json:"status"`
	Code       string `json:"code"`
}
