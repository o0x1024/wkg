package minipro

import (
	"net/http"
	"wkg/model/request"
	"wkg/services/srcService/miniprogramService"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func GetMiniInfo(c *gin.Context) {
	var err error
	var query = &request.Query{}
	if err = c.ShouldBindJSON(query); err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": err.Error()})
		return
	}
	data, total, err := miniprogramService.GetMiniInfo(query)
	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "success", "data": data, "total": int(total)})
}

func GetMiniInfoByCid(c *gin.Context) {
	var err error
	var query = &request.Query{}
	if err = c.ShouldBindJSON(query); err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": err.Error()})
		return
	}

	svc, total, err := miniprogramService.GetMiniInfoByCid(query)
	if err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "success", "data": svc, "total": int(total)})
}

func DelMiniById(c *gin.Context) {
	id := c.Query("id")

	if err := miniprogramService.DelMiniById(id); err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": err.Error()})

	}
	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "delete success."})
}

func GetNewMiniInfo(c *gin.Context) {
	var err error
	var query = &request.Query{}
	if err = c.ShouldBindJSON(query); err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": err.Error()})
		return
	}

	data, total, err := miniprogramService.GetNewMiniInfo(query)
	if err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "success", "data": data, "total": total})
}

func SearchMiniInfo(c *gin.Context) {
	var err error
	var query = &request.Query{}
	if err = c.ShouldBindJSON(query); err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": err.Error()})
		return
	}
	data, total, err := miniprogramService.SearchMiniInfo(query)
	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "success", "data": data, "total": int(total)})

}

func ReadFlagMiniInfoById(c *gin.Context) {
	var err error
	var pids = &request.ParamIds{}
	if err = c.ShouldBindJSON(pids); err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": err.Error()})
		return
	}
	err = miniprogramService.ReadFlagMiniInfoById(pids)
	if err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "success."})
}

func ReadAllFlagMiniInfo(c *gin.Context) {
	err := miniprogramService.ReadAllFlagMiniInfo()
	if err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "success."})
}
