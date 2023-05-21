package websiteService

import (
	"context"
	"crypto/tls"
	"io/ioutil"
	"net"
	"net/http"
	"regexp"
	"strings"
	"sync"
	"time"
	"wkg/model"
	"wkg/model/request"
	db2 "wkg/pkg/db"
	helper2 "wkg/pkg/libs/helper"

	"github.com/pkg/errors"
	"go.uber.org/zap"
	"gorm.io/gorm/clause"
)

//https://user.95516.com/favicon.ico
//检测域名库和IP库isNew字段是否为true，为true就拿过来测试站点是否能访问，能访问就记录相关数据

type mq struct {
	domain string
	cid    int
}

var tr = &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}
var client = http.Client{Timeout: 3 * time.Second, Transport: tr}

type WebSiteService struct {
	wsCache      sync.Map
	ctx          context.Context
	cancel       context.CancelFunc
	monitorTime  time.Duration
	monitorQueue []mq
	threadnum    int
	tch          chan mq
	client       http.Client
}

func GetWebSiteInfo(query *request.Query) (website []model.Websites, count int64, err error) {

	if query.AssetOption == "0" {
		if query.Keyword == "" {
			err = db2.Orm.Model(&model.Websites{}).Count(&count).Limit(query.PageSize).Offset((query.Page - 1) * query.PageSize).Find(&website).Error
		} else if query.Type == "title" {
			err = db2.Orm.Model(&model.Websites{}).Where(query.Type+" LIKE ?", "%"+query.Keyword+"%").Count(&count).Limit(query.PageSize).Offset((query.Page - 1) * query.PageSize).Find(&website).Error
		} else if query.Type == "website" {
			err = db2.Orm.Model(&model.Websites{}).Where("website LIKE ?", "%"+query.Keyword+"%").Count(&count).Limit(query.PageSize).Offset((query.Page - 1) * query.PageSize).Find(&website).Error
		} else if query.Type == "favicon" {
			err = db2.Orm.Model(&model.Websites{}).Where("favicon LIKE ?", "%"+query.Keyword+"%").Count(&count).Limit(query.PageSize).Offset((query.Page - 1) * query.PageSize).Find(&website).Error
		} else if query.Type == "ips" {
			err = db2.Orm.Model(&model.Websites{}).Where("ips LIKE ?", "%"+query.Keyword+"%").Count(&count).Limit(query.PageSize).Offset((query.Page - 1) * query.PageSize).Find(&website).Error
		}
		if err != nil {
			return
		}
	} else {
		if query.Keyword == "" {
			err = db2.Orm.Model(&model.Websites{}).Where("isNew=true").Count(&count).Limit(query.PageSize).Offset((query.Page - 1) * query.PageSize).Find(&website).Error
		} else if query.Type == "title" {
			err = db2.Orm.Model(&model.Websites{}).Where(query.Type+" LIKE ? and isNew=true", "%"+query.Keyword+"%").Count(&count).Limit(query.PageSize).Offset((query.Page - 1) * query.PageSize).Find(&website).Error
		} else if query.Type == "website" {
			err = db2.Orm.Model(&model.Websites{}).Where("website LIKE ? and isNew=true", "%"+query.Keyword+"%").Count(&count).Limit(query.PageSize).Offset((query.Page - 1) * query.PageSize).Find(&website).Error
		} else if query.Type == "favicon" {
			err = db2.Orm.Model(&model.Websites{}).Where("favicon LIKE ? and isNew=true", "%"+query.Keyword+"%").Count(&count).Limit(query.PageSize).Offset((query.Page - 1) * query.PageSize).Find(&website).Error
		} else if query.Type == "ips" {
			err = db2.Orm.Model(&model.Websites{}).Where("ips LIKE ? and isNew=true", "%"+query.Keyword+"%").Count(&count).Limit(query.PageSize).Offset((query.Page - 1) * query.PageSize).Find(&website).Error
		}
		if err != nil {
			return
		}
	}
	return
}

