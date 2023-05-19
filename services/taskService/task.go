package taskService

import (
	"context"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
	"strings"
	"time"
	"wkg/global"
	"wkg/model"
	"wkg/model/request"
	"wkg/model/response"
	"wkg/model/response/task"
	"wkg/pkg/db"
	go_queue "wkg/pkg/libs/go-queue"
	"wkg/utils"

	"github.com/pkg/errors"
	"go.uber.org/zap"
)

func NewTask(task *model.Task) error {
	var err error
	//先判断任务名称是否重复
	var count int64
	if err = db.Orm.Model(&model.Task{}).Where("taskname=?", task.TaskName).Count(&count).Error; err != nil {
		return err
	}
	if count >= 1 {
		return errors.New("任务名称重复")
	}
	//1.添加到数据库
	task.TaskId = fmt.Sprintf("%x", md5.Sum([]byte(time.Unix(time.Now().Unix(), 0).Format("2006-01-02 15:04:05"))))
	task.UpdateTime = time.Unix(time.Now().Unix(), 0).Format("2006-01-02 15:04:05")

	poc := &model.Poc{}
	if err := db.Orm.Model(&model.Poc{}).Where("pocName=?", task.PocName).Find(&poc).Error; err != nil {
		return errors.New("数据库查询错误")
	}
	task.PocId = poc.Id
	err = db.Orm.Model(&model.Task{}).Create(&task).Error
	if err != nil {
		return err
	}

	//如果任务状态为停止，则不启动任务
	if !*task.TaskStatus {
		return nil
	}

	tastContext, cancel := context.WithCancel(context.Background())
	global.TaskCore[task.TaskId] = cancel
	//，2.添加实际的任务
	go func(ctx context.Context, tm *model.Task) {
		zap.S().Info("task id:", tm.TaskId, " start.")
		tick := time.NewTicker(time.Duration(tm.Rate) * time.Minute)
		for {
			select {
			case <-tick.C:
				newtask(tm)
			case <-ctx.Done():
				zap.S().Info("task id:", tm.TaskId, " shutdown.")
				return
			}
		}
	}(tastContext, task)
	return nil
}

func UpdateTask(task *model.Task) error {
	var err error

	//1.添加到数据库
	task.UpdateTime = time.Unix(time.Now().Unix(), 0).Format("2006-01-02 15:04:05")

	poc := &model.Poc{}
	if err := db.Orm.Model(&model.Poc{}).Where("pocName=?", task.PocName).Find(&poc).Error; err != nil {
		return errors.New("数据库查询错误")
	}
	task.PocId = poc.Id
	err = db.Orm.Model(&model.Task{}).Where("taskId=?", task.TaskId).Updates(&task).Error
	if err != nil {
		return err
	}
	//先停止任务
	if _, ok := global.TaskCore[task.TaskId]; ok {
		cancel := global.TaskCore[task.TaskId].(context.CancelFunc)
		cancel()
		zap.S().Info("stop task id:%s", task.TaskId)
	} else {
		zap.S().Info("task not exist or not running")
		return errors.New("task not exist or not running")
	}

	//重新开始任务
	tastContext, cancel := context.WithCancel(context.Background())
	global.TaskCore[task.TaskId] = cancel
	//，2.添加实际的任务
	go func(ctx context.Context, tm *model.Task) {
		zap.S().Info("task id:", tm.TaskId, " start.")
		tick := time.NewTicker(time.Duration(tm.Rate) * time.Minute)
		for {
			select {
			case <-tick.C:
				newtask(tm)
			case <-ctx.Done():
				zap.S().Info("task id:", tm.TaskId, " shutdown.")
				return
			}
		}
	}(tastContext, task)

	return nil
}

func GetTask() (interface{}, bool) {

	v, ok, _ := global.Queue.Get()
	if ok {
		// t := v.(*task.ResTask)
		return v, ok
	}
	return nil, ok
}

func GetCompIds() ([]response.Option, error) {

	cmp := []model.Company{}

	err := db.Orm.Model(&model.Company{}).Find(&cmp).Error
	if err != nil {
		zap.S().Errorf("%s", err.Error())
		return nil, nil
	}
	var option []response.Option
	for _, v := range cmp {
		option = append(option, response.Option{Value: strconv.Itoa(v.Id), Lable: v.CompanyName})
	}

	return option, nil
}

