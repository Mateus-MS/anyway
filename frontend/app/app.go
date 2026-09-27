package app

import (
	"os"
	"sync"

	"github.com/gin-gonic/gin"
)

var (
	instance *application
	once     sync.Once
)

type application struct {
	Router *gin.Engine
}

func new() *application {
	if os.Getenv("APP_ENV") == "prod" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	return &application{
		Router: gin.Default(),
	}
}

func GetInstance() *application {
	once.Do(func() {
		instance = new()
	})

	return instance
}
