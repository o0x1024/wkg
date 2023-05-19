package services

import "wkg/services/taskService"

func InitService() {

	taskService.OnceInitTask()
}
