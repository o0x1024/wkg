package domainService

import (
	"bytes"
	"context"
	"crypto/tls"
	"fmt"
	"io"
	"log"
	"os"
	"wkg/global"
	"wkg/model"
	"wkg/model/request"
	db2 "wkg/pkg/db"
	helper2 "wkg/pkg/libs/helper"

	// "wkg/services"

	"go.uber.org/zap"
	"gopkg.in/yaml.v3"

	"io/ioutil"
	"math/rand"
	"net"
	"net/http"
	"regexp"
	"strings"
	"sync"
	"time"

	// "wkg/module/subdomainscan/amass/requests"

	"github.com/projectdiscovery/subfinder/v2/pkg/passive"
	"github.com/projectdiscovery/subfinder/v2/pkg/resolve"
	"github.com/projectdiscovery/subfinder/v2/pkg/runner"
)

type DomainMoniter struct {
	DomainCache           []model.Domain
	UpdateDomainCacheTime time.Duration
	threadNum             int
	cmpInfo               []model.Company
	updateCmpInfoTime     time.Duration
	ScanDomainInfoTime    time.Duration
	ctx                   context.Context
	cancel                context.CancelFunc
	newCmpFlag            chan bool
	once                  *sync.Once
	lock                  *sync.RWMutex
	wg                    *sync.WaitGroup
}

func NewDomainMoniter() *DomainMoniter {
	ctx, cancel := context.WithCancel(context.Background())
	return &DomainMoniter{
		threadNum:          1,
		wg:                 &sync.WaitGroup{},
		newCmpFlag:         make(chan bool),
		ctx:                ctx,
		once:               &sync.Once{},
		cancel:             cancel,
		lock:               &sync.RWMutex{},
		updateCmpInfoTime:  domainMonitorTime,
		ScanDomainInfoTime: domainMonitorTime,
	}
}

var domainMonitorTime = time.Duration(24) * time.Hour

func StartDomainMonitorService() {
	rand.Seed(time.Now().UTC().UnixNano())
	dm := NewDomainMoniter()
	dm.startDomainMonitorService()

}

func ScanDomain(cmp model.Company) {
	// sd := strings.Split(cmp.Domain, "|")
	// for _, vd := range sd {
	// 	var domainResult []requests.Output

	// 	log.Println("[*] domain scan by Subfinder:", vd)
	// 	subfindResultList := ScanDomainBySubfinder(vd)

	// 	if subfindResultList != nil {
	// 		var rq = requests.Output{}
	// 		for _, v := range subfindResultList {
	// 			rq.Domain = v
	// 			rq.Sources = []string{"subfinder"}
	// 			domainResult = append(domainResult, rq)
	// 			rq = requests.Output{}
	// 		}
	// 	}

	// 	if domainResult != nil {
	// 		for _, v := range domainResult {
	// 			//如果域名长度大于150，丢弃这个域名。
	// 			if len(v.Name) > 150 {
	// 				continue
	// 			}
	// 			//判断是否已经存在，如果已经存在就不储存
	// 			var count int64
	// 			err := db2.Orm.Model(&model.Domain{}).Where("domain=?", v.Domain).Count(&count).Error
	// 			if err != nil {
	// 				continue
	// 			}
	// 			if count <= 0 {
	// 				//如果记录不存在就插到数据库里
	// 				fp := true
	// 				var newDomain = model.Domain{IsNew: &fp}
	// 				newDomain.Domain = v.Domain
	// 				newDomain.Cid = cmp.Id
	// 				*newDomain.IsNew = true
	// 				if len(v.Sources) > 3 {
	// 					newDomain.Source = v.Sources[0]

	// 				} else {
	// 					newDomain.Source = fmt.Sprintf("%v", v.Sources)

	// 				}
	// 				newDomain.UpdateTime = helper2.GetCurTime()
	// 				err := db2.Orm.Model(&model.Domain{}).Create(&newDomain).Error
	// 				if err != nil {
	// 					zap.S().Errorf("%s", err.Error())
	// 					return
	// 				}
	// 			}
	// 		}
	// 	}
	// }
}

