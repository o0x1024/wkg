package company

import (
	"net/http"
	"wkg/model"
	"wkg/model/request"
	db2 "wkg/pkg/db"
	"wkg/services/srcService/companyService"
	"wkg/services/srcService/domainService"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func GetCompanyInfo(c *gin.Context) {
	var err error
	var param = &request.Query{}
	if err = c.ShouldBindJSON(param); err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": err.Error()})
		return
	}
	dom, total, err := companyService.GetCompanyInfo(param)
	if err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "成功", "data": dom, "total": int(total)})
}

func scanCompanyInfo(c *gin.Context) {
	var err error
	type TId struct {
		Id int `json:"id"`
	}
	var Tid = &TId{}
	if err = c.ShouldBindJSON(Tid); err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": err.Error()})
		return
	}
	var cmp model.Company
	err = db2.Orm.Model(&model.Company{}).Where("id=?", Tid.Id).Find(&cmp).Error
	if err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": err.Error()})
		return
	}
	if len(cmp.Domain) > 0 {
		go domainService.ScanDomain(cmp)
	} else {
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "scanning"})
	return

}

func GetCompanyByKey(c *gin.Context) {
	var err error
	var dss = &request.SearchStrut{}
	if err = c.ShouldBindJSON(dss); err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": err.Error()})
		return
	}
	dom, total, err := companyService.GetCompanyByKey(dss)
	if err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "成功", "data": dom, "total": int(total)})
}

func NewCompanyInfo(c *gin.Context) {
	var err error
	var com = &model.Company{}
	if err = c.ShouldBindJSON(com); err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": err.Error()})
		return
	}
	err = companyService.NewCompanyInfo(com)
	if err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "创建成功"})

}

func DelCompanyByCId(c *gin.Context) {
	cid := c.Query("cid")
	if cid == "" {
		zap.S().Errorf("param is null")
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": "param is null"})
		return
	}
	err := companyService.DelCompanyByCId(cid)
	if err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "删除成功"})
	return
}

func GetCompanyByCId(c *gin.Context) {

	cid := c.Query("cid")
	if cid == "" {
		zap.S().Errorf("param is null")
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": "param is null"})
		return
	}

	cmp, err := companyService.GetCompanyById(cid)
	if err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "成功", "data": cmp})
}

func UpdateCompanyInfo(c *gin.Context) {
	var err error
	var tc = &model.Company{}
	if err = c.ShouldBindJSON(tc); err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": err.Error()})
		return
	}

	err = companyService.UpdateCompanyInfo(tc)
	if err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "成功"})
}

func GetSelectOption(c *gin.Context) {

	option, err := companyService.GetSelectOption()
	if err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "成功", "data": option})
}
