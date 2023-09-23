package repositories

import (
	"pulsar/models"
)

// FindCommentById returns comment for given commentID
func FindCommentById(commentID uint64) (*models.Comment, error) {
	var comment models.Comment
	err := Repo().First(&comment, commentID).Error
	return &comment, err
}
