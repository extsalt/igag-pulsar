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

// A CreateCommentRequest represent payload for creating a comment
type CreateCommentRequest struct {
	Body     string `json:"body"`
	ImageUrl string `json:"imageUrl"`
}

// A CreateReplyRequest represent payload for creating a reply
type CreateReplyRequest struct {
	Body     string `json:"body"`
	ImageUrl string `json:"imageUrl"`
}