func GetWebSiteInfoByKey(dss *request.SearchStrut) (website []model.Websites, count int64, err error) {
	if dss.Type == "title" {
		err = db2.Orm.Model(&model.Websites{}).Where("title LIKE ?", "%"+dss.KeyWord+"%").Count(&count).Find(&website).Error
	} else if dss.Type == "website" {
		err = db2.Orm.Model(&model.Websites{}).Where("domain LIKE ?", "%"+dss.KeyWord+"%").Count(&count).Find(&website).Error
	} else if dss.Type == "favicon" {
		err = db2.Orm.Model(&model.Websites{}).Where("favicon LIKE ?", "%"+dss.KeyWord+"%").Count(&count).Find(&website).Error
	} else if dss.Type == "cid" {
		err = db2.Orm.Model(&model.Websites{}).Where("cid LIKE ?", "%"+dss.KeyWord+"%").Count(&count).Find(&website).Error
	}

	if err != nil {
		return nil, 0, err
	}
	return
}

func GetWebSiteInfoByCId(param *request.Query) (website []model.Websites, count int64, err error) {

	if param.AssetOption == "0" {
		if param.Keyword == "" {
			err = db2.Orm.Model(&model.Websites{}).Where("cid=?", param.Cid).Count(&count).Limit(param.PageSize).Offset((param.Page - 1) * param.PageSize).Find(&website).Error
		} else if param.Type == "title" {
			err = db2.Orm.Model(&model.Websites{}).Where("title LIKE ? and cid=?", "%"+param.Keyword+"%", param.Cid).Count(&count).Limit(param.PageSize).Offset((param.Page - 1) * param.PageSize).Find(&website).Error
		} else if param.Type == "website" {
			err = db2.Orm.Model(&model.Websites{}).Where("website LIKE ? and cid=?", "%"+param.Keyword+"%", param.Cid).Count(&count).Limit(param.PageSize).Offset((param.Page - 1) * param.PageSize).Find(&website).Error
		} else if param.Type == "favicon" {
			err = db2.Orm.Model(&model.Websites{}).Where("favicon LIKE ? and cid=?", "%"+param.Keyword+"%", param.Cid).Count(&count).Limit(param.PageSize).Offset((param.Page - 1) * param.PageSize).Find(&website).Error
		} else if param.Type == "ips" {
			err = db2.Orm.Model(&model.Websites{}).Where("ips LIKE ? and cid=?", "%"+param.Keyword+"%", param.Cid).Count(&count).Limit(param.PageSize).Offset((param.Page - 1) * param.PageSize).Find(&website).Error
		}
	} else {
		if param.Keyword == "" {
			err = db2.Orm.Model(&model.Websites{}).Where("cid=? and isNew=true", param.Cid).Count(&count).Limit(param.PageSize).Offset((param.Page - 1) * param.PageSize).Find(&website).Error
		} else if param.Type == "title" {
			err = db2.Orm.Model(&model.Websites{}).Where("title LIKE ? and cid=? and isNew=true", "%"+param.Keyword+"%", param.Cid).Count(&count).Limit(param.PageSize).Offset((param.Page - 1) * param.PageSize).Find(&website).Error
		} else if param.Type == "website" {
			err = db2.Orm.Model(&model.Websites{}).Where("website LIKE ? and cid=? and isNew=true", "%"+param.Keyword+"%", param.Cid).Count(&count).Limit(param.PageSize).Offset((param.Page - 1) * param.PageSize).Find(&website).Error
		} else if param.Type == "favicon" {
			err = db2.Orm.Model(&model.Websites{}).Where("favicon LIKE ? and cid=? and isNew=true", "%"+param.Keyword+"%", param.Cid).Count(&count).Limit(param.PageSize).Offset((param.Page - 1) * param.PageSize).Find(&website).Error
		} else if param.Type == "ips" {
			err = db2.Orm.Model(&model.Websites{}).Where("ips LIKE ? and cid=? and isNew=true", "%"+param.Keyword+"%", param.Cid).Count(&count).Limit(param.PageSize).Offset((param.Page - 1) * param.PageSize).Find(&website).Error
		}
	}
	if err != nil {
		return nil, 0, err
	}

	return
}

