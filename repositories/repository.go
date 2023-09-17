package repositories

import (
	"gorm.io/gorm"
	"pulsar/config"
)

const ResourceID = "resource_id"
const UserID = "user_id"

// Repo returns database connection
func Repo() *gorm.DB {
	return config.PulsarConfig.DB
}
