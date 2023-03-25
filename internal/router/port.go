package router

import (
	"github.com/secmohammed/deck-poc/container"
	log "github.com/siruspen/logrus"
)

type Type string

const (
	REST Type = "rest"
)

// Factory returns the requested config repo
func Factory(t Type, c container.Container) Repository {
	switch t {
	case REST:
		return NewRestRepository(c)
	default:
		log.Fatalf("Unknown config repository: %s", c)
	}

	return nil
}

// Repository port
type Repository interface {
	Expose() error
}