func ClsTaskQueue() {
	global.Queue = go_queue.NewQueue(4096 * 2)
}

func GetPocs() ([]response.Option, error) {

	poc := []model.Poc{}

	err := db.Orm.Model(&model.Poc{}).Find(&poc).Error
	if err != nil {
		return nil, nil
	}
	var option []response.Option
	for _, v := range poc {
		option = append(option, response.Option{Value: v.PocName, Lable: v.PocName})
	}

	return option, nil
}

// taskid
// 1000 信息搜集-域名-全部方式
// 1001 信息搜集-域名-接口扫描
// 1002 信息搜集-域名-暴力破解

// 2001 漏洞扫描-HTTP漏洞扫描

func newtask(tm *model.Task) {

	var err error
	rt := &task.ResTask{}

	if tm.TaskType == "0" {
		if tm.AssetRange == "1" && tm.DomainCollectionType == "0" {
			rt.TaskNo = 1000
		}
		if tm.AssetRange == "1" && tm.DomainCollectionType == "1" {
			rt.TaskNo = 1001
		}
		if tm.AssetRange == "1" && tm.DomainCollectionType == "2" {
			rt.TaskNo = 1002
		}

		if tm.AssetRange == "1" && tm.CompanyRange == "1" { //指定公司，搜集域名
			cids := strings.Split(tm.CompanyId, ",")
			cids = cids[:len(cids)-1]
			for _, v := range cids {
				cmp := &model.Company{}
				if err = db.Orm.Model(&model.Company{}).Where("id=?", v).Find(cmp).Error; err != nil {
					zap.S().Errorf(err.Error())
					return
				}

				domains := strings.Split(cmp.Domain, "|")
				dm := task.Domain{}
				dm.CId = cmp.Id
				dm.RootDomains = domains
				rt.Domains = append(rt.Domains, dm)
			}
			rt.TaskId = tm.TaskId

		} else if tm.AssetRange == "1" && tm.CompanyRange == "0" { //全部公司
			cmps := []model.Company{}
			if err = db.Orm.Model(&model.Company{}).Find(&cmps).Error; err != nil {
				zap.S().Errorf(err.Error())
				return
			}
			for _, v := range cmps {
				domains := strings.Split(v.Domain, "|")
				dm := task.Domain{}
				dm.CId = v.Id
				dm.RootDomains = domains

				rt.Domains = append(rt.Domains, dm)
			}
			rt.TaskId = tm.TaskId

		}
	} else if tm.TaskType == "1" {
		rt.TaskNo = 2001 //漏洞扫描-HTTP漏洞扫描
		poc := &model.Poc{}
		if err := db.Orm.Model(&model.Poc{}).Where("id=?", tm.PocId).Find(poc).Error; err != nil {
			zap.S().Info("%s", err.Error())
			return
		}

		poc.PocContent = url.QueryEscape(poc.PocContent)

		//扫描站点
		if tm.CompanyRange == "0" && tm.AssetRange == "2" { //全部公司   扫描IP
			ips := []model.Ips{}
			if err = db.Orm.Model(&model.Ips{}).Select("ip").Find(&ips).Error; err != nil {
				zap.S().Errorf(err.Error())
				return
			}

			vs := task.VSAsset{}
			vs.PocContent = poc.PocContent
			count := 0 //500个目标做为一个任务
			for _, v := range ips {
				vs.Targets = append(vs.Targets, v.Ip)
				if count > 500 {
					rt.VSAssets = vs
					rt.TaskId = tm.TaskId
					ok, _ := global.Queue.Put(rt)
					if !ok {
						zap.S().Info("put failed!")
						return
					}
				}
				count++
			}
			if count != 0 {
				rt.TaskId = tm.TaskId
				rt.VSAssets = vs
				ok, _ := global.Queue.Put(rt)
				if !ok {
					zap.S().Info("put failed!")
					return
				}
			}

		} else if tm.CompanyRange == "0" && tm.AssetRange == "0" { //全部公司   站点和IP

			ips := []model.Ips{}
			if err = db.Orm.Model(&model.Ips{}).Select("ip").Find(&ips).Error; err != nil {
				zap.S().Errorf(err.Error())
				return
			}

			vs := task.VSAsset{}
			vs.PocContent = poc.PocContent
			count := 0 //500个目标做为一个任务
			for _, v := range ips {
				vs.Targets = append(vs.Targets, v.Ip)
				if count > 500 {
					rt.VSAssets = vs
					rt.TaskId = tm.TaskId
					ok, _ := global.Queue.Put(rt)
					if !ok {
						zap.S().Info("put failed!")
						return
					}
				}
				count++
			}
			if count != 0 {
				rt.TaskId = tm.TaskId
				rt.VSAssets = vs
				ok, _ := global.Queue.Put(rt)
				if !ok {
					zap.S().Info("put failed!")
					return
				}
			}

			//---------------------------
			website := []model.Websites{}
			if err = db.Orm.Model(&model.Websites{}).Select("website").Find(&website).Error; err != nil {
				zap.S().Errorf(err.Error())
				return
			}

			vs.PocContent = poc.PocContent
			count = 0 //500个目标做为一个任务
			for _, v := range website {
				vs.Targets = append(vs.Targets, v.Website)
				if count > 500 {
					rt.VSAssets = vs
					rt.TaskId = tm.TaskId
					ok, _ := global.Queue.Put(rt)
					if !ok {
						zap.S().Info("put failed!")
						return
					}
				}
				count++
			}
			if count != 0 {
				rt.TaskId = tm.TaskId
				rt.VSAssets = vs
				ok, _ := global.Queue.Put(rt)
				if !ok {
					zap.S().Info("put failed!")
					return
				}
			}
		} else if tm.CompanyRange == "0" && tm.AssetRange == "4" { //全部公司   站点
			vs := task.VSAsset{}
			vs.PocContent = poc.PocContent
			count := 0 //500个目标做为一个任务

			//---------------------------
			website := []model.Websites{}
			if err = db.Orm.Model(&model.Websites{}).Select("website").Find(&website).Error; err != nil {
				zap.S().Errorf(err.Error())
				return
			}

			vs.PocContent = poc.PocContent
			count = 0 //500个目标做为一个任务
			for _, v := range website {
				vs.Targets = append(vs.Targets, v.Website)
				if count > 500 {
					rt.VSAssets = vs
					rt.TaskId = tm.TaskId
					ok, _ := global.Queue.Put(rt)
					if !ok {
						zap.S().Info("put failed!")
						return
					}
				}
				count++
			}
			if count != 0 {
				rt.TaskId = tm.TaskId
				rt.VSAssets = vs
				ok, _ := global.Queue.Put(rt)
				if !ok {
					zap.S().Info("put failed!")
					return
				}
			}
		} else if tm.CompanyRange == "1" && tm.AssetRange == "4" { //单或多个公司  扫描站点
			cids := strings.Split(tm.CompanyId, ",")
			cids = cids[:len(cids)-1]
			for _, v := range cids {
				website := []model.Websites{}
				if err = db.Orm.Model(&model.Websites{}).Select("website").Where("cid=?", v).Find(&website).Error; err != nil {
					zap.S().Errorf(err.Error())
					return
				}

				vs := task.VSAsset{}
				vs.PocContent = poc.PocContent
				count := 0 //500个目标做为一个任务
				for _, v := range website {
					vs.Targets = append(vs.Targets, v.Website)
					if count > 500 {
						rt.VSAssets = vs
						rt.TaskId = tm.TaskId
						ok, _ := global.Queue.Put(rt)
						if !ok {
							zap.S().Info("put failed!")
							return
						}
						count = 0
					}
					count++
				}
				if count != 0 {
					rt.TaskId = tm.TaskId
					rt.VSAssets = vs
					ok, _ := global.Queue.Put(rt)
					if !ok {
						zap.S().Info("put failed!")
						return
					}
				}

			}

		} else if tm.CompanyRange == "1" && tm.AssetRange == "2" { //单或多个公司  扫描IP
			cids := strings.Split(tm.CompanyId, ",")
			cids = cids[:len(cids)-1]
			for _, v := range cids {
				ips := []model.Ips{}
				if err = db.Orm.Model(&model.Websites{}).Select("ip").Where("cid=?", v).Find(&ips).Error; err != nil {
					zap.S().Errorf(err.Error())
					return
				}

				vs := task.VSAsset{}
				vs.PocContent = poc.PocContent
				count := 0 //500个目标做为一个任务
				for _, v := range ips {
					vs.Targets = append(vs.Targets, v.Ip)
					if count > 500 {
						rt.VSAssets = vs
						rt.TaskId = tm.TaskId
						ok, _ := global.Queue.Put(rt)
						if !ok {
							zap.S().Info("put failed!")
							return
						}
						count = 0
					}
					count++
				}
				if count != 0 {
					rt.TaskId = tm.TaskId
					rt.VSAssets = vs
					ok, _ := global.Queue.Put(rt)
					if !ok {
						zap.S().Info("put failed!")
						return
					}
				}

			}

		} else if tm.CompanyRange == "1" && tm.AssetRange == "0" { //单或多个公司  站点和IP
			cids := strings.Split(tm.CompanyId, ",")
			cids = cids[:len(cids)-1]
			for _, v := range cids {
				website := []model.Websites{}
				if err = db.Orm.Model(&model.Websites{}).Select("website").Where("cid=?", v).Find(&website).Error; err != nil {
					zap.S().Errorf(err.Error())
					return
				}

				vs := task.VSAsset{}
				vs.PocContent = poc.PocContent
				count := 0 //500个目标做为一个任务
				for _, v := range website {
					vs.Targets = append(vs.Targets, v.Website)
					if count > 500 {
						rt.VSAssets = vs
						rt.TaskId = tm.TaskId
						ok, _ := global.Queue.Put(rt)
						if !ok {
							zap.S().Info("put failed!")
							return
						}
						count = 0
					}
					count++
				}
				if count != 0 {
					rt.TaskId = tm.TaskId
					rt.VSAssets = vs
					ok, _ := global.Queue.Put(rt)
					if !ok {
						zap.S().Info("put failed!")
						return
					}
				}

				ips := []model.Ips{}
				if err = db.Orm.Model(&model.Websites{}).Select("ip").Where("cid=?", v).Find(&ips).Error; err != nil {
					zap.S().Errorf(err.Error())
					return
				}

				count = 0 //500个目标做为一个任务
				for _, v := range ips {
					vs.Targets = append(vs.Targets, v.Ip)
					if count > 500 {
						rt.VSAssets = vs
						rt.TaskId = tm.TaskId
						ok, _ := global.Queue.Put(rt)
						if !ok {
							zap.S().Info("put failed!")
							return
						}
						count = 0
					}
					count++
				}
				if count != 0 {
					rt.TaskId = tm.TaskId
					rt.VSAssets = vs
					ok, _ := global.Queue.Put(rt)
					if !ok {
						zap.S().Info("put failed!")
						return
					}
				}
			}
		}
	}

	buf, err := json.Marshal(rt)
	if err != nil {
		zap.S().Errorf(err.Error())
		return
	}
	//加密数据
	data, err := utils.EncryptByAes(buf)
	if err != nil {
		zap.S().Errorf(err.Error())
		return
	}
	ok, _ := global.Queue.Put(data)
	if !ok {
		zap.S().Errorf("put failed")
		return
	}
}

