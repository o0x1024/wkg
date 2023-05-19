package dash

import (
	"net/http"
	"wkg/model"
	db2 "wkg/pkg/db"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func GetDashboardInfo(c *gin.Context) {
	var TotalDomain int64
	err := db2.Orm.Model(&model.Domain{}).Count(&TotalDomain).Error
	if err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": "json unmarshal error"})
		return
	}

	var NewDomain int64
	err = db2.Orm.Model(&model.Domain{}).Where("isNew=true").Count(&NewDomain).Error
	if err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": "json unmarshal error"})
		return
	}

	var TotalWebwite int64
	err = db2.Orm.Model(&model.Websites{}).Count(&TotalWebwite).Error
	if err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": "json unmarshal error"})
		return
	}

	var NewWebsite int64
	err = db2.Orm.Model(&model.Websites{}).Where("isNew=true").Count(&NewWebsite).Error
	if err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": "json unmarshal error"})
		return
	}

	var TotalIp int64
	err = db2.Orm.Model(&model.Ips{}).Count(&TotalIp).Error
	if err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": "json unmarshal error"})
		return
	}

	var NewIp int64
	err = db2.Orm.Model(&model.Ips{}).Where("isNew=true").Count(&NewIp).Error
	if err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": "json unmarshal error"})
		return
	}

	var TotalService int64
	err = db2.Orm.Model(&model.Services{}).Count(&TotalService).Error
	if err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": "json unmarshal error"})
		return
	}

	var NewService int64
	err = db2.Orm.Model(&model.Services{}).Where("isNew=true").Count(&NewService).Error
	if err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": "json unmarshal error"})
		return
	}

	var TotalMini int64
	err = db2.Orm.Model(&model.MiniProgram{}).Count(&TotalMini).Error
	if err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": "json unmarshal error"})
		return
	}

	var NewMini int64
	err = db2.Orm.Model(&model.MiniProgram{}).Where("isNew=true").Count(&NewMini).Error
	if err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": "json unmarshal error"})
		return
	}

	var TotalWOA int64
	err = db2.Orm.Model(&model.WebChatOfficeAccount{}).Count(&TotalWOA).Error
	if err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": "json unmarshal error"})
		return
	}

	var NewWOA int64
	err = db2.Orm.Model(&model.WebChatOfficeAccount{}).Where("isNew=true").Count(&NewWOA).Error
	if err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": "json unmarshal error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "成功", "data": gin.H{
		"totalWebSite": TotalWebwite, "newWebSite": NewWebsite, "totalDomain": TotalDomain, "newDomain": NewDomain,
		"totalIP": TotalIp, "newIP": NewIp, "totalService": TotalService, "newService": NewService,
		"totalMini": TotalMini, "newMini": NewMini, "totalWOA": TotalWOA, "newWOA": NewWOA,
	}})
}
