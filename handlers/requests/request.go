package requests

type CreatePostRequest struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}
