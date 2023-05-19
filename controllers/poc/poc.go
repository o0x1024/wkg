package poc

import (
	"net/http"
	"wkg/model"
	"wkg/model/request"
	"wkg/services/pocService"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func DelPocById(c *gin.Context) {
	id := c.Query("id")

	if err := pocService.DelPocById(id); err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": err.Error()})

	}
	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "删除成功."})

}

func SaveEditPoc(c *gin.Context) {
	var err error
	Poc := &model.Poc{}
	if err = c.ShouldBindJSON(Poc); err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": err.Error()})
		return
	}

	err = pocService.SaveEditPoc(Poc)
	if err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "success"})
}

func SavePoc(c *gin.Context) {
	var err error
	Poc := &model.Poc{}
	if err = c.ShouldBindJSON(Poc); err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": err.Error()})
		return
	}

	err = pocService.SavePoc(Poc)
	if err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "success"})

}

func GetPocDetailById(c *gin.Context) {
	var err error
	var id = c.Query("id")

	svc, err := pocService.GetPocDetailById(id)
	if err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "success", "data": svc})
}

func GetPocInfo(c *gin.Context) {
	var err error
	var query = &request.Query{}
	if err = c.ShouldBindJSON(query); err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": err.Error()})
		return
	}

	svc, total, err := pocService.GetPocInfo(query)
	if err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "success", "data": svc, "total": int(total)})
}

func GetPocInfoByCid(c *gin.Context) {
	var err error
	var query = &request.Query{}
	if err = c.ShouldBindJSON(query); err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": err.Error()})
		return
	}

	svc, total, err := pocService.GetPocInfoByCid(query)
	if err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "success", "data": svc, "total": int(total)})
}
