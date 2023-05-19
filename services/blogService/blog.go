package blogService

import (
	"strconv"
	"time"
	"wkg/model"
	"wkg/model/request"
	"wkg/model/response"
	"wkg/pkg/db"
)

func GetShareknowledgeList(page *request.PageParam) ([]model.ListData, int64, error) {

	sknow := []model.Knowledge{}
	if err := db.Orm.Model(&model.Knowledge{}).Limit(page.PageSize).Offset((page.Page-1)*page.PageSize).Where("shared=?", true).Order("id desc").Find(&sknow).Error; err != nil {
		return nil, 0, err
	}

	lds := []model.ListData{}

	dMap := make(map[int][]model.Knows)

	for i := 0; i < len(sknow); i++ {
		k := model.Knows{}
		ct, err := time.Parse("2006-01-02 15:04:05", sknow[i].CreateTime)
		if err != nil {
			return nil, 0, err
		}

		k.Href = "/shareview?key=" + sknow[i].CKey
		k.Title = sknow[i].Title
		k.CreateTime = ct.Format("Jan,02,2006")
		//同样年份的放到同一key里
		dMap[ct.Year()] = append(dMap[ct.Year()], k)

	}

	for k, v := range dMap {
		ld := model.ListData{}
		ld.Year = strconv.Itoa(k)

		ld.Knows = v
		lds = append(lds, ld)
	}

	var total int64
	if err := db.Orm.Model(&model.Knowledge{}).Where("shared=?", true).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	return lds, total, nil
}

func GetShareknowledgeByKey(key string) (*model.Knowledge, error) {
	sknow := &model.Knowledge{}
	if err := db.Orm.Model(&model.Knowledge{}).Where("ckey=? and shared=?", key, true).Find(&sknow).Error; err != nil {
		return nil, err
	}

	//更新hit
	sknow.Hit += 1
	if err := db.Orm.Model(&model.Knowledge{}).Where("ckey=?", sknow.CKey).Update("hit", sknow.Hit).Error; err != nil {
		return nil, err
	}

	return sknow, nil
}

func SearchShareKnowledgeByWord(page *request.PageParam) ([]model.ListData, int64, error) {

	sknow := []model.Knowledge{}
	if err := db.Orm.Model(&model.Knowledge{}).Limit(page.PageSize).Offset((page.Page-1)*page.PageSize).Where("shared=? and title like ?", true, "%"+page.KeyWord+"%").Order("id desc").Find(&sknow).Error; err != nil {
		return nil, 0, err
	}

	lds := []model.ListData{}

	dMap := make(map[int][]model.Knows)

	for i := 0; i < len(sknow); i++ {
		k := model.Knows{}
		ct, err := time.Parse("2006-01-02 15:04:05", sknow[i].CreateTime)
		if err != nil {
			return nil, 0, err
		}

		k.Href = "/shareview?key=" + sknow[i].CKey
		k.Title = sknow[i].Title
		k.CreateTime = ct.Format("Jan,02,2006")
		//同样年份的放到同一key里
		dMap[ct.Year()] = append(dMap[ct.Year()], k)

	}

	for k, v := range dMap {
		ld := model.ListData{}
		ld.Year = strconv.Itoa(k)

		ld.Knows = v
		lds = append(lds, ld)
	}

	var total int64
	if err := db.Orm.Model(&model.Knowledge{}).Where("shared=? and title like ?", true, "%"+page.KeyWord+"%").Count(&total).Error; err != nil {
		return nil, 0, err
	}

	return lds, total, nil
}

func GetsharedKnowledgeListbyTag(page *request.PageParam) ([]model.ListData, int64, error) {
	sknow := []model.Knowledge{}
	if err := db.Orm.Model(&model.Knowledge{}).Limit(page.PageSize).Offset((page.Page-1)*page.PageSize).Where("shared=? and tags like ?", true, "%"+page.KeyWord+"%").Order("id desc").Find(&sknow).Error; err != nil {
		return nil, 0, err
	}

	lds := []model.ListData{}

	dMap := make(map[int][]model.Knows)

	for i := 0; i < len(sknow); i++ {
		k := model.Knows{}
		ct, err := time.Parse("2006-01-02 15:04:05", sknow[i].CreateTime)
		if err != nil {
			return nil, 0, err
		}

		k.Href = "/shareview?key=" + sknow[i].CKey
		k.Title = sknow[i].Title
		k.CreateTime = ct.Format("Jan,02,2006")
		//同样年份的放到同一key里
		dMap[ct.Year()] = append(dMap[ct.Year()], k)

	}

	for k, v := range dMap {
		ld := model.ListData{}
		ld.Year = strconv.Itoa(k)

		ld.Knows = v
		lds = append(lds, ld)
	}

	var total int64
	if err := db.Orm.Model(&model.Knowledge{}).Where("shared=? and title like ?", true, "%"+page.KeyWord+"%").Count(&total).Error; err != nil {
		return nil, 0, err
	}

	return lds, total, nil

}

func GetTags() ([]response.RespTags, error) {

	knows := []model.Knowledge{}

	if err := db.Orm.Model(&model.Knowledge{}).Where("shared=true").Order("id desc").Find(&knows).Error; err != nil {
		return nil, err
	}

	rtags := []response.RespTags{}
	tagMaps := make(map[string]int)
	for _, v := range knows {
		for _, vt := range v.Tags {
			if _, ok := tagMaps[vt]; !ok {
				tagMaps[vt] = 1
			} else {
				tagMaps[vt] = tagMaps[vt] + 1
			}
		}
	}

	for v, k := range tagMaps {
		rtag := response.RespTags{}
		rtag.TagName = v
		rtag.Count = k
		rtags = append(rtags, rtag)
	}

	return rtags, nil
}
