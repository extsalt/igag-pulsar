package resources

import (
	"pulsar/models"
	"strconv"
)

type UserResource struct {
	Avatar   string `json:"avatar"`
	Username string `json:"username"`
}
type PostResource struct {
	ID        string       `json:"id"`
	Title     string       `json:"title"`
	Slug      string       `json:"slug"`
	Body      string       `json:"body"`
	Image     string       `json:"image"`
	SmImage   string       `json:"smImage"`
	MdImage   string       `json:"mdImage"`
	LgImage   string       `json:"lgImage"`
	CreatedAt string       `json:"createdAt"`
	User      UserResource `json:"user"`
}

func PostsJsonResource(posts *[]models.Post) *[]PostResource {
	if len(*posts) == 0 {
		return &([]PostResource{})
	}
	var resource []PostResource
	for _, post := range *posts {
		resource = append(resource, PostResource{
			ID:        strconv.FormatUint(post.ID, 10),
			Title:     post.Title,
			Slug:      post.Slug,
			Body:      post.Body,
			Image:     post.OriginalImage,
			CreatedAt: post.CreatedAt.String(),
			User:      UserResource{},
		})
	}
	return &resource
}
