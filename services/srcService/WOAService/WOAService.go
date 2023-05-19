package WOAService

import (
	"strings"
	"wkg/model"
	"wkg/model/request"
	db2 "wkg/pkg/db"
)

func GetWOAInfo(query *request.Query) ([]model.WebChatOfficeAccount, int64, error) {
	var err error
	var WOA = []model.WebChatOfficeAccount{}
	var total int64

	if query.Keyword != "" {
		if err = db2.Orm.Model(&model.WebChatOfficeAccount{}).Where(query.Type+" like ?", "%"+query.Keyword+"%").Count(&total).Error; err != nil {
			return nil, 0, err
		}
		if err = db2.Orm.Model(&model.WebChatOfficeAccount{}).Limit(query.PageSize).Offset((query.Page-1)*query.PageSize).Where(query.Type+" like ?", "%"+query.Keyword+"%").Find(&WOA).Error; err != nil {
			return nil, 0, err
		}
	} else {
		if err = db2.Orm.Model(&model.WebChatOfficeAccount{}).Count(&total).Error; err != nil {
			return nil, 0, err
		}
		if err = db2.Orm.Model(&model.WebChatOfficeAccount{}).Limit(query.PageSize).Offset((query.Page - 1) * query.PageSize).Find(&WOA).Error; err != nil {
			return nil, 0, err
		}
	}

	return WOA, total, nil
}

func GetWOAInfoByCid(query *request.Query) ([]model.WebChatOfficeAccount, int64, error) {

	var svc []model.WebChatOfficeAccount
	err := db2.Orm.Model(&model.WebChatOfficeAccount{}).Where("cid=?", query.Cid).Limit(query.PageSize).Offset((query.Page - 1) * query.PageSize).Find(&svc).Error
	if err != nil {
		return nil, 0, err
	}

	var total int64
	//查询总数
	err = db2.Orm.Model(&model.WebChatOfficeAccount{}).Where("cid=?", query.Cid).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	return svc, total, nil

}

func DelWOAById(id string) error {
	var WOA model.WebChatOfficeAccount
	if err := db2.Orm.Where("id=?", id).Delete(&WOA).Error; err != nil {
		return err
	}
	return nil
}

func GetNewWOAInfo(query *request.Query) ([]model.WebChatOfficeAccount, int64, error) {
	var err error
	var total int64
	//查询总数
	err = db2.Orm.Model(&model.WebChatOfficeAccount{}).Where("isNew=true").Count(&total).Error
	if err != nil {
		return nil, 0, err
	}
	//根据page和pagesize查询数据
	var WOA []model.WebChatOfficeAccount
	err = db2.Orm.Model(&model.WebChatOfficeAccount{}).Where("isNew=true").Limit(query.PageSize).Offset((query.Page - 1) * query.PageSize).Find(&WOA).Error
	if err != nil {
		return nil, 0, err
	}

	return WOA, total, nil
}

func SearchWOAInfo(query *request.Query) ([]model.WebChatOfficeAccount, int64, error) {
	var err error
	var WOA = []model.WebChatOfficeAccount{}
	var total int64
	if err = db2.Orm.Model(&model.WebChatOfficeAccount{}).Where(query.Type+" like ?", "%"+query.Keyword+"%").Count(&total).Error; err != nil {
		return nil, 0, err
	}
	if err = db2.Orm.Model(&model.WebChatOfficeAccount{}).Limit(query.PageSize).Offset((query.Page-1)*query.PageSize).Where(query.Type+" like ?", "%"+query.Keyword+"%").Find(&WOA).Error; err != nil {
		return nil, 0, err
	}

	return WOA, total, nil
}

func ReadFlagWOAInfoById(pids *request.ParamIds) error {
	id := strings.Split(pids.Id, ",")

	for _, v := range id {
		err := db2.Orm.Model(&model.WebChatOfficeAccount{}).Where("id=?", v).Update("isNew", false).Error
		if err != nil {
			return err
		}
	}
	return nil
}

func ReadAllFlagWOAInfo() error {
	err := db2.Orm.Model(&model.WebChatOfficeAccount{}).Where("1=1").Update("isNew", false).Error
	if err != nil {
		return err
	}
	return nil
}
