package repositories

import (
	"gorm.io/gorm"
	"pulsar/pkg/oauth2identity"
)

func FindUserByOauth2Identify(identity *oauth2identity.Oauth2Identity) (*gorm.Model, error) {
	return nil, nil
}
