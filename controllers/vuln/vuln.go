package vuln

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
	"wkg/model"
	"wkg/model/request"
	response2 "wkg/model/response"
	db2 "wkg/pkg/db"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func GetRootDomain(c *gin.Context) {
	var cmp []model.Company
	err := db2.Orm.Model(&model.Company{}).Select("domain,monitorStatus").Find(&cmp).Error
	if err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": err.Error()})
		return
	}
	var bf response2.BugFindRes
	if len(cmp) > 0 {
		for _, v := range cmp {
			if *v.MonitorStatus {
				sd := strings.Split(v.Domain, "|")
				for _, vx := range sd {
					bf.Data = append(bf.Data, vx)
				}
			}
		}
	}

	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "成功", "data": bf})

}

func GetFofaRule(c *gin.Context) {
	var bf response2.BugFindRes
	domainRuleList := []string{"cert=\"_gelen_\""}
	keyWordRuleList := []string{"title=\"_gelen_\"", "body=\"_gelen_\""}
	//favicon := "icon_hash=\"_gelen_\""   ,"header=\"_gelen_\""
	var cmp []model.Company
	err := db2.Orm.Model(&model.Company{}).Select("domain,keyWord").Find(&cmp).Error
	if err != nil {
		zap.S().Errorf("%s", err.Error())
		log.Println("company.go line:81 error:[", err, "]")
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": err.Error()})
		return
	}

	//域名相关规则
	if len(cmp) > 0 {
		for _, v := range cmp {
			sd := strings.Split(v.Domain, "|")
			for _, vx := range sd {
				for _, iv := range domainRuleList {
					bf.Data = append(bf.Data, strings.Replace(iv, "_gelen_", vx, -1))
				}
			}

			kw := strings.Split(v.KeyWord, "|")
			for _, vx := range kw {
				if vx == "-" {
					continue
				}
				for _, iv := range keyWordRuleList {
					bf.Data = append(bf.Data, strings.Replace(iv, "_gelen_", vx, -1))
				}
			}
		}
	}

	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "成功", "data": bf})

}

func GetNewBug(c *gin.Context) {
	var err error
	var bug []model.BugReport
	var param = &request.PageParam{}
	if err = c.ShouldBindJSON(param); err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": err.Error()})
		return
	}

	err = db2.Orm.Model(&model.BugReport{}).Limit(param.PageSize).Offset((param.Page - 1) * param.PageSize).Where("isNew=true").Find(&bug).Error
	if err != nil {
		zap.S().Errorf("%s", err.Error())
		log.Println("GetNewBug.go line:116 error:[", err, "]")
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": err.Error()})
		return
	}

	var count int64
	err = db2.Orm.Model(&model.BugReport{}).Where("isNew=true").Count(&count).Error
	if err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "成功", "data": bug, "total": int(count)})
}

func GetOldBug(c *gin.Context) {
	var err error
	var bug []model.BugReport
	var param = &request.PageParam{}
	if err = c.ShouldBindJSON(param); err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": err.Error()})
		return
	}

	err = db2.Orm.Model(&model.BugReport{}).Limit(param.PageSize).Offset((param.Page - 1) * param.PageSize).Where("isNew=false").Find(&bug).Error
	if err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": err.Error()})
		return
	}

	var count int64
	err = db2.Orm.Model(&model.BugReport{}).Limit(param.PageSize).Offset((param.Page - 1) * param.PageSize).Where("isNew=false").Count(&count).Error
	if err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "成功", "data": bug, "total": int(count)})
}

func BatchRead(c *gin.Context) {
	var err error
	type DelId struct {
		Id string `json:"id"`
	}
	var ei = &DelId{}
	if err = c.ShouldBindJSON(ei); err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": err.Error()})
		return
	}
	id := strings.Split(ei.Id, ",")

	for _, v := range id {
		err := db2.Orm.Model(&model.BugReport{}).Where("id=?", v).Update("isNew", false).Error
		if err != nil {
			zap.S().Errorf("%s", err.Error())
			c.JSON(http.StatusOK, gin.H{"code": 400, "msg": err.Error()})
			return
		}
	}
	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "修改成功."})
	return
}

func DelBugCollectById(c *gin.Context) {
	id := c.Query("id")

	if err := db2.Orm.Where("id=?", id).Delete(&model.BugCollect{}).Error; err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "成功"})
	return
}

