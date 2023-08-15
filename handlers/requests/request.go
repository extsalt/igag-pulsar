package requests

// A CreatePostRequest represent payload for creating a post
type CreatePostRequest struct {
	Title    string `json:"title"`
	Body     string `json:"body"`
	ImageUrl string `json:"imageUrl"`
}

// A RegisterUserRequest represent payload for register user request
type RegisterUserRequest struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
