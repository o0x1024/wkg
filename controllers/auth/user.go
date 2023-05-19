package auth

import (
	"net/http"
	"wkg/model"
	"wkg/model/request"
	"wkg/services/userService"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func SaveUser(c *gin.Context) {
	var err error
	var user = &model.Users{}
	if err = c.ShouldBindJSON(user); err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": err.Error()})
		return
	}

	err = userService.SaveUser(user)
	if err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "success"})
}

func DelUser(c *gin.Context) {
	id := c.Query("id")
	if id == "1" {
		zap.S().Errorf("can not delete this user")
		c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "can not delete this user"})
		return
	}
	if err := userService.DelUser(id); err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": err.Error()})

	}
	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "delete success."})
}

func ResetPwd(c *gin.Context) {
	id := c.Query("id")
	if id == "1" {
		zap.S().Errorf("can not reset this user")
		c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "can not reset this user"})
		return
	}
	if err := userService.ResetPwd(id); err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": err.Error()})

	}
	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "reset success."})
}

func GetUserList(c *gin.Context) {
	var err error
	var query = &request.Query{}
	if err = c.ShouldBindJSON(query); err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": err.Error()})
		return
	}

	svc, total, err := userService.GetUserList(query)
	if err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "success", "data": svc, "total": int(total)})
}
