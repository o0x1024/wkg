package IPService

import (
	"strings"
	"wkg/model"
	"wkg/model/request"
	db2 "wkg/pkg/db"
)

func GetIPsInfo(query *request.Query) ([]model.Ips, int64, error) {

	var total int64
	var IPs []model.Ips
	if query.AssetOption == "1" {
		//查询总数
		err := db2.Orm.Model(&model.Ips{}).Where("ip like ? and isNew=true", "%"+query.Keyword+"%").Count(&total).Error
		if err != nil {
			return nil, 0, err
		}
		//根据page和pagesize查询数据

		err = db2.Orm.Model(&model.Ips{}).Where("ip like ? and isNew=true", "%"+query.Keyword+"%").Limit(query.PageSize).Offset((query.Page - 1) * query.PageSize).Find(&IPs).Error
		if err != nil {
			return nil, 0, err
		}

	} else {
		//查询总数
		err := db2.Orm.Model(&model.Ips{}).Where("ip like ? ", "%"+query.Keyword+"%").Count(&total).Error
		if err != nil {
			return nil, 0, err
		}
		//根据page和pagesize查询数据
		err = db2.Orm.Model(&model.Ips{}).Where("ip like ?", "%"+query.Keyword+"%").Limit(query.PageSize).Offset((query.Page - 1) * query.PageSize).Find(&IPs).Error
		if err != nil {
			return nil, 0, err
		}

	}

	return IPs, total, nil
}

func SearchIPsInfo(query *request.Query) ([]model.Ips, int64, error) {
	var err error
	var ips = []model.Ips{}

	if err = db2.Orm.Model(&model.Ips{}).Where("ip LIKE ?", "%"+query.Keyword+"%").Find(&ips).Error; err != nil {
		return nil, 0, nil
	}

	var total int64
	if err = db2.Orm.Model(&model.Ips{}).Where("ip LIKE ?", "%"+query.Keyword+"%").Count(&total).Error; err != nil {
		return nil, 0, nil
	}

	return ips, total, nil
}

func GetNewIPsInfo(query *request.Query) ([]model.Ips, int64, error) {
	var err error
	var total int64

	//查询总数
	err = db2.Orm.Model(&model.Ips{}).Where("isNew=true").Count(&total).Error
	if err != nil {
		return nil, 0, err
	}
	//根据page和pagesize查询数据
	var Ips []model.Ips
	err = db2.Orm.Model(&model.Ips{}).Where("isNew=true").Limit(query.PageSize).Offset((query.Page - 1) * query.PageSize).Find(&Ips).Error
	if err != nil {
		return nil, 0, err
	}

	return Ips, total, nil
}

func ReadAllFlagIPInfo() error {
	err := db2.Orm.Model(&model.Ips{}).Where("1=1").Update("isNew", false).Error
	if err != nil {
		return err
	}
	return nil
}

func ReadFlagIPInfoById(cids *request.ParamIds) error {
	id := strings.Split(cids.Id, ",")

	for _, v := range id {
		err := db2.Orm.Model(&model.Ips{}).Where("id=?", v).Update("isNew", false).Error
		if err != nil {
			return err
		}
	}
	return nil
}

func DelIPById(id string) error {
	var ips model.Ips
	err := db2.Orm.Where("id=?", id).Delete(&ips).Error
	if err != nil {
		return err
	}
	return nil
}

func GetIPsInfoByCid(query *request.Query) ([]model.Ips, int64, error) {
	var IPs []model.Ips
	var total int64
	if query.AssetOption == "0" {
		err := db2.Orm.Model(&model.Ips{}).Where("cid=?", query.Cid).Limit(query.PageSize).Offset((query.Page - 1) * query.PageSize).Find(&IPs).Error
		if err != nil {
			return nil, 0, err
		}
		//查询总数
		err = db2.Orm.Model(&model.Ips{}).Where("cid=?", query.Cid).Count(&total).Error
		if err != nil {
			return nil, 0, err
		}

	} else {
		err := db2.Orm.Model(&model.Ips{}).Where("cid=? and isNew=true", query.Cid).Limit(query.PageSize).Offset((query.Page - 1) * query.PageSize).Find(&IPs).Error
		if err != nil {
			return nil, 0, err
		}
		//查询总数
		err = db2.Orm.Model(&model.Ips{}).Where("cid=? and isNew=true", query.Cid).Count(&total).Error
		if err != nil {
			return nil, 0, err
		}

	}

	return IPs, total, nil
}
