package repositories

import (
	"fmt"
	"pulsar/config"
	"pulsar/models"
	"pulsar/pkg/oauth2identity"
)

func FindUserByOauth2Identify(identity oauth2identity.Oauth2Identity) (*models.User, error) {
	email, err := identity.GetEmail()
	if err != nil {
		return nil, err
	}
	user, err := FindUserByEmail(email)
	if err != nil {
		err = CreateUserByEmail(email)
		if err != nil {
			return nil, err
		}
		user, err = FindUserByEmail(email)
	}
	return user, err
}

func CreateUserByEmail(email string) error {
	err := config.PulsarConfig.DB.Create(&models.User{
		Email: email,
	}).Error
	if err != nil {
		return err
	}
	var lastInsertId struct {
		ID int
	}
	err = config.PulsarConfig.DB.Raw("select last_insert_id() as id limit 1").Scan(&lastInsertId).Error
	if err != nil {
		return err
	}
	err = config.PulsarConfig.DB.Exec("UPDATE users SET username = ? WHERE id = ?", fmt.Sprintf("user%d", lastInsertId.ID), lastInsertId.ID).Error
	return err
}

func FindUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := config.PulsarConfig.DB.Where(&models.User{Email: email}).First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}
