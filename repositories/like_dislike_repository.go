package repositories

import (
	"pulsar/config"
	"pulsar/models"
)

// FindLikeForResourceForUser returns returns like of user for a post
func FindLikeForResourceForUser(resourceID uint64, resourceType string, userID uint64) (*models.LikeDislike, error) {
	var likeDislike models.LikeDislike
	err := config.PulsarConfig.DB.
		Where(map[string]interface{}{"action": 1, "resource_id": resourceID, "resource_type": resourceType, "user_id": userID}).
		Limit(1).
		First(&likeDislike).
		Error
	if err != nil {
		return nil, err
	}
	return &likeDislike, err
}