func StopTask(taskId string) error {
	//根据任务名称停止任务
	var err error
	tk := model.Task{}
	if err = db.Orm.Model(&model.Task{}).Where("taskId=?", taskId).Find(&tk).Error; err != nil {
		return err
	}
	*tk.TaskStatus = !*tk.TaskStatus
	err = db.Orm.Model(&model.Task{}).Where("taskName=?", tk.TaskName).Updates(&tk).Error
	if err != nil {
		return err
	}

	if _, ok := global.TaskCore[taskId]; ok {
		cancel := global.TaskCore[taskId].(context.CancelFunc)
		cancel()
		zap.S().Info("stop task id:%s", taskId)
		return nil
	} else {
		zap.S().Info("task not exist")
		return errors.New("task not exist")
	}
}

func TaskInfo(taskId string) (*model.Task, error) {

	var task = &model.Task{}
	err := db.Orm.Where("taskId=?", taskId).Find(&task).Error
	if err != nil {
		return nil, err
	}

	return task, nil

}

func DelTaskByTaskId(taskId string) error {

	var task model.Task
	err := db.Orm.Where("taskId=?", taskId).Delete(&task).Error
	if err != nil {
		return err
	}
	//删除当前正在运行的任务
	if _, ok := global.TaskCore[taskId]; ok {
		cancel := global.TaskCore[taskId].(context.CancelFunc)
		cancel()
		zap.S().Info("stop task id: %s", taskId)
		return nil
	} else {
		zap.S().Info("task:[" + taskId + "] not exist")
		return errors.New("task:[" + taskId + "]  not exist")
	}

}

