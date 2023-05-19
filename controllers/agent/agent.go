package agent

import (
	"time"
	"wkg/model"
	"wkg/model/request"
	"wkg/model/response"
	db2 "wkg/pkg/db"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func HeartBeat(c *gin.Context) {
	var err error
	var agent = model.Agent{}
	var res = response.ResBase{}
	err = c.BindJSON(&agent)
	if err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(400, gin.H{"code": "400", "msg": "param error."})
		return
	}
	//更新时间
	agent.UpdateTime = time.Now().Format("2006-01-02 15:04:05")
	var count int64
	if err = db2.Orm.Model(&model.Agent{}).Where("agent_id=?", agent.AgentId).Count(&count).Error; err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(400, gin.H{"code": "400", "msg": "db error."})
		return
	}
	if count == 0 {

		err = db2.Orm.Model(&model.Agent{}).Create(&agent).Error
		if err != nil {
			zap.S().Errorf("%s", err.Error())
			c.JSON(400, gin.H{"code": "400", "msg": "param error."})
			return
		}
		res.Status = 201
		res.Msg = "success"
		zap.S().Info("agent register success,agentID:%s", agent.AgentId)
		c.JSON(200, res)
		return
	} else {

		err = db2.Orm.Model(&model.Agent{}).Where("agent_id=?", agent.AgentId).Updates(&agent).Error
		if err != nil {
			zap.S().Errorf("%s", err.Error())
			c.JSON(400, gin.H{"code": "400", "msg": "db error."})
			return
		}
		res.Status = 201
		res.Msg = "success"
		c.JSON(200, res)
		return
	}

}

const DEFAULT_OFFLINE_DURATION = 10 * 60

func AgentList(c *gin.Context) {
	var err error
	agentList := []model.Agent{}
	page := request.PageParam{}

	type RespAgent struct {
		AgentId    string `json:"agent_id"`
		Hostname   string `json:"host_name"`
		Platform   string `json:"platform" `
		Status     string `json:"status"`
		CpuUsage   string `json:"cpu_usage" `
		MemUsage   string `json:"mem_usage" `
		UpdateTime string `json:"update_time"`
	}

	ras := []RespAgent{}

	if err = c.BindJSON(&page); err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(400, gin.H{"msg": "param error."})
		return
	}
	if err = db2.Orm.Model(&model.Agent{}).Limit(page.PageSize).Offset((page.Page - 1) * page.PageSize).Find(&agentList).Error; err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(400, gin.H{"msg": "database error."})
		return
	}

	for _, v := range agentList {
		ra := RespAgent{}

		ra.AgentId = v.AgentId
		ra.Hostname = v.Hostname
		ra.Platform = v.Platform
		ra.CpuUsage = v.CpuUsage[2:4] + "%"
		ra.MemUsage = v.MemUsage[2:4] + "%"
		ra.UpdateTime = v.UpdateTime

		now := time.Now().Unix()
		loc, _ := time.LoadLocation("Local")
		uptime, err := time.ParseInLocation("2006-01-02 15:04:05", v.UpdateTime, loc)
		if err != nil {
			zap.S().Error(err)
		}

		up := uptime.Local().Unix()
		subnum := now - up
		if subnum > DEFAULT_OFFLINE_DURATION {
			ra.Status = "offline"
		} else {
			ra.Status = "running"
		}
		ras = append(ras, ra)
	}

	var count int64
	if err = db2.Orm.Model(&model.Agent{}).Count(&count).Error; err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(400, gin.H{"msg": "database error."})
		return
	}
	c.JSON(200, gin.H{"msg": "success", "data": ras, "total": count})

}

func GetAgentById(c *gin.Context) {
	id := c.Query("id")
	var agent = model.Agent{}
	if id != "" {
		if err := db2.Orm.Model(model.Agent{}).Where("agent_id=?", id).Find(&agent).Error; err != nil {
			zap.S().Errorf("%s", err.Error())
			c.JSON(400, gin.H{"msg": "Database delete error."})
			return
		}
	}
	c.JSON(200, gin.H{"msg": "success", "data": agent})
}

func DeleteAgent(c *gin.Context) {
	id := c.Query("id")
	if id != "" {
		var agent = &model.Agent{}
		agent.AgentId = id
		if err := db2.Orm.Model(&model.Agent{}).Delete(&agent).Error; err != nil {
			zap.S().Errorf("%s", err.Error())
			c.JSON(400, gin.H{"msg": "Database delete error."})
			return
		}
		c.JSON(200, gin.H{"msg": "success"})
	} else {
		c.JSON(400, gin.H{"msg": "can not found id"})

	}
}

func AgentDetail(c *gin.Context) {
	id := c.Query("agent_id")
	if id != "" {
		var agent = model.Agent{}
		if err := db2.Orm.Where("agent_id=?", id).Find(&agent).Error; err != nil {
			zap.S().Errorf("%s", err.Error())
			c.JSON(400, gin.H{"msg": "Database delete error."})
			return
		}
		c.JSON(200, gin.H{"msg": "success", "data": agent})
	} else {
		c.JSON(400, gin.H{"msg": "can not found id"})

	}
}