// 每5秒查看一次是否有新公司
func (d *DomainMoniter) monitorNewCompany() {
	ticker := time.NewTicker(5 * time.Second)
	for {
		select {
		case <-ticker.C:
			if global.HasNewCompanyFlag {
				d.newCmpFlag <- true
				global.HasNewCompanyFlag = false
			}
		case <-d.ctx.Done():
			return

		}
	}
}

func (d *DomainMoniter) startDomainMonitorService() {
	log.Println("[*] Start domain monitor service.")
	//按给定时间读取数据库中的公司信息

	//监控是否有新公司被添加
	go d.monitorNewCompany()
	go d.loopUpdateCmpInfo()
	//域名监控功能
	//按设定的总时间一次性从数据库里读取信息，进行域名监控扫描。
	go d.loopMonitorDomain()
	//1.每个公司开一个协程，根据monitorstatus来确定是否需要监控

}

func (d *DomainMoniter) loopMonitorDomain() {
	//ticker := time.NewTicker(tim*time.Hour)
	ticker := time.NewTicker(d.ScanDomainInfoTime)
	for {
		select {
		case <-ticker.C:
			d.scanDomainByAmass()
		case <-d.newCmpFlag:
			d.reloadCmpInfo()
			d.onceCcanDomainByAmass()
		case <-d.ctx.Done():
			return
		}
	}
}

func (d *DomainMoniter) onceCcanDomainByAmass() {
	//Glog.Println("[*] start domain monitor queue.")
	if len(d.cmpInfo) <= 0 {
		return
	}
	//判断监控状态
	if !*d.cmpInfo[len(d.cmpInfo)-1].MonitorStatus {
		return
	}
	sd := strings.Split(d.cmpInfo[len(d.cmpInfo)-1].Domain, "|")
	err := db2.Orm.Model(&model.Domain{}).Where("cid=?", d.cmpInfo[len(d.cmpInfo)-1].Id).Find(&d.DomainCache).Error
	if err != nil {
		zap.S().Errorf("%s", err.Error())
		return
	}
	for _, vd := range sd {
		DRes := ScanDomainBySubfinder(vd)
		fp := true
		if DRes != nil {
			for _, v := range DRes {
				//如果域名长度大于150，丢弃这个域名。
				if len(v) > 150 {
					continue
				}
				//判断是否已经存在，如果已经存在就不储存
				var isExsit = false
				for _, s := range d.DomainCache {
					if v == s.Domain {
						isExsit = true
						break
					}
				}
				if !isExsit {
					//如果记录不存在就插到数据库里
					var newDomain = model.Domain{IsNew: &fp}
					newDomain.Domain = v
					newDomain.Cid = d.cmpInfo[len(d.cmpInfo)-1].Id
					*newDomain.IsNew = true

					newDomain.Source = fmt.Sprintf("%v", "subfinder")

					newDomain.UpdateTime = helper2.GetCurTime()
					d.DomainCache = append(d.DomainCache, newDomain)
					err := db2.Orm.Model(&model.Domain{}).Create(&newDomain).Error
					if err != nil {
						zap.S().Errorf("%s", err.Error())
						return
					}
					isExsit = false
				}
			}
		}
	}

}

