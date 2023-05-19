package blog

import (
	"net/http"
	"sort"
	"wkg/model/request"
	"wkg/services/blogService"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func GetShareknowledgeList(c *gin.Context) {
	var err error
	page := &request.PageParam{}
	if err = c.ShouldBindJSON(page); err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": err.Error()})
		return
	}

	lds, total, err := blogService.GetShareknowledgeList(page)
	if err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": err.Error()})
		return
	}

	sort.Slice(lds, func(i, j int) bool {
		return lds[i].Year > lds[j].Year
	})

	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "成功", "data": lds, "total": int(total)})
}

func GetShareknowledgeByKey(c *gin.Context) {
	key := c.Query("key")

	data, err := blogService.GetShareknowledgeByKey(key)
	if err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "成功", "data": data})
}

func SearchShareKnowledgeByWord(c *gin.Context) {
	page := &request.PageParam{}
	if err := c.ShouldBindJSON(page); err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": err.Error()})
		return
	}

	dataListString, total, err := blogService.SearchShareKnowledgeByWord(page)
	if err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "成功", "data": dataListString, "total": int(total)})
}

func GetsharedKnowledgeListbyTag(c *gin.Context) {
	page := &request.PageParam{}
	if err := c.ShouldBindJSON(page); err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": err.Error()})
		return
	}

	dataList, total, err := blogService.GetsharedKnowledgeListbyTag(page)
	if err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "成功", "data": dataList, "total": int(total)})
}

func GetTags(c *gin.Context) {
	tags, err := blogService.GetTags()
	if err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "success", "data": tags})

}
