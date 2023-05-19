package report

import (
	"github.com/gin-gonic/gin"
)

/*
	Type

数据类型，可选择：
1 - Agent心跳数据
2 - 攻击信息
3 - 异常信息
*/
func Upload(c *gin.Context) {
	// var up = request.UploadReq{}
	// var err error

	// err = c.BindJSON(&up)
	// if err != nil {
	// 	zap.S().Errorf("%s", err.Error())
	// 	c.JSON(400, gin.H{"msg": "param error."})
	// 	return
	// }
	// if 1 == up.Type {
	// 	err := agentService.HeartBeat(up)
	// 	if err != nil {
	// 		zap.S().Errorf("%s", err.Error())
	// 		c.JSON(400, gin.H{"msg": "param error."})
	// 		return
	// 	}

	// 	c.JSON(200, gin.H{"msg": "heartbeat", "status": 201})
	// 	return
	// }

	// if 2 == up.Type {

	// }
	// if 3 == up.Type {

	// }

}
