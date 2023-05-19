package tool

import (
	"net/http"
	"wkg/model/request"
	"wkg/services/toolService"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func SendQuestion(c *gin.Context) {

	gpt := &request.ChatGPT{}
	if err := c.ShouldBindJSON(gpt); err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": err.Error()})
		return
	}

	request, err := toolService.SendQuestion(gpt)
	if err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "success", "data": request})
}
