package miniprogramService

import (
	"strings"
	"wkg/model"
	"wkg/model/request"
	db2 "wkg/pkg/db"
)

func GetMiniInfo(query *request.Query) ([]model.MiniProgram, int64, error) {
	var err error
	var mini = []model.MiniProgram{}
	var total int64

	if query.Keyword != "" {
		if err = db2.Orm.Model(&model.MiniProgram{}).Where(query.Type+" like ?", "%"+query.Keyword+"%").Count(&total).Error; err != nil {
			return nil, 0, err
		}
		if err = db2.Orm.Model(&model.MiniProgram{}).Limit(query.PageSize).Offset((query.Page-1)*query.PageSize).Where(query.Type+" like ?", "%"+query.Keyword+"%").Find(&mini).Error; err != nil {
			return nil, 0, err
		}
	} else {
		if err = db2.Orm.Model(&model.MiniProgram{}).Count(&total).Error; err != nil {
			return nil, 0, err
		}
		if err = db2.Orm.Model(&model.MiniProgram{}).Limit(query.PageSize).Offset((query.Page - 1) * query.PageSize).Find(&mini).Error; err != nil {
			return nil, 0, err
		}
	}

	return mini, total, nil
}

func GetMiniInfoByCid(query *request.Query) ([]model.MiniProgram, int64, error) {

	var svc []model.MiniProgram
	err := db2.Orm.Model(&model.MiniProgram{}).Where("cid=?", query.Cid).Limit(query.PageSize).Offset((query.Page - 1) * query.PageSize).Find(&svc).Error
	if err != nil {
		return nil, 0, err
	}

	var total int64
	//查询总数
	err = db2.Orm.Model(&model.MiniProgram{}).Where("cid=?", query.Cid).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}
	return svc, total, nil

}

func DelMiniById(id string) error {
	var mini model.MiniProgram
	if err := db2.Orm.Where("id=?", id).Delete(&mini).Error; err != nil {
		return err
	}
	return nil
}

func GetNewMiniInfo(query *request.Query) ([]model.MiniProgram, int64, error) {
	var err error
	var total int64
	//查询总数
	err = db2.Orm.Model(&model.MiniProgram{}).Where("isNew=true").Count(&total).Error
	if err != nil {
		return nil, 0, err
	}
	//根据page和pagesize查询数据
	var mini []model.MiniProgram
	err = db2.Orm.Model(&model.MiniProgram{}).Where("isNew=true").Limit(query.PageSize).Offset((query.Page - 1) * query.PageSize).Find(&mini).Error
	if err != nil {
		return nil, 0, err
	}

	return mini, total, nil
}

func SearchMiniInfo(query *request.Query) ([]model.MiniProgram, int64, error) {
	var err error
	var mini = []model.MiniProgram{}
	var total int64
	if err = db2.Orm.Model(&model.MiniProgram{}).Where(query.Type+" like ?", "%"+query.Keyword+"%").Count(&total).Error; err != nil {
		return nil, 0, err
	}
	if err = db2.Orm.Model(&model.MiniProgram{}).Limit(query.PageSize).Offset((query.Page-1)*query.PageSize).Where(query.Type+" like ?", "%"+query.Keyword+"%").Find(&mini).Error; err != nil {
		return nil, 0, err
	}

	return mini, total, nil
}

func ReadFlagMiniInfoById(pids *request.ParamIds) error {
	id := strings.Split(pids.Id, ",")

	for _, v := range id {
		err := db2.Orm.Model(&model.MiniProgram{}).Where("id=?", v).Update("isNew", false).Error
		if err != nil {
			return err
		}
	}
	return nil
}

func ReadAllFlagMiniInfo() error {
	err := db2.Orm.Model(&model.MiniProgram{}).Where("1=1").Update("isNew", false).Error
	if err != nil {
		return err
	}
	return nil
}
