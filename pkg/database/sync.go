package database

import (
	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/secmohammed/deck-poc/internal/entity"
	"gorm.io/gorm"
)

func synchronize(db *gorm.DB) error {
	m := gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		{
			ID: "initial",
			Migrate: func(tx *gorm.DB) error {
				if err := tx.AutoMigrate(
					&entity.CardTemplate{},
					&entity.Deck{},
					&entity.CardDeck{},
				); err != nil {
					return err
				}
				return SeedTemplates(tx)
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Migrator().DropTable(
					&entity.CardDeck{},
					&entity.CardTemplate{},
					&entity.Deck{},
				)
			},
		},
	})
	return m.Migrate()
}
