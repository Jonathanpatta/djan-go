package djan_go

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/qor/roles"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Config struct {
	GormDb *gorm.DB
	DbPath string
	Debug  bool
	Router *mux.Router
}

type DataModelConfig struct {
	GlobalConfig *Config
	Permission   *roles.Permission
}

func NewDefaultConfig() (*Config, error) {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	router := mux.NewRouter()
	router.Use(CorsMiddleware)
	router.StrictSlash(false)
	return &Config{
		Debug:  true,
		GormDb: db,
		Router: router,
	}, nil
}

func NewGormConfig(db *gorm.DB) (*Config, error) {
	router := mux.NewRouter()
	router.Use(CorsMiddleware)
	router.StrictSlash(false)
	return &Config{
		Debug:  true,
		GormDb: db,
		Router: router,
	}, nil
}

func NewPostgresConfig(pgurl string) (*Config, error) {
	db, err := gorm.Open(postgres.Open(pgurl), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
	router := mux.NewRouter()
	router.Use(CorsMiddleware)
	router.StrictSlash(false)
	return &Config{
		Debug:  true,
		GormDb: db,
		Router: router,
	}, nil
}
