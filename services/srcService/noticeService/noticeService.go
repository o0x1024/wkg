package noticeService

import (
	"time"
	"wkg/model"
	"wkg/model/request"
	db2 "wkg/pkg/db"
)

func GetNoticeByKeyword(query *request.Query) ([]model.Notice, int64, error) {
	var err error
	var Notice = []model.Notice{}
	var total int64

	if query.Keyword != "" {
		if err = db2.Orm.Model(&model.Notice{}).Where("title like ?", "%"+query.Keyword+"%").Count(&total).Error; err != nil {
			return nil, 0, err
		}
		if err = db2.Orm.Model(&model.Notice{}).Limit(query.PageSize).Offset((query.Page-1)*query.PageSize).Where("title like ?", "%"+query.Keyword+"%").Find(&Notice).Error; err != nil {
			return nil, 0, err
		}
	} else {
		if err = db2.Orm.Model(&model.Notice{}).Count(&total).Error; err != nil {
			return nil, 0, err
		}
		if err = db2.Orm.Model(&model.Notice{}).Limit(query.PageSize).Offset((query.Page - 1) * query.PageSize).Find(&Notice).Error; err != nil {
			return nil, 0, err
		}
	}

	return Notice, total, nil
}

func GetNoticeDetailById(id string) (*model.Notice, error) {

	var notice = &model.Notice{}
	err := db2.Orm.Model(&model.Notice{}).Where("id=? ", id).Find(&notice).Error
	if err != nil {
		return nil, err
	}

	return notice, nil
}

func SaveNotice(notice *model.Notice) error {

	notice.UpdateTime = time.Unix(time.Now().Unix(), 0).Format("2006-01-02 15:04:05")

	err := db2.Orm.Model(&model.Notice{}).Create(&notice).Error
	if err != nil {
		return err
	}
	return nil
}

func SaveEditNotice(notice *model.Notice) error {
	notice.UpdateTime = time.Unix(time.Now().Unix(), 0).Format("2006-01-02 15:04:05")

	err := db2.Orm.Model(&model.Notice{}).Where("id=? and cid=?", notice.Id, notice.Cid).Updates(&notice).Error
	if err != nil {
		return err
	}
	return nil
}

func GetNoticeList(query *request.Query) ([]model.Notice, int64, error) {

	var svc []model.Notice
	err := db2.Orm.Model(&model.Notice{}).Select("id,title,updateTime").Limit(query.PageSize).Offset((query.Page - 1) * query.PageSize).Order("updateTime desc").Find(&svc).Error
	if err != nil {
		return nil, 0, err
	}

	var total int64
	//查询总数
	err = db2.Orm.Model(&model.Notice{}).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	return svc, total, nil

}

func DelNoticeById(id string) error {
	var Notice model.Notice
	if err := db2.Orm.Where("id=?", id).Delete(&Notice).Error; err != nil {
		return err
	}
	return nil
}
