package company

import (
	"net/http"
	"os"
	"strconv"
	"wkg/model"
	"wkg/model/request"
	"wkg/pkg/db"
	db2 "wkg/pkg/db"
	"wkg/services/srcService/companyService"
	"wkg/services/srcService/domainService"

	"github.com/gin-gonic/gin"
	"github.com/tealeg/xlsx"
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

type XID struct {
	Id int `json:"id"`
}

func Export(c *gin.Context) {
	var xid XID
	var sheet *xlsx.Sheet
	var row *xlsx.Row
	var cell *xlsx.Cell
	var err error

	if err = c.ShouldBindJSON(&xid); err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": err.Error()})
		return
	}

	wb := xlsx.NewFile()
	myStyle := xlsx.NewStyle()
	myStyle.Font.Size = 11
	myStyle.Font.Name = "Microsoft YaHei UI Light"
	//取出域名信息
	var dom []model.Domain
	err = db.Orm.Model(&model.Domain{}).Where("cid=?", xid.Id).Find(&dom).Error
	if err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": err.Error()})
		return
	}
	sheet, err = wb.AddSheet("domain")
	err = sheet.SetColWidth(20, 20, 100)
	if err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": err.Error()})
		return
	}
	row = sheet.AddRow()
	cell = row.AddCell()
	cell.SetStyle(myStyle)
	cell.Value = "id"
	cell = row.AddCell()
	cell.SetStyle(myStyle)
	cell.Value = "cid"
	cell = row.AddCell()
	cell.SetStyle(myStyle)
	cell.Value = "domain"
	cell = row.AddCell()
	cell.SetStyle(myStyle)
	cell.Value = "type"
	cell = row.AddCell()
	cell.SetStyle(myStyle)
	cell.Value = "updateTime"
	cell = row.AddCell()
	cell.SetStyle(myStyle)
	cell.Value = "isNew"

	for _, v := range dom {
		row = sheet.AddRow()
		cell = row.AddCell()
		cell.SetStyle(myStyle)
		cell.Value = strconv.Itoa(v.Id)
		cell = row.AddCell()
		cell.SetStyle(myStyle)
		cell.Value = strconv.Itoa(v.Cid)
		cell = row.AddCell()
		cell.SetStyle(myStyle)
		cell.Value = v.Domain
		cell = row.AddCell()
		cell.SetStyle(myStyle)
		cell.Value = v.Type
		cell = row.AddCell()
		cell.SetStyle(myStyle)
		cell.Value = v.UpdateTime
		cell = row.AddCell()
		cell.SetStyle(myStyle)
		if *v.IsNew {
			cell.Value = "true"
		}
	}
	//取出域名信息
	var ws []model.Websites
	err = db.Orm.Model(&model.Websites{}).Where("cid=?", xid.Id).Find(&ws).Error
	if err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": err.Error()})
		return
	}
	sheet, err = wb.AddSheet("website")

	if err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": err.Error()})
		return
	}
	{
		row = sheet.AddRow()
		cell = row.AddCell()
		cell.SetStyle(myStyle)
		cell.Value = "#"
		cell = row.AddCell()
		cell.SetStyle(myStyle)
		cell.Value = "Website"
		cell = row.AddCell()
		cell.SetStyle(myStyle)
		cell.Value = "Title"
		cell = row.AddCell()
		cell.SetStyle(myStyle)
		cell.Value = "Ip"
		cell = row.AddCell()
		cell.SetStyle(myStyle)
		cell.Value = "Favicon"
		cell = row.AddCell()
		cell.SetStyle(myStyle)
		cell.Value = "Headers"
		cell = row.AddCell()
		cell.SetStyle(myStyle)
		cell.Value = "UpdateTime"
		cell = row.AddCell()
		cell.SetStyle(myStyle)
		cell.Value = "CDN"
		cell = row.AddCell()
		cell.SetStyle(myStyle)
		cell.Value = "Cert"
		for _, v := range ws {
			row = sheet.AddRow()
			cell = row.AddCell()
			cell.SetStyle(myStyle)
			cell.Value = strconv.Itoa(v.Id)
			cell = row.AddCell()
			cell.SetStyle(myStyle)
			cell.Value = strconv.Itoa(v.Cid)
			cell = row.AddCell()
			cell.SetStyle(myStyle)
			cell.Value = v.Domain
			cell = row.AddCell()
			cell.SetStyle(myStyle)
			cell.Value = v.Website
			cell = row.AddCell()
			cell.SetStyle(myStyle)
			cell.Value = v.Ips
			cell = row.AddCell()
			cell.SetStyle(myStyle)
			cell.Value = v.Favicon
			cell = row.AddCell()
			cell.SetStyle(myStyle)
			cell.Value = v.FaviconUrl
			cell = row.AddCell()
			cell.SetStyle(myStyle)
			cell.Value = v.Title
			cell = row.AddCell()
			cell.SetStyle(myStyle)
			cell.Value = v.Headers
			cell = row.AddCell()
			cell.SetStyle(myStyle)
			cell.Value = v.UpdateTime
			cell = row.AddCell()
			cell.SetStyle(myStyle)
			if *v.CDN {
				cell.Value = "Y"
			}
			cell = row.AddCell()
			cell.SetStyle(myStyle)
			cell.Value = v.Cert

		}
	}
	path := "tmp/result.xlsx"
	err = wb.Save(path)
	if err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": err.Error()})
		return
	}

	content, err := os.ReadFile("tmp/result.xlsx")
	if err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": err.Error()})
		return
	}

	c.Writer.WriteHeader(http.StatusOK)
	c.Header("Content-Disposition", `attachment; filename=result.xlsx`)
	c.Header("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet; "+"filename=result.xlsx")
	c.Writer.Write(content)
}