func ScanDomainBySubfinder(domain string) []string {

	//读取subfinder.yaml
	providers := &runner.Providers{}

	fp, err := os.Open("conf/subfinder.yaml")
	if err != nil {
		zap.S().Errorf("can not open conf/subfinder.yaml file.")
		return nil
	}
	defer fp.Close()
	ymlbuf, err := ioutil.ReadAll(fp)
	if err != nil {
		zap.S().Errorf(err.Error())
		return nil
	}
	err = yaml.Unmarshal(ymlbuf, providers)
	if err != nil {
		zap.S().Errorf(err.Error())
		return nil
	}

	runnerInstance, err := runner.NewRunner(&runner.Options{
		Threads:            10,                              // Thread controls the number of threads to use for active enumerations
		Timeout:            30,                              // Timeout is the seconds to wait for sources to respond
		MaxEnumerationTime: 10,                              // MaxEnumerationTime is the maximum amount of time in mins to wait for enumeration
		Resolvers:          resolve.DefaultResolvers,        // Use the default list of resolvers by marshaling it to the config
		Sources:            passive.DefaultSources,          // Use the default list of passive sources
		AllSources:         passive.DefaultAllSources,       // Use the default list of all passive sources
		Recursive:          passive.DefaultRecursiveSources, // Use the default list of recursive sources
		Providers:          providers,                       // Use empty api keys for all providers
	})
	if err != nil {
		zap.S().Errorf(err.Error())
		return nil
	}

	buf := bytes.Buffer{}
	err = runnerInstance.EnumerateSingleDomain(context.Background(), domain, []io.Writer{&buf})
	if err != nil {
		zap.S().Errorf(err.Error())
		return nil
	}

	data, err := ioutil.ReadAll(&buf)
	if err != nil || data == nil {
		zap.S().Errorf(err.Error())
		return nil
	}
	domainList := strings.Split(string(data), "\n")
	return domainList

}

func (d *DomainMoniter) scanDomainByAmass() {
	// //Glog.Println("[*] start domain monitor queue.")
	// if len(d.cmpInfo) <= 0 {
	// 	zap.S().Info("company count 0")
	// 	return
	// }
	// for i := len(d.cmpInfo) - 1; i > 0; i-- {
	// 	//判断监控状态
	// 	if !*d.cmpInfo[i].MonitorStatus {
	// 		continue
	// 	}
	// 	sd := strings.Split(d.cmpInfo[i].Domain, "|")
	// 	err := db2.Orm.Model(&model.Domain{}).Where("cid=?", d.cmpInfo[i].Id).Find(&d.DomainCache).Error
	// 	if err != nil {
	// 		zap.S().Errorf(err.Error())
	// 	}
	// 	for _, vd := range sd {
	// 		log.Println("[*] domain monitor:", vd)

	// 		var domainResult []requests.Output
	// 		var errCh = make(chan error, 1)
	// 		safe2.Go(func() error {
	// 			domainResult = subdomainscan2.DomainBrute(vd)
	// 			return nil
	// 		}, errCh)
	// 		if <-errCh != nil {
	// 			zap.S().Errorf("%s", <-errCh)
	// 			continue
	// 		}
	// 		log.Println("[+] found new domain count:", len(domainResult))
	// 		fp := true
	// 		if domainResult != nil {
	// 			for _, v := range domainResult {
	// 				//如果域名长度大于150，丢弃这个域名。
	// 				if len(v.Name) > 150 {
	// 					continue
	// 				}
	// 				//判断是否已经存在，如果已经存在就不储存
	// 				var isExsit = false
	// 				for _, s := range d.DomainCache {
	// 					if v.Name == s.Domain {
	// 						isExsit = true
	// 						break
	// 					}
	// 				}
	// 				if !isExsit {
	// 					//如果记录不存在就插到数据库里
	// 					var newDomain = model.Domain{IsNew: &fp}
	// 					newDomain.Domain = v.Name
	// 					newDomain.Cid = d.cmpInfo[i].Id
	// 					*newDomain.IsNew = true
	// 					if len(v.Sources) > 3 {
	// 						newDomain.Source = v.Sources[0]

	// 					} else {
	// 						newDomain.Source = fmt.Sprintf("%v", v.Sources)

	// 					}
	// 					newDomain.UpdateTime = helper2.GetCurTime()
	// 					d.DomainCache = append(d.DomainCache, newDomain)
	// 					err := db2.Orm.Model(&model.Domain{}).Create(&newDomain).Error
	// 					if err != nil {
	// 						zap.S().Errorf("%s", err.Error())
	// 						continue
	// 					}
	// 					isExsit = false
	// 				}
	// 			}
	// 		}
	// 	}
	// 	zap.S().Info("domain cache:%d", len(d.DomainCache))
	// }
}

