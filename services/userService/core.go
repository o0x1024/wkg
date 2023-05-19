package userService

import (
	"time"
	"wkg/model"
	"wkg/model/request"
	db2 "wkg/pkg/db"
)

func SaveUser(user *model.Users) error {

	user.CreateTime = time.Unix(time.Now().Unix(), 0).Format("2006-01-02 15:04:05")
	err := db2.Orm.Model(&model.Users{}).Create(&user).Error
	if err != nil {
		return err
	}
	return nil

}

func DelUser(id string) error {
	var user model.Users
	if err := db2.Orm.Where("id=?", id).Delete(&user).Error; err != nil {
		return err
	}
	return nil
}

func ResetPwd(id string) error {
	var user = &model.Users{}

	if err := db2.Orm.Model(&model.Users{}).Where("id=?", id).Find(user).Error; err != nil {
		return err
	}
	user.Password = "u^b^4oaWQ@pcQ"
	if err := db2.Orm.Model(&model.Users{}).Where("id=?", id).Updates(&user).Error; err != nil {
		return err
	}
	return nil
}

func GetUserList(query *request.Query) ([]model.Users, int64, error) {
	var err error
	var user = []model.Users{}
	var total int64
	if query.Keyword != "" {
		if err = db2.Orm.Model(&model.Users{}).Where(query.Type+" like ?", "%"+query.Keyword+"%").Count(&total).Error; err != nil {
			return nil, 0, err
		}
		if err = db2.Orm.Model(&model.Users{}).Limit(query.PageSize).Offset((query.Page-1)*query.PageSize).Where(query.Type+" like ?", "%"+query.Keyword+"%").Find(&user).Error; err != nil {
			return nil, 0, err
		}
	} else {
		if err = db2.Orm.Model(&model.Users{}).Count(&total).Error; err != nil {
			return nil, 0, err
		}
		if err = db2.Orm.Model(&model.Users{}).Limit(query.PageSize).Offset((query.Page - 1) * query.PageSize).Find(&user).Error; err != nil {
			return nil, 0, err
		}
	}

	return user, total, nil
}
