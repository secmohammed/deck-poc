package container

import (
	"github.com/secmohammed/deck-poc/config"
	"github.com/secmohammed/deck-poc/pkg/database"
	"sync"
)

var (
	instantiateAppOnce sync.Once
	appInstance        *container
)

type container struct {
	c  config.Repository
	db database.Repository
}

type Container interface {
	Config() config.Repository
	Database() database.Repository
}

func NewApplication(c config.Repository) Container {
	instantiateAppOnce.Do(func() {
		appInstance = &container{
			c:  c,
			db: database.NewDatabaseConnection(c),
		}
	})
	return appInstance
}

func (c *container) Get() *container {
	return c
}

func (c *container) Database() database.Repository {
	return c.db
}
func (c *container) Config() config.Repository {
	return c.c
}
