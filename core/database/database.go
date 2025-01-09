package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type (
	Client struct {
		DB *gorm.DB
	}
	Config struct {
		Host     string
		Port     string
		User     string
		Password string
		DBName   string
		SSLMode  string
	}
)

type (
	Connection interface {
		CreateRecord(value interface{}) error
		FindRecord(model interface{}, id interface{}, preloadAssociations ...string) error
		UpdateRecord(model interface{}) error
		DeleteRecord(model interface{}) error
	}
)

var _ Connection = (*Client)(nil)

// Initialize creates a connection to the database and
// stores the reference to `DB` which can be used for further database operations
func Initialize(config *Config) (*Client, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.Host, config.Port, config.User, config.Password, config.DBName, config.SSLMode,
	)

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logger.Info, // Log level
			Colorful:      true,        // Enable color
		},
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})

	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	// Set connection pool settings
	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get database instance: %v", err)
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	return &Client{
		DB: db,
	}, nil
}

func (c Client) GetModel(name string) Client {
	return Client{DB: c.DB.Table(name)}
}

func (c Client) Migrate(models ...interface{}) error {
	return c.DB.AutoMigrate(models...)
}

func (c Client) CreateRecord(value interface{}) error {
	return c.DB.Create(value).Error
}

func (c Client) FindRecord(model interface{}, id interface{}, preloadAssociations ...string) error {
	query := c.DB
	for _, association := range preloadAssociations {
		query = query.Preload(association)
	}
	return query.First(model, id).Error
}

func (c Client) UpdateRecord(model interface{}) error {
	return c.DB.Save(model).Error
}

func (c Client) DeleteRecord(model interface{}) error {
	return c.DB.Delete(model).Error
}
