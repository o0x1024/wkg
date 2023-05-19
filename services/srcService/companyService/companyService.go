package companyService

import (
	"strconv"
	"time"
	"wkg/model"
	"wkg/model/request"
	"wkg/model/response"
	db2 "wkg/pkg/db"
)

func GetSelectOption() (option []response.Option, err error) {
	cmp := []model.Company{}
	err = db2.Orm.Model(&model.Company{}).Find(&cmp).Error
	if err != nil {

		return
	}

	for _, v := range cmp {
		option = append(option, response.Option{Value: strconv.Itoa(v.Id), Lable: v.CompanyName})
	}

	return
}

func UpdateCompanyInfo(tc *model.Company) (err error) {
	tc.LastUpdateTime = time.Unix(time.Now().Unix(), 0).Format("2006-01-02 15:04:05")

	//输入的时候就把回车换
	err = db2.Orm.Model(&model.Company{}).Where("id=?", tc.Id).Updates(&tc).Error
	if err != nil {
		return
	} else {
		return
	}
}

func GetCompanyById(cid string) (OneCmp *model.Company, err error) {
	err = db2.Orm.Model(&model.Company{}).Find(&OneCmp, "id=?", cid).Error
	if err != nil {
		return
	}
	return

}

func DelCompanyByCId(cid string) (err error) {
	var OneCmp model.Company
	err = db2.Orm.Delete(&OneCmp, cid).Error
	if err != nil {
		return
	}

	var tDom model.Domain
	err = db2.Orm.Where("cid=?", cid).Delete(&tDom).Error
	if err != nil {
		return
	}

	var tws model.Websites
	err = db2.Orm.Where("cid=?", cid).Delete(&tws).Error
	if err != nil {
		return
	}

	var svc model.Services
	err = db2.Orm.Where("cid=?", cid).Delete(&svc).Error
	if err != nil {
		return
	}

	var ips model.Ips
	err = db2.Orm.Where("cid=?", cid).Delete(&ips).Error
	if err != nil {
		return
	}

	var mini model.MiniProgram
	err = db2.Orm.Where("cid=?", cid).Delete(&mini).Error
	if err != nil {
		return
	}

	var wc model.WebChatOfficeAccount
	err = db2.Orm.Where("cid=?", cid).Delete(&wc).Error
	if err != nil {
		return
	}

	var notice model.Notice
	err = db2.Orm.Where("cid=?", cid).Delete(&notice).Error
	if err != nil {
		return
	}

	return
}

func GetCompanyInfo(param *request.Query) (dom []model.Company, total int64, err error) {
	//if param.Keyword == "" {
	//	err = db2.Orm.Model(&model.Company{}).Count(&count).Error
	//} else if param.Type == "companyName" {
	//	err = db2.Orm.Model(&model.Company{}).Where("companyName LIKE ?", "%"+param.Keyword+"%").Count(&count).Error
	//} else if param.Type == "domain" {
	//	err = db2.Orm.Model(&model.Company{}).Where("domain LIKE ?", "%"+param.Keyword+"%").Count(&count).Error
	//} else if param.Type == "keyWord" {
	//	err = db2.Orm.Model(&model.Company{}).Where("keyWord LIKE ?", "%"+param.Keyword+"%").Count(&count).Error
	//}
	//if err != nil {
	//
	//	return
	//}

	if param.Keyword == "" {
		err = db2.Orm.Model(&model.Company{}).Count(&total).Limit(param.PageSize).Offset((param.Page - 1) * param.PageSize).Find(&dom).Error
	} else if param.Type == "companyName" {
		err = db2.Orm.Model(&model.Company{}).Where("companyName LIKE ?", "%"+param.Keyword+"%").Count(&total).Limit(param.PageSize).Offset((param.Page - 1) * param.PageSize).Find(&dom).Error
	} else if param.Type == "domain" {
		err = db2.Orm.Model(&model.Company{}).Where("domain LIKE ?", "%"+param.Keyword+"%").Count(&total).Limit(param.PageSize).Offset((param.Page - 1) * param.PageSize).Find(&dom).Error
	} else if param.Type == "keyWord" {
	}
	if err != nil {
		return
	}
	return
}

func NewCompanyInfo(com *model.Company) (err error) {

	com.LastUpdateTime = time.Unix(time.Now().Unix(), 0).Format("2006-01-02 15:04:05")
	err = db2.Orm.Model(&model.Company{}).Create(&com).Error
	if err != nil {

		return
	}
	return

}

func GetCompanyByKey(dss *request.SearchStrut) ([]model.Company, int64, error) {
	var err error
	var total int64
	var dom = []model.Company{}
	if dss.Type == "companyName" {
		err = db2.Orm.Model(&model.Company{}).Where("companyName LIKE ?", "%"+dss.KeyWord+"%").Find(&dom).Count(&total).Error
	} else if dss.Type == "domain" {
		err = db2.Orm.Model(&model.Company{}).Where("domain LIKE ?", "%"+dss.KeyWord+"%").Count(&total).Find(&dom).Error
	} else if dss.Type == "keyWord" {
		err = db2.Orm.Model(&model.Company{}).Where("keyWord LIKE ?", "%"+dss.KeyWord+"%").Count(&total).Find(&dom).Error
	}
	if err != nil {
		return nil, 0, err
	}
	return dom, total, nil
}
