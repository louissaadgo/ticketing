package models

type Error struct {
	Message string `json:"message"`
	Field   string `json:"field"`
}

type ErrorResponse struct {
	Errors []Error `json:"errors"`
}
