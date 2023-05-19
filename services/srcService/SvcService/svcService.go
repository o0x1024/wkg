package SvcService

import (
	"strings"
	"wkg/model"
	"wkg/model/request"
	db2 "wkg/pkg/db"
)

func GetServiceInfo(query *request.Query) ([]model.Services, int64, error) {
	var err error
	var service = []model.Services{}
	var total int64
	if query.AssetOption == "1" {
		if query.Keyword != "" {
			if err = db2.Orm.Model(&model.Services{}).Where(query.Type+" like ? and isNew=true", "%"+query.Keyword+"%").Count(&total).Error; err != nil {
				return nil, 0, err
			}
			if err = db2.Orm.Model(&model.Services{}).Limit(query.PageSize).Offset((query.Page-1)*query.PageSize).Where(query.Type+" like ? and isNew=tru", "%"+query.Keyword+"%").Find(&service).Error; err != nil {
				return nil, 0, err
			}
		} else {
			if err = db2.Orm.Model(&model.Services{}).Where("isNew=true").Count(&total).Error; err != nil {
				return nil, 0, err
			}
			if err = db2.Orm.Model(&model.Services{}).Limit(query.PageSize).Offset((query.Page - 1) * query.PageSize).Where("isNew=true").Find(&service).Error; err != nil {
				return nil, 0, err
			}
		}
	} else {
		if query.Keyword != "" {
			if err = db2.Orm.Model(&model.Services{}).Where(query.Type+" like ?", "%"+query.Keyword+"%").Count(&total).Error; err != nil {
				return nil, 0, err
			}
			if err = db2.Orm.Model(&model.Services{}).Limit(query.PageSize).Offset((query.Page-1)*query.PageSize).Where(query.Type+" like ?", "%"+query.Keyword+"%").Find(&service).Error; err != nil {
				return nil, 0, err
			}
		} else {
			if err = db2.Orm.Model(&model.Services{}).Count(&total).Error; err != nil {
				return nil, 0, err
			}
			if err = db2.Orm.Model(&model.Services{}).Limit(query.PageSize).Offset((query.Page - 1) * query.PageSize).Find(&service).Error; err != nil {
				return nil, 0, err
			}
		}
	}

	return service, total, nil
}

func GetServiceInfoByCid(query *request.Query) ([]model.Services, int64, error) {

	var svc []model.Services
	var total int64

	if query.AssetOption == "0" {
		err := db2.Orm.Model(&model.Services{}).Where("cid=?", query.Cid).Limit(query.PageSize).Offset((query.Page - 1) * query.PageSize).Find(&svc).Error
		if err != nil {
			return nil, 0, err
		}

		//查询总数
		err = db2.Orm.Model(&model.Services{}).Where("cid=?", query.Cid).Count(&total).Error
		if err != nil {
			return nil, 0, err
		}
	} else {
		err := db2.Orm.Model(&model.Services{}).Where("cid=? and isNew=true", query.Cid).Limit(query.PageSize).Offset((query.Page - 1) * query.PageSize).Find(&svc).Error
		if err != nil {
			return nil, 0, err
		}

		//查询总数
		err = db2.Orm.Model(&model.Services{}).Where("cid=? and isNew=true", query.Cid).Count(&total).Error
		if err != nil {
			return nil, 0, err
		}
	}

	return svc, total, nil

}

func DelServiceById(id string) error {
	var svc model.Services
	if err := db2.Orm.Where("id=?", id).Delete(&svc).Error; err != nil {
		return err
	}
	return nil
}

func GetNewServiceInfo(query *request.Query) ([]model.Domain, int64, error) {
	var err error
	var total int64
	//查询总数
	err = db2.Orm.Model(&model.Ips{}).Where("isNew=true").Count(&total).Error
	if err != nil {
		return nil, 0, err
	}
	//根据page和pagesize查询数据
	var dom = []model.Domain{}
	err = db2.Orm.Model(&model.Domain{}).Where("isNew=true").Limit(query.PageSize).Offset((query.Page - 1) * query.PageSize).Find(&dom).Error
	if err != nil {
		return nil, 0, err
	}

	return dom, total, nil
}

func SearchServiceInfo(query *request.Query) ([]model.Services, int64, error) {
	var err error
	var service = []model.Services{}
	var total int64
	if err = db2.Orm.Model(&model.Services{}).Where(query.Type+" like ?", "%"+query.Keyword+"%").Count(&total).Error; err != nil {
		return nil, 0, err
	}
	if err = db2.Orm.Model(&model.Services{}).Limit(query.PageSize).Offset((query.Page-1)*query.PageSize).Where(query.Type+" like ?", "%"+query.Keyword+"%").Find(&service).Error; err != nil {
		return nil, 0, err
	}

	return service, total, nil
}

func ReadFlagServiceInfoById(pids *request.ParamIds) error {
	id := strings.Split(pids.Id, ",")

	for _, v := range id {
		err := db2.Orm.Model(&model.Services{}).Where("id=?", v).Update("isNew", false).Error
		if err != nil {
			return err
		}
	}
	return nil
}

func ReadAllFlagServiceInfo() error {
	err := db2.Orm.Model(&model.Services{}).Where("1=1").Update("isNew", false).Error
	if err != nil {
		return err
	}
	return nil
}
