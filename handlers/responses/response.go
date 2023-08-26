package responses

type NotFoundResponse map[string]any

func GetNotFoundResponse() *NotFoundResponse {
	var response NotFoundResponse
	response["message"] = "Resource does not exists"
	return &response
}
