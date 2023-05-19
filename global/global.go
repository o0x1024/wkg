package global

import go_queue "wkg/pkg/libs/go-queue"

var (
	HasNewCompanyFlag = false
	IPReqRecord       = make(map[string]int)

	TaskCore = make(map[string]interface{})

	Queue = go_queue.NewQueue(4096 * 2)
)

func init() {

}
