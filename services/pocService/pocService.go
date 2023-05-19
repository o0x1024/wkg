package pocService

import (
	"strings"
	"time"
	"wkg/model"
	"wkg/model/request"
	db2 "wkg/pkg/db"
)

func GetPocInfo(query *request.Query) ([]model.Poc, int64, error) {
	var err error
	var Poc = []model.Poc{}
	var total int64

	if query.Keyword != "" {
		if err = db2.Orm.Model(&model.Poc{}).Where("pocName like ?", "%"+query.Keyword+"%").Count(&total).Error; err != nil {
			return nil, 0, err
		}
		if err = db2.Orm.Model(&model.Poc{}).Limit(query.PageSize).Offset((query.Page-1)*query.PageSize).Where("pocName like ?", "%"+query.Keyword+"%").Find(&Poc).Error; err != nil {
			return nil, 0, err
		}
	} else {
		if err = db2.Orm.Model(&model.Poc{}).Count(&total).Error; err != nil {
			return nil, 0, err
		}
		if err = db2.Orm.Model(&model.Poc{}).Limit(query.PageSize).Offset((query.Page - 1) * query.PageSize).Find(&Poc).Error; err != nil {
			return nil, 0, err
		}
	}

	return Poc, total, nil
}

func GetPocDetailById(id string) (*model.Poc, error) {

	var Poc = &model.Poc{}
	err := db2.Orm.Model(&model.Poc{}).Where("id=? ", id).Find(&Poc).Error
	if err != nil {
		return nil, err
	}

	return Poc, nil
}

func SavePoc(Poc *model.Poc) error {

	Poc.UpdateTime = time.Unix(time.Now().Unix(), 0).Format("2006-01-02 15:04:05")
	err := db2.Orm.Model(&model.Poc{}).Create(&Poc).Error
	if err != nil {
		return err
	}
	return nil
}

func SaveEditPoc(Poc *model.Poc) error {
	Poc.UpdateTime = time.Unix(time.Now().Unix(), 0).Format("2006-01-02 15:04:05")

	err := db2.Orm.Model(&model.Poc{}).Where("id=?", Poc.Id).Updates(&Poc).Error
	if err != nil {
		return err
	}
	return nil
}

func GetPocInfoByCid(query *request.Query) ([]model.Poc, int64, error) {

	var svc []model.Poc
	err := db2.Orm.Model(&model.Poc{}).Select("id,pocName,updateTime").Limit(query.PageSize).Offset((query.Page - 1) * query.PageSize).Find(&svc).Error
	if err != nil {
		return nil, 0, err
	}

	var total int64
	//查询总数
	err = db2.Orm.Model(&model.Poc{}).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	return svc, total, nil

}

func DelPocById(id string) error {
	var Poc model.Poc
	if err := db2.Orm.Where("id=?", id).Delete(&Poc).Error; err != nil {
		return err
	}
	return nil
}

func GetNewPocInfo(query *request.Query) ([]model.Poc, int64, error) {
	var err error
	var total int64
	//查询总数
	err = db2.Orm.Model(&model.Poc{}).Where("isNew=true").Count(&total).Error
	if err != nil {
		return nil, 0, err
	}
	//根据page和pagesize查询数据
	var Poc []model.Poc
	err = db2.Orm.Model(&model.Poc{}).Where("isNew=true").Limit(query.PageSize).Offset((query.Page - 1) * query.PageSize).Find(&Poc).Error
	if err != nil {
		return nil, 0, err
	}

	return Poc, total, nil
}

func SearchPocInfo(query *request.Query) ([]model.Poc, int64, error) {
	var err error
	var Poc = []model.Poc{}
	var total int64
	if err = db2.Orm.Model(&model.Poc{}).Where(query.Type+" like ?", "%"+query.Keyword+"%").Count(&total).Error; err != nil {
		return nil, 0, err
	}
	if err = db2.Orm.Model(&model.Poc{}).Limit(query.PageSize).Offset((query.Page-1)*query.PageSize).Where(query.Type+" like ?", "%"+query.Keyword+"%").Find(&Poc).Error; err != nil {
		return nil, 0, err
	}

	return Poc, total, nil
}

func ReadFlagPocInfoById(pids *request.ParamIds) error {
	id := strings.Split(pids.Id, ",")

	for _, v := range id {
		err := db2.Orm.Model(&model.Poc{}).Where("id=?", v).Update("isNew", false).Error
		if err != nil {
			return err
		}
	}
	return nil
}

func ReadAllFlagPocInfo() error {
	err := db2.Orm.Model(&model.Poc{}).Where("1=1").Update("isNew", false).Error
	if err != nil {
		return err
	}
	return nil
}
