package server

import (
	"aws-test/settings"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
)


func Init() {
	log := logrus.New()
	environments := settings.GetEnv()
	engine := gin.New()
	log.SetFormatter(&logrus.JSONFormatter{})

	engine.Use(cors.Default(),
		settings.LogMiddleware(log),
		gin.Recovery(),
		gzip.Gzip(gzip.DefaultCompression),
	)

	RouterInit(engine)
	engine.Run(fmt.Sprintf(":%v", environments.Port))
}
