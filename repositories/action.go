package repositories

import (
	"pulsar/models"
)

const ResourceType = "resource_type"

func FindActionOnResource(resourceID uint64, resourceType models.Resource, userID uint64) (*models.LikeDislike, error) {
	var likeDislike models.LikeDislike
	err := Repo().
		Where(map[string]interface{}{ResourceID: resourceID, ResourceType: resourceType, UserID: userID}).
		Limit(1).
		First(&likeDislike).
		Error
	return &likeDislike, err
}

// AddActionToResource Add like to resource
func AddActionToResource(resourceID uint64, resource models.Resource, userID uint64, action models.Action) (*models.LikeDislike, error) {
	actionResource := models.LikeDislike{
		UserID:       userID,
		ResourceID:   resourceID,
		ResourceType: resource,
		Action:       action,
	}
	err := Repo().Create(&actionResource).Error
	return &actionResource, err
}

// RemoveActionFromResource Remove action from resource
func RemoveActionFromResource(actionId uint64, resourceID uint64, resource models.Resource, userID uint64, action models.Action) (*models.LikeDislike, error) {
	actionResource := models.LikeDislike{
		ID:           actionId,
		UserID:       userID,
		ResourceID:   resourceID,
		ResourceType: resource,
		Action:       action,
	}
	err := Repo().Delete(&actionResource).Error
	return &actionResource, err
}