func NewWebSiteService() *WebSiteService {
	ctx, cancel := context.WithCancel(context.Background())
	tr := &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}
	return &WebSiteService{
		ctx:         ctx,
		cancel:      cancel,
		monitorTime: wesiteMonitorTime,
		threadnum:   30,
		tch:         make(chan mq, 10),
		client:      http.Client{Timeout: 2 * time.Second, Transport: tr},
	}
}

var wesiteMonitorTime = time.Hour * 24

func (ws *WebSiteService) monitorWebSiteFromDomain() {

	var domains []model.Domain

	////每次重新读取数据库的时候清空一次
	//ws.monitorQueue=nil
	if len(ws.monitorQueue) <= 0 {
		err := db2.Orm.Model(&model.Domain{}).Select("domain", "cid").Where("isNew=true").Find(&domains).Error
		if err != nil {
			zap.S().Errorf("%s", err.Error())
			return
		}
		if len(domains) <= 0 {
			return
		}
		for _, v := range domains {
			var m mq
			m.domain = v.Domain
			m.cid = v.Cid
			ws.monitorQueue = append(ws.monitorQueue, m)

		}
	}

}

func (ws *WebSiteService) monitorWebSiteFromIP() {

}

func (ws *WebSiteService) worker() {
	tk := time.NewTicker(ws.monitorTime)

	//开启多个扫描groutine，太多数据库怕抗不住
	for i := 0; i < ws.threadnum; i++ {
		go ws.run()
	}

	for {
		select {
		case <-tk.C:
			for i := len(ws.monitorQueue) - 1; i > 0; i-- {
				ws.tch <- ws.monitorQueue[i]
			}
			ws.monitorQueue = nil
		}
	}
}

func ScanCompanyWebsiteInfo(cmp model.Company) error {

	//从域名扫描，只扫描80，443
	var domains []model.Domain
	err := db2.Orm.Model(&model.Domain{}).Select("domain,cid").Where("cid=?", cmp.Id).Find(&domains).Error
	if err != nil || len(domains) <= 0 {
		zap.S().Errorf("%s", err.Error())
		return errors.New("database error or not found new?")
	}
	go func() {
		for _, v := range domains {
			alivesite, err := getAliveSite(v.Domain)
			if err != nil {
				continue
			}
			for _, alv := range alivesite {
				//f:= true
				//判断是不是已经有了，如果有过
				var count int64 = 0
				if err := db2.Orm.Model(&model.Websites{}).Where("website=?", alv).Count(&count).Error; err != nil {
					zap.S().Errorf("%s", err.Error())
					return
				}
				if count != 0 {
					count = 0
					continue
				}

				website := model.Websites{IsNew: new(bool), CDN: new(bool)}
				title, headers, err := getTitleandHeader(alv)
				if err != nil {
					zap.S().Errorf("%s", err.Error())
					continue
				}
				website.Cid = v.Cid
				website.Website = alv
				website.Headers = headers
				var tmpstr string
				for _, v := range title {
					if ('\u4e00' < v && v < '\u9fa5') || ('\x00' < v && v < '\x7e') {
						tmpstr += string(v)
					}
				}
				website.Domain = v.Domain
				website.Title = tmpstr
				tmpstr = ""
				website.UpdateTime = helper2.GetCurTime()
				favhash, favhashUrl, err := getFaviconHash(alv)
				website.FaviconUrl = favhashUrl
				ips, err := getIP(v.Domain)
				if err == nil {
					if len(ips) > 1 {
						*website.CDN = true
					} else {
						website.Ips = ips[0]
						//添加IP到IP数据库
						var tip = model.Ips{IsNew: new(bool)}
						tip.Ip = ips[0]
						tip.Cid = v.Cid
						*tip.IsNew = true
						tip.UpdateTime = helper2.GetCurTime()
						err = db2.Orm.Model(&model.Ips{}).Create(&tip).Error
						if err != nil {
							zap.S().Errorf("%s", err.Error())
							return
						}
					}
				}
				*website.IsNew = true
				if err == nil {
					website.Favicon = favhash
				}
				err = db2.Orm.Model(&model.Websites{}).Create(&website).Error
				if err != nil {
					zap.S().Errorf("%s", err.Error())
					return
				}

			}
		}
	}()

	//从服务中扫描
	var svcList []model.Services
	if err = db2.Orm.Model(&model.Services{}).Select("cid,service,host,port").Where("cid=?", cmp.Id).Find(&svcList).Error; err != nil {
		zap.S().Errorf("database error or not found new?")
		return errors.New("database error or not found new?")
	}
	go func() {
		for _, v := range svcList {
			if strings.Contains(v.Service, "http") {
				website := model.Websites{IsNew: new(bool), CDN: new(bool)}
				alv := "http://" + v.Host + ":" + v.Port
				title, headers, errx := getTitleandHeader(alv)
				if errx != nil {
					continue
				}
				website.Cid = v.Cid
				website.Website = alv
				website.Headers = headers
				var tmpstr string
				for _, v := range title {
					if ('\u4e00' < v && v < '\u9fa5') || ('\x00' < v && v < '\x7e') {
						tmpstr += string(v)
					}
				}
				website.Domain = v.Host
				website.Title = tmpstr
				tmpstr = ""
				website.UpdateTime = helper2.GetCurTime()
				favhash, favhashUrl, err := getFaviconHash(alv)
				website.FaviconUrl = favhashUrl
				website.Ips = v.Host
				*website.IsNew = true
				if err == nil {
					website.Favicon = favhash
				}
				err = db2.Orm.Model(&model.Websites{}).Create(&website).Error
				if err != nil {
					zap.S().Errorf("%s", err.Error())
					return
				}
			}
		}
	}()

	return nil
}