func RunTask(taskId string) error {
	//根据任务名称停止任务
	var err error
	tk := model.Task{}
	if err = db.Orm.Model(&model.Task{}).Where("taskId=?", taskId).Find(&tk).Error; err != nil {
		return err
	}
	*tk.TaskStatus = !*tk.TaskStatus
	err = db.Orm.Model(&model.Task{}).Where("taskId=?", taskId).Updates(&tk).Error
	if err != nil {
		return err
	}

	if *tk.TaskStatus {
		tastContext, cancel := context.WithCancel(context.Background())
		global.TaskCore[tk.TaskId] = cancel

		go func(ctx context.Context, tm *model.Task) {
			zap.S().Info("task id:%s start.", tm.TaskId)
			tick := time.NewTicker(time.Duration(tm.Rate) * time.Minute)
			for {
				select {
				case <-tick.C:
					newtask(tm)
				case <-ctx.Done():
					zap.S().Info("task id:%s shutdown.", tm.TaskId)
					return
				}
			}
		}(tastContext, &tk)
	}
	return nil
}

func GetTaskList(query *request.Query) ([]model.Task, int, error) {
	tk := []model.Task{}
	err := db.Orm.Model(&model.Task{}).Limit(query.PageSize).Offset((query.Page - 1) * query.PageSize).Order("updateTime desc").Find(&tk).Error
	if err != nil {
		zap.S().Errorf("%s", err.Error())
		return nil, 0, nil
	}

	for i := 0; i < len(tk); i++ {
		cids := strings.Split(tk[i].CompanyId, ",")
		var cName string
		for _, v := range cids {
			comp := model.Company{}
			err := db.Orm.Model(&model.Company{}).Where("id=?", v).Find(&comp).Error
			if err != nil {
				zap.S().Errorf("%s", err.Error())
				return nil, 0, nil
			}
			cName += comp.CompanyName + " "
		}
		tk[i].CompanyId = cName
	}
	var count int64
	err = db.Orm.Model(&model.Task{}).Count(&count).Error
	if err != nil {
		zap.S().Errorf("%s", err.Error())
		return nil, 0, nil
	}

	return tk, 0, nil
}

