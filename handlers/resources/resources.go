package resources

import (
	"pulsar/models"
	"strconv"
)

type PostResource struct {
	ID            string `json:"id"`
	Title         string `json:"title"`
	Slug          string `json:"slug"`
	Body          string `json:"body"`
	OriginalImage string `json:"originalImage"`
	SmImage       string `json:"smImage"`
	MdImage       string `json:"mdImage"`
	LgImage       string `json:"lgImage"`
	CreatedAt     string `json:"createdAt"`
	User          struct {
		Username string `json:"username"`
		Avatar   string `json:"avatar"`
	} `json:"user"`
}

func PostsJsonResource(posts *[]models.Post) *[]PostResource {
	if len(*posts) == 0 {
		return &([]PostResource{})
	}
	var resource []PostResource
	for _, post := range *posts {
		resource = append(resource, PostResource{
			ID:            strconv.FormatUint(post.ID, 10),
			Title:         post.Title,
			Slug:          post.Slug,
			Body:          post.Body,
			OriginalImage: post.OriginalImage,
			CreatedAt:     post.CreatedAt.String(),
		})
	}
	return &resource
}
