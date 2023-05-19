package woa

import (
	"net/http"
	"wkg/model/request"
	"wkg/services/srcService/WOAService"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func GetWOAInfo(c *gin.Context) {
	var err error
	var query = &request.Query{}
	if err = c.ShouldBindJSON(query); err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": err.Error()})
		return
	}
	data, total, err := WOAService.GetWOAInfo(query)
	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "success", "data": data, "total": int(total)})
}

func GetWOAInfoByCid(c *gin.Context) {
	var err error
	var query = &request.Query{}
	if err = c.ShouldBindJSON(query); err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": err.Error()})
		return
	}

	woas, total, err := WOAService.GetWOAInfoByCid(query)
	if err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "success", "data": woas, "total": int(total)})
}

func DelWOAById(c *gin.Context) {
	id := c.Query("id")

	if err := WOAService.DelWOAById(id); err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": err.Error()})

	}
	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "delete success."})
}

func GetNewWOAInfo(c *gin.Context) {
	var err error
	var query = &request.Query{}
	if err = c.ShouldBindJSON(query); err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": err.Error()})
		return
	}

	woas, total, err := WOAService.GetNewWOAInfo(query)
	if err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "success", "data": (woas), "total": total})
}

func SearchWOAInfo(c *gin.Context) {
	var err error
	var query = &request.Query{}
	if err = c.ShouldBindJSON(query); err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": err.Error()})
		return
	}
	data, total, err := WOAService.SearchWOAInfo(query)
	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "success", "data": data, "total": int(total)})

}

func ReadFlagWOAInfoById(c *gin.Context) {
	var err error
	var pids = &request.ParamIds{}
	if err = c.ShouldBindJSON(pids); err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": err.Error()})
		return
	}
	err = WOAService.ReadFlagWOAInfoById(pids)
	if err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "success."})
}

func ReadAllFlagWOAInfo(c *gin.Context) {
	err := WOAService.ReadAllFlagWOAInfo()
	if err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "success."})
}
