package service

import (
	"net/http"
	"wkg/model/request"
	"wkg/services/srcService/SvcService"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func GetNewServiceInfo(c *gin.Context) {
	var err error
	var query = &request.Query{}
	if err = c.ShouldBindJSON(query); err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": err.Error()})
		return
	}

	data, total, err := SvcService.GetNewServiceInfo(query)
	if err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "成功", "data": data, "total": total})
}

func DelServiceById(c *gin.Context) {
	id := c.Query("id")

	if err := SvcService.DelServiceById(id); err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": err.Error()})

	}
	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "删除成功."})
}

func GetServiceInfoByCid(c *gin.Context) {
	var err error
	var query = &request.Query{}
	if err = c.ShouldBindJSON(query); err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": err.Error()})
		return
	}

	svc, total, err := SvcService.GetServiceInfoByCid(query)
	if err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "成功", "data": svc, "total": int(total)})
}

func GetServiceInfo(c *gin.Context) {
	var err error
	var query = &request.Query{}
	if err = c.ShouldBindJSON(query); err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": err.Error()})
		return
	}
	data, total, err := SvcService.GetServiceInfo(query)
	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "成功", "data": data, "total": int(total)})

}

func ReadAllFlagServiceInfo(c *gin.Context) {
	err := SvcService.ReadAllFlagServiceInfo()
	if err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "成功."})
}

func ReadFlagServiceInfoById(c *gin.Context) {
	var err error
	var pids = &request.ParamIds{}
	if err = c.ShouldBindJSON(pids); err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": err.Error()})
		return
	}
	err = SvcService.ReadFlagServiceInfoById(pids)
	if err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "成功."})
}

func SearchServiceInfo(c *gin.Context) {
	var err error
	var query = &request.Query{}
	if err = c.ShouldBindJSON(query); err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": err.Error()})
		return
	}
	data, total, err := SvcService.SearchServiceInfo(query)
	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "成功", "data": data, "total": int(total)})

}