func LoopScanNewWebSite() {
	tick := time.NewTicker(24 * time.Hour)
	for {
		select {
		case <-tick.C:
			err := ScanNewWebsiteInfo()
			if err != nil {
				zap.S().Errorf(err.Error())
				continue
			}
		}
	}
}

func ScanNewWebsiteInfo() error {

	var domains []model.Domain
	err := db2.Orm.Model(&model.Domain{}).Select("domain", "cid").Where("isNew=true").Find(&domains).Error
	if err != nil || len(domains) <= 0 {
		zap.S().Errorf("%s or database error or not found new?", err.Error())
		return errors.New("database error or not found new?")
	}
	go func() {
		for _, v := range domains {
			alivesite, err := getAliveSite(v.Domain)
			if err != nil {
				continue
			}
			for _, alv := range alivesite {
				//f:= true
				website := model.Websites{IsNew: new(bool), CDN: new(bool)}
				title, headers, errx := getTitleandHeader(alv)
				if errx != nil {
					zap.S().Errorf("%s", errx.Error())
					continue
				}
				website.Cid = v.Cid
				website.Website = alv
				website.Headers = headers
				var tmpstr string
				for _, v := range title {
					if ('\u4e00' < v && v < '\u9fa5') || ('\x00' < v && v < '\x7e') {
						tmpstr += string(v)
					}
				}
				website.Domain = v.Domain
				website.Title = tmpstr
				website.UpdateTime = helper2.GetCurTime()
				favhash, favhashUrl, err := getFaviconHash(alv)
				if err == nil {
					website.Favicon = favhash
				}
				website.FaviconUrl = favhashUrl
				ips, err := getIP(v.Domain)
				if err == nil {
					if len(ips) > 1 {
						*website.CDN = true
					} else {
						*website.CDN = false
						//添加IP到IP数据库
						{
							var tip = model.Ips{IsNew: new(bool)}
							tip.Ip = ips[0]
							tip.Cid = v.Cid
							*tip.IsNew = true
							tip.UpdateTime = helper2.GetCurTime()

							var count int64
							err = db2.Orm.Model(&model.Ips{}).Where("ip=?", tip.Ip).Count(&count).Error
							if err != nil {
								zap.S().Errorf("%s", err.Error())
								return
							}

							if count == 0 {
								err = db2.Orm.Model(&model.Ips{}).Clauses(clause.Insert{Modifier: "ignore"}).Create(&tip).Error
								if err != nil {
									zap.S().Errorf("%s", err.Error())
									return
								}
							}

						}
					}
				}
				*website.IsNew = true
				website.Ips = ips[0]

				var count int64 = 0
				err = db2.Orm.Model(&model.Websites{}).Clauses(clause.Insert{Modifier: "ignore"}).Where("website=?", website.Website).Count(&count).Error
				if err != nil {
					zap.S().Errorf("%s", err.Error())
					return
				}
				if count == 0 {
					err = db2.Orm.Model(&model.Websites{}).Create(&website).Error
					if err != nil {
						zap.S().Errorf("%s", err.Error())
						return
					}
				}
			}
		}
	}()

	return nil
}

