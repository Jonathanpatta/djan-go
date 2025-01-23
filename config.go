package djan_go

import (
	"github.com/gorilla/mux"
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
	Auth         bool
	EndPointName string
	GlobalConfig *Config
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
