package config

import (
	log "github.com/siruspen/logrus"
)

type Type string

const (
	// Local local config
	Local Type = "local"
)

// Factory returns the requested config repo
func Factory(c Type) Repository {
	switch Type(c) {
	case Local:
		return NewLocalRepository()

	default:
		log.Fatalf("Unknown config repository: %s", c)
	}

	return nil
}

// Repository port
type Repository interface {
	Get(key string) (interface{}, error)
	GetInt(key string) int64
	GetFloat(key string) float64
	GetString(key string) string
	GetBool(key string) bool
	GetStringSlice(key string) []string
}