func (ws *WebSiteService) run() {
	for v := range ws.tch {
		//查找当前缓存，当前站点是不是已经存在了
		_, ok := ws.wsCache.Load(v.domain)
		if ok {
			continue
		}
		//只要走到这一步就把域名添加到缓存里。
		ws.wsCache.Store(v.domain, v.domain)
		alivesite, err := getAliveSite(v.domain)
		if err != nil {
			continue
		}
		for _, alv := range alivesite {
			//f:= true
			website := model.Websites{IsNew: new(bool), CDN: new(bool)}
			title, headers, errx := getTitleandHeader(alv)
			if errx != nil {
				continue
			}
			website.Cid = v.cid
			website.Website = alv
			website.Headers = headers
			var tmpstr string
			for _, v := range title {
				if ('\u4e00' < v && v < '\u9fa5') || ('\x00' < v && v < '\x7e') {
					tmpstr += string(v)
				}
			}
			website.Domain = v.domain
			website.Title = tmpstr
			tmpstr = ""
			website.UpdateTime = helper2.GetCurTime()
			favhash, favhashUrl, err := getFaviconHash(alv)
			website.FaviconUrl = favhashUrl
			ips, err := getIP(v.domain)
			if err == nil {
				if len(ips) > 1 {
					*website.CDN = true
				} else {
					website.Ips = ips[0]
					//添加IP到IP数据库
					var tip = model.Ips{IsNew: new(bool)}
					tip.Ip = ips[0]
					tip.Cid = v.cid
					*tip.IsNew = true
					tip.UpdateTime = helper2.GetCurTime()
					err = db2.Orm.Model(&model.Ips{}).Create(&tip).Error
					if err != nil {
						zap.S().Errorf("%s", err.Error())
						return
					}
				}
			}
			*website.IsNew = true
			if err == nil {
				website.Favicon = favhash
			}
			err = db2.Orm.Model(&model.Websites{}).Create(&website).Error
			if err != nil {
				zap.S().Errorf("%s", err.Error())
				return
			}

		}
	}
}

func getAliveSite(target string) ([]string, error) {
	var testPort = []string{"80", "443"}
	var aliveUrl []string
	var url string
	for _, v := range testPort {
		if v == "443" {
			url = "https://" + target
		} else {
			url = "http://" + target + ":" + v
		}
		_, err := client.Get(url)
		if err != nil {
			continue
		}
		aliveUrl = append(aliveUrl, url)
		url = ""
	}
	if len(aliveUrl) == 0 {
		zap.S().Errorf("%s", "no alive")
		return nil, errors.New("no alive")
	}
	return aliveUrl, nil
}

func getIP(domain string) (ip []string, err error) {
	//先获取IP地址
	ips, err := net.LookupIP(domain)
	if err != nil {
		zap.S().Errorf("%s", err.Error())
		return
	}

	for _, v := range ips {
		ip = append(ip, v.String())
	}
	return
}

