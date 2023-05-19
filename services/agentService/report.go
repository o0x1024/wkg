package agentService

import (
	"wkg/model/request"
)

// func HeartBeat(up request.UploadReq) error {
// 	var err error
// 	//心跳包，更新一下服务器状态
// 	agent := &model.Agent{}
// 	if err = db.Orm.Model(&model.Agent{}).Where("agentId=?", up.Detail.AgentId).Find(agent).Error; err != nil {
// 		return err
// 	}
// 	agent. = up.Detail.Disk
// 	agent.Cpu = up.Detail.Cpu
// 	agent.Memory = up.Detail.Memory
// 	agent.LastUpdateTime = time.Now().Format("2006-01-02 15:04:05")
// 	if err = db.Orm.Model(&model.Agent{}).Where("agentId=?", agent.AgentId).Updates(agent).Error; err != nil {
// 		return err
// 	}

// 	return nil

// }

func UploadException(up request.UploadReq) error {

	return nil
}
