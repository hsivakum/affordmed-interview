package init

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Init(router *gin.Engine) {
	db := Db()
	if err := db.Ping(); err != nil {
		logrus.Errorf("unable to ping db %v", err)
	}

	initObjects(db)
	initRoutes(router)
}