func (d *DomainMoniter) reloadCmpInfo() {
	err := db2.Orm.Model(&model.Company{}).Find(&d.cmpInfo).Error
	if err != nil {
		log.Println("[!] domainService.go  line:137 db error:", err)
	}
}

func GetDomainInfo(query *request.Query) ([]model.Domain, int64, error) {
	var err error
	var count int64
	//查询总数

	//根据page和pagesize查询数据
	var dom []model.Domain
	if query.AssetOption == "0" {
		err = db2.Orm.Model(&model.Domain{}).Count(&count).Limit(query.PageSize).Offset((query.Page - 1) * query.PageSize).Find(&dom).Error
		if err != nil {
			return nil, 0, err
		}
	} else {
		err = db2.Orm.Model(&model.Domain{}).Limit(query.PageSize).Offset((query.Page - 1) * query.PageSize).Where("isNew=true").Find(&dom).Error
		if err != nil {
			return nil, 0, err
		}
		err = db2.Orm.Model(&model.Domain{}).Where("isNew=true").Count(&count).Error
		if err != nil {
			return nil, 0, err
		}
	}

	return dom, count, nil
}

func GetNewDomainInfo(query *request.Query) ([]model.Domain, int64, error) {
	var err error
	var count int64
	//查询总数
	err = db2.Orm.Model(&model.Domain{}).Where("isNew=true").Count(&count).Error
	if err != nil {
		return nil, 0, err
	}
	//根据page和pagesize查询数据
	var dom []model.Domain
	err = db2.Orm.Model(&model.Domain{}).Where("isNew=true").Limit(query.PageSize).Offset((query.Page - 1) * query.PageSize).Find(&dom).Error
	if err != nil {
		return nil, 0, err
	}

	return dom, count, nil
}

func ReadFlagDomainInfoById(ids *request.Ids) error {

	_ids := strings.Split(ids.Id, ",")
	for i, v := range _ids {
		if v == "" && i == len(_ids)-1 {
			break
		}
		err := db2.Orm.Model(&model.Domain{}).Where("id=?", v).Update("isNew", false).Error
		if err != nil {
			return err
		}
	}
	return nil
}

func GetDomainInfoByCid(query *request.Query) ([]model.Domain, int64, error) {
	var err error
	var count int64
	//查询总数
	// err = db2.Orm.Model(&model.Domain{}).Where("cid=?", query.Cid).Count(&count).Error
	// if err != nil {
	// 	return nil, 0, err
	// }
	//根据page和pagesize查询数据
	var dom []model.Domain

	if query.AssetOption == "0" {
		if query.Keyword != "" {
			err = db2.Orm.Model(&model.Domain{}).Limit(query.PageSize).Offset((query.Page-1)*query.PageSize).Where("cid=? and domain=?", query.Cid, query.Keyword).Find(&dom).Error
			if err != nil {
				return nil, 0, err
			}
			if err = db2.Orm.Model(&model.Domain{}).Where("cid=? and isNew=true", query.Cid).Count(&count).Error; err != nil {
				return nil, 0, err
			}
		} else {
			err = db2.Orm.Model(&model.Domain{}).Limit(query.PageSize).Offset((query.Page-1)*query.PageSize).Where("cid=?", query.Cid).Find(&dom).Error
			if err != nil {
				return nil, 0, err
			}
			if err = db2.Orm.Model(&model.Domain{}).Where("cid=? and isNew=true", query.Cid).Count(&count).Error; err != nil {
				return nil, 0, err
			}
		}
	} else {
		if query.Keyword != "" {
			err = db2.Orm.Model(&model.Domain{}).Limit(query.PageSize).Offset((query.Page-1)*query.PageSize).Where("cid=? and domain=? and isNew=true", query.Cid, query.Keyword).Find(&dom).Error
			if err != nil {
				return nil, 0, err
			}
			if err = db2.Orm.Model(&model.Domain{}).Where("cid=? and isNew=true", query.Cid).Count(&count).Error; err != nil {
				return nil, 0, err
			}
		} else {
			err = db2.Orm.Model(&model.Domain{}).Limit(query.PageSize).Offset((query.Page-1)*query.PageSize).Where("cid=? and isNew=true", query.Cid).Find(&dom).Error

			if err != nil {
				return nil, 0, err
			}
			if err = db2.Orm.Model(&model.Domain{}).Where("cid=? and isNew=true", query.Cid).Count(&count).Error; err != nil {
				return nil, 0, err
			}
		}
	}

	return dom, count, nil
}

