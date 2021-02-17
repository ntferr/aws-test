package settings

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"time"
)

func LogMiddleware(logger logrus.FieldLogger) gin.HandlerFunc {

	return func(c *gin.Context) {

		if c.Request.URL.Path != fmt.Sprintf("%s:%s/aws/healthcheck",
			GetEnv().Host, GetEnv().Port) {

			c.Next()
			statusCode := c.Writer.Status()
			entry := logger.WithFields(logrus.Fields{
				"log_version": 	"1.0",
				"request_id": 	"undefined",
				"date_time":	time.Now(),
			})

			if len(c.Errors) > 0 {
				entry.Error(c.Errors.ByType(gin.ErrorTypePrivate).String())
			} else {
				if statusCode > 499 {
					entry.Error()
				} else if statusCode > 399 {
					entry.Warn()
				} else {
					entry.Info()
				}
			}
		}
	}
}
