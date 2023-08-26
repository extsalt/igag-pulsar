package repositories

import (
	"pulsar/config"
	"pulsar/models"
)

// FindPostById returns post for given postID
func FindPostById(postID string) (*models.Post, error) {
	var post models.Post
	err := config.PulsarConfig.DB.First(&post, postID).Error
	if err != nil {
		return nil, err
	}
	return &post, err
}