func GetDomainInfoByKey(query *request.Query) ([]model.Domain, int64, error) {

	var err error

	var count int64
	//查询总数
	err = db2.Orm.Model(&model.Domain{}).Where("cid=?", query.Cid).Count(&count).Error
	if err != nil {
		return nil, 0, err
	}

	var dom = []model.Domain{}
	if query.Type == "ip" {
		err = db2.Orm.Model(&model.Domain{}).Where("ip LIKE ?", "%"+query.Keyword+"%").Find(&dom).Error
	} else if query.Type == "domain" {
		err = db2.Orm.Model(&model.Domain{}).Where("domain LIKE ?", "%"+query.Keyword+"%").Find(&dom).Error
	} else if query.Type == "title" {
		err = db2.Orm.Model(&model.Domain{}).Where("title LIKE ?", "%"+query.Keyword+"%").Find(&dom).Error
	}
	if err != nil {
		return nil, 0, err
	}

	return dom, count, err
}

func (d *DomainMoniter) loopUpdateCmpInfo() {

	err := db2.Orm.Model(&model.Company{}).Find(&d.cmpInfo).Error
	if err != nil {
		zap.S().Errorf("%s", err.Error())
		return
	}

	ticker := time.NewTicker(d.updateCmpInfoTime)
	for {
		select {
		case <-ticker.C:
			//获取各公司的域名信息和待监控状态
			err := db2.Orm.Model(&model.Company{}).Find(&d.cmpInfo).Error
			if err != nil {
				zap.S().Errorf("%s", err.Error())
				return
			}

		}
	}

}

func (d *DomainMoniter) getIPAndTitle(domain string) (ip string, title string) {

	//先获取IP地址
	ips, err := net.ResolveIPAddr("ip", domain)
	if err != nil {
		//Glog.Println("Resolution error", err.Error())
		return
	}

	ip = ips.String()
	if len(ips.String()) <= 0 {
		return
	}

	domain = strings.ReplaceAll(domain, "\r", "")
	domain = strings.ReplaceAll(domain, "\n", "")
	v1 := "http://" + domain
	body, err := httpGet(v1)
	if err != nil {
		zap.S().Errorf("%s", err.Error())
		return
	}
	exp := regexp.MustCompile(`<title>(.*?)</title>`)
	result := exp.FindAllStringSubmatch(string(body), -1)
	for _, text := range result {
		title = text[1]
	}
	if len(title) >= 0 {
		return
	}

	//https请求
	v1 = "https://" + domain
	body, err = httpGet(v1)
	if err != nil {
		zap.S().Errorf("%s", err.Error())
		return
	}
	exp = regexp.MustCompile(`<title>(.*?)</title>`)
	result = exp.FindAllStringSubmatch(string(body), -1)
	for _, text := range result {
		title = text[1]
	}

	return
}

func httpGet(url string) (result string, err error) {
	tr := &http.Transport{
		//Proxy:               p,
		TLSClientConfig:       &tls.Config{InsecureSkipVerify: true},
		TLSHandshakeTimeout:   8 * time.Second,
		ResponseHeaderTimeout: 4 * time.Second,
		DisableKeepAlives:     false,
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		zap.S().Errorf("%s", err.Error())
		return
	}
	client := http.Client{Transport: tr, Timeout: 2 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func DelDomainById(id string) error {
	var domain model.Domain
	err := db2.Orm.Where("id=?", id).Delete(&domain).Error
	if err != nil {
		return err
	}
	return nil
}
