package users

import (
	"github.com/HaitaoYue/myweb/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"net/http"
)

var db = make(map[string]string)

func UserRetrieve(c *gin.Context) {
	user := c.Params.ByName("name")
	value, ok := db[user]
	glog.Infof("ok is %s %t %s", utils.Green, ok, utils.Reset)
	if ok {
		c.JSON(http.StatusOK, gin.H{"user": user, "value": value})
	} else {
		c.JSON(http.StatusOK, gin.H{"user": user, "status": "no value"})
	}
}
