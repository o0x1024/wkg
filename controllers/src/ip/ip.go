package ip

import (
	"net/http"
	"wkg/model/request"
	"wkg/services/srcService/IPService"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func DelIPById(c *gin.Context) {
	id := c.Query("id")

	if err := IPService.DelIPById(id); err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": err.Error()})

	}
	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "删除成功."})

}

func GetIPsInfoByCid(c *gin.Context) {
	var err error
	var query = &request.Query{}
	if err = c.ShouldBindJSON(query); err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": err.Error()})
		return
	}

	ips, total, err := IPService.GetIPsInfoByCid(query)
	if err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "成功", "data": ips, "total": int(total)})

}

func GetIPsInfo(c *gin.Context) {
	var err error
	var query = &request.Query{}
	if err = c.ShouldBindJSON(query); err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": err.Error()})
		return
	}
	data, total, err := IPService.GetIPsInfo(query)
	if err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "成功", "data": data, "total": total})
}

func ReadFlagIPInfoById(c *gin.Context) {

	var err error
	var pids = &request.ParamIds{}
	if err = c.ShouldBindJSON(pids); err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": err.Error()})
		return
	}
	err = IPService.ReadFlagIPInfoById(pids)
	if err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "成功."})
}

func ReadAllFlagIPInfo(c *gin.Context) {
	err := IPService.ReadAllFlagIPInfo()
	if err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "成功."})
}

func SearchIPsInfo(c *gin.Context) {
	var err error
	var query = &request.Query{}
	if err = c.ShouldBindJSON(query); err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": err.Error()})
		return
	}

	data, total, err := IPService.SearchIPsInfo(query)
	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "成功", "data": data, "total": total})
	return
}

func GetNewIPsInfo(c *gin.Context) {
	var err error
	var query = &request.Query{}
	if err = c.ShouldBindJSON(query); err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": err.Error()})
		return
	}

	data, total, err := IPService.GetNewIPsInfo(query)
	if err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "成功", "data": data, "total": total})
}
