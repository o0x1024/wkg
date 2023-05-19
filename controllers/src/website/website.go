package website

import (
	"net/http"
	"wkg/model"
	"wkg/model/request"
	db2 "wkg/pkg/db"
	"wkg/services/srcService/websiteService"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func ScanNewWebSite(c *gin.Context) {
	err := websiteService.ScanNewWebsiteInfo()
	if err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": err.Error()})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "scanning"})
		return
	}
}

func GetWebSiteInfoByCid(c *gin.Context) {
	var err error
	var param = &request.Query{}
	if err = c.ShouldBindJSON(param); err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": err.Error()})
		return
	}

	website, total, err := websiteService.GetWebSiteInfoByCId(param)
	if err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "成功", "data": website, "total": total})
}

func GetNewWebSiteInfo(c *gin.Context) {
	var err error
	var query = &request.Query{}
	if err = c.ShouldBindJSON(query); err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": err.Error()})
		return
	}

	var count int64
	//根据page和pagesize查询数据
	var dom []model.Websites
	err = db2.Orm.Model(&model.Websites{}).Where("isNew=true").Count(&count).Limit(query.PageSize).Offset((query.Page - 1) * query.PageSize).Find(&dom).Error
	if err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "成功", "data": dom, "total": int(count)})
}

func DelWebSiteById(c *gin.Context) {
	id := c.Query("id")

	if err := websiteService.DelWebSiteById(id); err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": err.Error()})

	}
	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "删除成功."})
}

func GetWebSiteInfoByKey(c *gin.Context) {
	var err error
	var dss = &request.SearchStrut{}
	if err = c.ShouldBindJSON(dss); err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": err.Error()})
		return
	}

	website, total, err := websiteService.GetWebSiteInfoByKey(dss)
	if err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "成功", "data": website, "total": total})
}

func ReadFlagWebSitefoById(c *gin.Context) {
	var err error
	var pids = &request.ParamIds{}
	if err = c.ShouldBindJSON(pids); err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": err.Error()})
		return
	}
	err = websiteService.ReadFlagWebSitefoById(pids)
	if err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "成功"})
}

func ReadAllFlagWebSiteInfo(c *gin.Context) {
	err := db2.Orm.Model(&model.Websites{}).Where("1=1").Update("isNew", false).Error
	if err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": "db error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "成功"})
	return
}

func GetWebSiteInfo(c *gin.Context) {
	var err error
	var query = &request.Query{}
	if err = c.ShouldBindJSON(query); err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": "unmarshal error"})
		return
	}
	website, total, err := websiteService.GetWebSiteInfo(query)
	if err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": "db error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "成功", "data": website, "total": int(total)})
	return
}