func SearchBugCollectByKey(c *gin.Context) {
	var err error
	type PageParam struct {
		Page     int    `json:"page"`
		PageSize int    `json:"pagesize"`
		KeyWord  string `json:"keyword"`
	}

	var param = &PageParam{}
	if err = c.ShouldBindJSON(param); err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": err.Error()})
		return
	}

	var count int64 = 0
	if err := db2.Orm.Model(&model.BugCollect{}).Where("detail LIKE ?", "%"+param.KeyWord+"%").Count(&count).Error; err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": err.Error()})
		return
	}

	var bug = &[]model.BugCollect{}
	if err := db2.Orm.Limit(param.PageSize).Offset((param.Page-1)*param.PageSize).Model(&model.BugCollect{}).Where("detail LIKE ?", "%"+param.KeyWord+"%").Find(&bug).Error; err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "成功", "data": bug, "total": int(count)})
}

func SearchOldBugReport(c *gin.Context) {
	var err error
	type PageParam struct {
		Page     int    `json:"page"`
		PageSize int    `json:"pagesize"`
		KeyWord  string `json:"keyword"`
	}

	var param = &PageParam{}
	if err = c.ShouldBindJSON(param); err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": err.Error()})
		return
	}

	var count int64 = 0
	if err := db2.Orm.Model(&model.BugReport{}).Where("detail LIKE ?", "%"+param.KeyWord+"%").Where("isNew=false").Count(&count).Error; err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": err.Error()})
		return
	}

	var bug = &[]model.BugReport{}
	if err := db2.Orm.Limit(param.PageSize).Offset((param.Page-1)*param.PageSize).Model(&model.BugReport{}).Where("detail LIKE ?", "%"+param.KeyWord+"%").Where("isNew=false").Find(&bug).Error; err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "成功", "data": bug, "total": int(count)})
}

func SearchNewBugReport(c *gin.Context) {
	type PageParam struct {
		Page     int    `json:"page"`
		PageSize int    `json:"pagesize"`
		KeyWord  string `json:"keyword"`
	}
	var err error
	var param = &PageParam{}
	if err = c.ShouldBindJSON(param); err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": err.Error()})
		return
	}

	var count int64 = 0
	if err := db2.Orm.Model(&model.BugReport{}).Where("detail LIKE ?", "%"+param.KeyWord+"%").Where("isNew=true").Count(&count).Error; err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": err.Error()})
		return
	}

	var bug = &[]model.BugReport{}
	if err := db2.Orm.Limit(param.PageSize).Offset((param.Page-1)*param.PageSize).Model(&model.BugReport{}).Where("detail LIKE ?", "%"+param.KeyWord+"%").Where("isNew=true").Find(&bug).Error; err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "成功", "data": bug, "total": int(count)})
}

func GetBugCollect(c *gin.Context) {
	var bug []model.BugCollect
	var param = &request.PageParam{}
	var err error
	if err = c.ShouldBindJSON(param); err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": err.Error()})
		return
	}

	err = db2.Orm.Model(&model.BugCollect{}).Limit(param.PageSize).Offset((param.Page - 1) * param.PageSize).Find(&bug).Error
	if err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": err.Error()})
		return
	}

	var count int64
	err = db2.Orm.Model(&model.BugCollect{}).Count(&count).Error
	if err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "success", "data": bug, "total": int(count)})
}

func Collect(c *gin.Context) {
	id := c.Query("id")
	br := &model.BugReport{}
	bc := &model.BugCollect{}

	err := db2.Orm.Model(&model.BugReport{}).Where("id=?", id).Find(br).Error
	if err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": err.Error()})
		return
	}

	bc.UpdateTime = br.UpdateTime
	bc.Detail = br.Detail

	err = db2.Orm.Model(&model.BugCollect{}).Create(bc).Error
	if err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "success"})
	return
}

func Read(c *gin.Context) {
	id := c.Query("id")
	err := db2.Orm.Model(&model.BugReport{}).Where("id=?", id).Update("isNew", false).Error
	if err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "success"})
	return
}

func BugReport(c *gin.Context) {
	var bug model.BugReport
	b, _ := ioutil.ReadAll(c.Request.Body)
	detail := string(b)
	t := true
	bug.Detail = detail
	bug.IsNew = &t
	bug.UpdateTime = time.Unix(time.Now().Unix(), 0).Format("2006-01-02 15:04:05")

	err := db2.Orm.Model(&model.BugReport{}).Create(&bug).Error
	if err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": err.Error()})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "success"})
		return
	}

}
