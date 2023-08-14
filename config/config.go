package config

import "gorm.io/gorm"

// A PulsarConfig is globally available config
var PulsarConfig *Config

type Config struct {
	DB *gorm.DB
}
