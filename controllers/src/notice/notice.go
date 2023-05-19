package notice

import (
	"net/http"
	"wkg/model"
	"wkg/model/request"
	"wkg/services/srcService/noticeService"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func DelNoticeById(c *gin.Context) {
	id := c.Query("id")

	if err := noticeService.DelNoticeById(id); err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": err.Error()})

	}
	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "删除成功."})

}

func SaveEditNotice(c *gin.Context) {
	var err error
	notice := &model.Notice{}
	if err = c.ShouldBindJSON(notice); err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": err.Error()})
		return
	}

	err = noticeService.SaveEditNotice(notice)
	if err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "success"})
}

func SaveNotice(c *gin.Context) {
	var err error
	notice := &model.Notice{}
	if err = c.ShouldBindJSON(notice); err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": err.Error()})
		return
	}

	err = noticeService.SaveNotice(notice)
	if err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "success"})

}

func GetNoticeDetailById(c *gin.Context) {
	var err error
	var id = c.Query("id")

	svc, err := noticeService.GetNoticeDetailById(id)
	if err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "success", "data": svc})
}

func GetNoticeList(c *gin.Context) {
	var err error
	var query = &request.Query{}
	if err = c.ShouldBindJSON(query); err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": err.Error()})
		return
	}

	svc, total, err := noticeService.GetNoticeList(query)
	if err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "success", "data": svc, "total": int(total)})
}

func GetNoticeByKeyword(c *gin.Context) {
	var err error
	var query = &request.Query{}
	if err = c.ShouldBindJSON(query); err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": err.Error()})
		return
	}

	svc, total, err := noticeService.GetNoticeByKeyword(query)
	if err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "success", "data": svc, "total": int(total)})
}
