package database

import (
	"fmt"
	"github.com/secmohammed/deck-poc/config"
	log "github.com/siruspen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"net/url"
)

type Repository interface {
	Get() *gorm.DB
}
type databaseConnection struct {
	DB *gorm.DB
}

func NewDatabaseConnection(config config.Repository) Repository {
	user := config.GetString("app.db.username")
	password := config.GetString("app.db.password")
	database := config.GetString("app.db.database")
	host := config.GetString("app.db.host")
	port := config.GetInt("app.db.port")
	var enableLogging logger.Interface
	enableDBLogging := config.GetBool("app.db.log")
	if enableDBLogging {
		enableLogging = logger.Default
	}
	dsn := url.URL{
		User:     url.UserPassword(user, password),
		Scheme:   "postgres",
		Host:     fmt.Sprintf("%s:%d", host, port),
		RawQuery: (&url.Values{"sslmode": []string{"disable"}}).Encode(),
	}
	fmt.Println(dsn.String())
	db, err := gorm.Open(postgres.Open(dsn.String()), &gorm.Config{
		Logger: enableLogging,
	})

	if err != nil {
		log.Fatalf("database connection failed: %s", err)
	}

	// check if db exists
	stmt := fmt.Sprintf("SELECT * FROM pg_database WHERE datname = '%s';", database)
	result := db.Raw(stmt)
	if result.Error != nil {
		log.Fatalf("database connection failed: %s", result.Error)

	}
	var rec = make(map[string]interface{})
	if result.Find(rec); len(rec) == 0 {
		stmt := fmt.Sprintf("CREATE DATABASE %s;", database)
		if rs := db.Exec(stmt); rs.Error != nil {
			log.Fatalf("database connection failed: %s", result.Error)
		}
	}
	dsn.Path = database
	fmt.Println(dsn.String())
	db, err = gorm.Open(postgres.Open(dsn.String()), &gorm.Config{
		Logger: enableLogging,
	})
	if err != nil {
		log.Fatalf("database connection failed: %s", err)

	}
	sync := config.GetBool("app.db.sync")
	if err != nil {
		log.Fatalf("database connection failed: %s", err)

	}
	if sync {
		if err := synchronize(db); err != nil {
			log.Fatalf("database connection failed: %s", err)

		}
	}
	return &databaseConnection{DB: db}
}
func (d *databaseConnection) Get() *gorm.DB {
	return d.DB
}