func getTitleandHeader(url string) (title string, header string, errs error) {

	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/95.0.4638.69 Safari/537.36 Edg/95.0.1020.53")
	if err != nil {
		zap.S().Errorf("%s", err.Error())
		return
	}
	resp, err := client.Do(req)
	if err != nil {
		zap.S().Errorf("%s", err.Error())
		return "", "", err
	}

	for k, v := range resp.Header {
		header += k
		header += ": "
		for _, vx := range v {
			header += vx

		}
		header += "\n"
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		zap.S().Errorf("%s", err.Error())
		return "", "", err
	}
	exp := regexp.MustCompile(`<title.*>(.*?)</title>`)
	result := exp.FindAllStringSubmatch(string(body), -1)
	for _, text := range result {
		title = text[1]
	}
	if len(title) >= 0 {
		return
	}

	return
}

func getFaviconHash(rooturl string) (string, string, error) {
	url := rooturl
	tr := &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}
	client := http.Client{Transport: tr, Timeout: 3 * time.Second}

	resp, err := client.Get(url)
	if err != nil {
		zap.S().Errorf("%s", err.Error())
		return "", "", err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		zap.S().Errorf("%s", err.Error())
		return "", "", err
	}

	var path string
	rex := regexp.MustCompile(`<link.*rel="{0,1}.*icon"{0,1}.*href="{0,1}(.*ico|png|icon|jpg)"{0,1}/{0,1}>`)
	resslice := rex.FindAllStringSubmatch(string(body), -1)
	for _, v := range resslice {
		if len(v[1]) > 0 {
			path = v[1]
		}
	}
	if len(path) <= 0 {
		rex = regexp.MustCompile("<link.*href=\"(.*)\".*rel=\".*icon\"")
		resslice = rex.FindAllStringSubmatch(string(body), -1)
		for _, v := range resslice {
			if len(v[1]) > 0 {
				path = v[1]
			}
		}
	}

	if len(path) <= 0 {
		fUrl := url + "/favicon.ico"
		resp, err = client.Get(fUrl)
		if err != nil {
			zap.S().Errorf("%s", err.Error())
			return "", "", err
		}
		if resp.StatusCode == 200 && strings.Contains(resp.Header.Get("Content-Type"), "image") {
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				zap.S().Errorf("%s", err.Error())
				return "", "", err
			}
			b64body := helper2.StandBase64(body)
			res := helper2.Mmh3Hash32([]byte(b64body))
			return res, path, nil
		}

	}

	if len(path) > 0 && !strings.Contains(path, "http") {
		fUrl := url + path
		resp, err = http.Get(fUrl)
		if err != nil {
			zap.S().Errorf("%s", err.Error())
			return "", "", err
		}
		body, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			zap.S().Errorf("%s", err.Error())
			return "", "", err
		}

		b64body := helper2.StandBase64(body)
		res := helper2.Mmh3Hash32([]byte(b64body))
		return res, fUrl, nil
	} else {
		resp, err = http.Get(path)
		if err != nil {
			zap.S().Errorf("%s", err.Error())
			return "", "", err
		}
		body, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			zap.S().Errorf("%s", err.Error())
			return "", "", err
		}

		b64body := helper2.StandBase64(body)
		res := helper2.Mmh3Hash32([]byte(b64body))
		return res, path, nil
	}

}

func ReadFlagWebSitefoById(cid *request.ParamIds) error {
	ids := strings.Split(cid.Id, ",")

	for i, v := range ids {
		if v == "" && i == len(ids)-1 {
			break
		}
		err := db2.Orm.Model(&model.Websites{}).Where("id=?", v).Update("isNew", false).Error
		if err != nil {

			return err
		}
	}
	return nil
}

func DelWebSiteById(id string) error {
	var website model.Websites
	err := db2.Orm.Where("id=?", id).Delete(&website).Error
	if err != nil {
		return err
	}
	return nil
}