func OnceInitTask() {
	zap.S().Info("init task")
	tk := []model.Task{}
	err := db.Orm.Model(&model.Task{}).Find(&tk).Error
	if err != nil {
		zap.S().Errorf("%s", err.Error())
		return
	}

	for _, v := range tk {
		if *v.TaskStatus {
			tastContext, cancel := context.WithCancel(context.Background())
			global.TaskCore[v.TaskId] = cancel

			go func(ctx context.Context, tm model.Task) {
				zap.S().Info("task id: " + tm.TaskId + " start.")
				tick := time.NewTicker(time.Duration(tm.Rate) * time.Minute)
				for {
					select {
					case <-tick.C:
						newtask(&tm)
					case <-ctx.Done():
						zap.S().Info("task id:%s shutdown.", tm.TaskId)
						return
					}
				}
			}(tastContext, v)
		}
	}
}

func UploadDomains(dsr request.DomainScanRes) error {

	domains := strings.Split(dsr.Domains, ",")

	for _, v := range domains {
		var count int64
		err := db.Orm.Model(&model.Domain{}).Where("domain=?", v).Count(&count).Error
		if err != nil {
			zap.S().Errorf("%s", err.Error())
			return err
		}

		dom := model.Domain{}
		dom.Cid = dsr.Cid
		dom.Domain = v
		dom.Source = "admin"
		dom.UpdateTime = time.Now().Format("2006-01-02 15:04:05")
		*dom.IsNew = true

		if count <= 0 {
			err = db.Orm.Model(&model.Domain{}).Create(&dom).Error
			if err != nil {
				zap.S().Errorf("%s", err.Error())
				return err
			}
		}

	}
	return nil
}
