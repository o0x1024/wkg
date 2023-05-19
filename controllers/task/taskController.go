package task

import (
	"net/http"
	"wkg/global"
	"wkg/model"
	"wkg/model/request"
	"wkg/services/taskService"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func NewTask(c *gin.Context) {
	task := &model.Task{}
	if err := c.ShouldBindJSON(task); err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": err.Error()})
		return
	}

	err := taskService.NewTask(task)
	if err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": err.Error()})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "任务新建成功."})
		return
	}
}

func UpdateTask(c *gin.Context) {
	task := &model.Task{}
	if err := c.ShouldBindJSON(task); err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": err.Error()})
		return
	}

	err := taskService.UpdateTask(task)
	if err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": err.Error()})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "任务新建成功."})
		return
	}
}

func TaskInfo(c *gin.Context) {
	tid := c.Query("taskId")
	if tid == "" {
		zap.S().Errorf("param is null")
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": "param is null"})
		return
	}

	task, err := taskService.TaskInfo(tid)
	if err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "success", "data": task})
}

func GetTaskList(c *gin.Context) {
	query := &request.Query{}
	if err := c.ShouldBindJSON(query); err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": err.Error()})
		return
	}

	tasklist, total, err := taskService.GetTaskList(query)
	if err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": err.Error()})
		return
	}
	queueNum := global.Queue.Quantity()
	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "success", "data": tasklist, "queueNum": queueNum, "total": int(total)})
}

func GetTask(c *gin.Context) {
	task, ok := taskService.GetTask()
	if !ok {
		zap.S().Errorf("no task")
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": "no task"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "成功", "data": task})

}

func StopTask(c *gin.Context) {
	tid := c.Query("taskId")
	if tid == "" {
		zap.S().Errorf("param is null")
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": "param is null"})
		return
	}
	err := taskService.StopTask(tid)
	if err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "暂停成功"})

}

func RunTask(c *gin.Context) {
	tid := c.Query("taskId")
	if tid == "" {
		zap.S().Errorf("param is null")
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": "param is null"})
		return
	}
	if err := taskService.RunTask(tid); err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "运行成功"})

}

func DelTaskByTaskId(c *gin.Context) {
	tid := c.Query("taskId")
	if tid == "" {
		zap.S().Errorf("param is null")
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": "param is null"})
		return
	}
	if err := taskService.DelTaskByTaskId(tid); err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "删除成功"})

}

func GetCompIds(c *gin.Context) {

	options, err := taskService.GetCompIds()
	if err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "成功", "data": options})

}

func GetPocs(c *gin.Context) {
	options, err := taskService.GetPocs()
	if err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "成功", "data": options})
}

func ClsTaskQueue(c *gin.Context) {
	taskService.ClsTaskQueue()
	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "成功"})
}

func UploadDomains(c *gin.Context) {

	dsr := request.DomainScanRes{}
	if err := c.ShouldBindJSON(dsr); err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": err.Error()})
		return
	}

	err := taskService.UploadDomains(dsr)
	if err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "成功"})

}
