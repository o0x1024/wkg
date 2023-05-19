package knowledgeService

import (
	"errors"
	"strconv"
	"time"
	"wkg/conf"
	"wkg/model"
	"wkg/model/request"
	db2 "wkg/pkg/db"
	helper2 "wkg/pkg/libs/helper"
)

type KnowledgeService struct {
}

type Option struct {
	Value string `json:"value"`
	Lable string `json:"label"`
}

type KList struct {
	Key        string `json:"key"`
	Title      string `json:"title"`
	UpdateTime string `json:"updateTime"`
}

func Shareknowledge(key string) (string, error) {

	know := model.Knowledge{}
	if err := db2.Orm.Model(&model.Knowledge{}).Where("ckey=?", key).Find(&know).Error; err != nil {
		return "", err
	}

	*know.Shared = !*know.Shared
	if err := db2.Orm.Model(&model.Knowledge{}).Where("ckey=?", key).Update("shared", *know.Shared).Error; err != nil {
		return "", err
	}
	shareUrl := conf.Cfg.BaseUrl + "shareview?key=" + key
	return shareUrl, nil

}

func GetHotList() ([]request.ListData, error) {

	var err error
	sknow := []model.Knowledge{}
	if err = db2.Orm.Model(&model.Knowledge{}).Limit(10).Order("hit desc").Where("shared=?", true).Find(&sknow).Error; err != nil {
		return nil, err
	}

	lds := []request.ListData{}
	//处理一下内容的长度
	for i := 0; i < len(sknow); i++ {
		ld := request.ListData{}
		ld.Href = "/shareview?key=" + sknow[i].CKey
		ld.Title = sknow[i].Title
		ld.Hit = sknow[i].Hit
		lds = append(lds, ld)
	}

	return lds, err
}

func SearchShareKnowledgeByWord(page *request.PageParam) ([]model.ListData, int64, error) {

	sknow := []model.Knowledge{}
	if err := db2.Orm.Model(&model.Knowledge{}).Limit(page.PageSize).Offset((page.Page-1)*page.PageSize).Where("shared=? and title like ?", true, "%"+page.KeyWord+"%").Order("id desc").Find(&sknow).Error; err != nil {
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
	if err := db2.Orm.Model(&model.Knowledge{}).Where("shared=? and title like ?", true, "%"+page.KeyWord+"%").Count(&total).Error; err != nil {
		return nil, 0, err
	}

	return lds, total, nil
}

func GetKnowledgeCategoriesList(key string) ([]model.Knowledge, error) {
	knowledge := []model.Knowledge{}
	var pid int
	err := db2.Orm.Model(&model.Category{}).Select("id").Where("ckey=?", key).Find(&pid).Error
	if err != nil {
		return nil, err
	}
	err = db2.Orm.Model(&model.Knowledge{}).Select("title,ckey,isLeaf").Where("parentId=? ", pid).Find(&knowledge).Error
	if err != nil {
		return nil, err
	}
	return knowledge, nil
}

func buildTree(categories []model.CategoryTree) []*model.CategoryTree {

	// 存储每个类别节点
	treeMap := make(map[string]*model.CategoryTree, len(categories))

	// 遍历切片，将每个节点添加到字典中
	for i := range categories {
		treeMap[strconv.Itoa(categories[i].Id)] = &categories[i]
	}

	// 再遍历切片，将每个节点的父节点与子节点建立关联
	var roots []*model.CategoryTree
	for i := range categories {
		if categories[i].ParentId == 0 {
			roots = append(roots, &categories[i])
		} else if parent, ok := treeMap[strconv.Itoa(categories[i].ParentId)]; ok {
			parent.Children = append(parent.Children, &categories[i])
		}
	}

	return roots
}

func GetTree() ([]*model.CategoryTree, error) {
	categories := []model.Category{}

	// var treeCache = make(map[int][]model.CategoryTree)
	//一级菜单
	err := db2.Orm.Model(&model.Category{}).Select("id,parentId,title,ckey,isLeaf").Order("parentId ASC").Find(&categories).Error
	if err != nil {
		return nil, err
	}

	var _trees []model.CategoryTree
	var _tree model.CategoryTree
	for i := 0; i < len(categories); i++ {
		_tree.Id = categories[i].Id
		_tree.ParentId = categories[i].ParentId
		_tree.Key = categories[i].CKey
		_tree.Title = categories[i].Title
		_tree.IsLeaf = categories[i].IsLeaf
		_trees = append(_trees, _tree)
	}

	res := buildTree(_trees)
	return res, err
}

func GetKnowledgeByKey(key string) (*model.Knowledge, error) {
	knowledge := model.Knowledge{}

	err := db2.Orm.Model(&model.Knowledge{}).Where("ckey=? ", key).Find(&knowledge).Error
	if err != nil {
		return nil, err
	}

	return &knowledge, nil
}

func GetSummary() ([]KList, error) {
	var knowledgeList []KList
	var know []model.Knowledge

	err := db2.Orm.Model(&model.Knowledge{}).Select("ckey,title,updateTime").Order("updateTime desc").Limit(10).Find(&know).Error
	if err != nil {
		return nil, err
	}

	for _, v := range know {
		var _klist KList
		_klist.Key = v.CKey
		_klist.Title = v.Title
		_klist.UpdateTime = v.UpdateTime
		knowledgeList = append(knowledgeList, _klist)
	}

	return knowledgeList, nil
}

func AddNode(parentId int, nodeName string) (*model.Category, error) {
	var cate = &model.Category{}

	if parentId != 0 {
		cate.ParentId = parentId
	}

	cate.Title = nodeName
	flg := false
	cate.IsLeaf = &flg
	cate.CKey = helper2.Md5(time.Unix(time.Now().Unix(), 0).Format("2006-01-02 15:04:05"))

	err := db2.Orm.Model(&model.Category{}).Create(&cate).Error
	if err != nil {
		return cate, err
	}
	return cate, nil
}

func GetKnowlegeInBackend(page *request.PageParam) ([]model.Knowledge, int64, error) {

	sknow := []model.Knowledge{}
	if err := db2.Orm.Debug().Model(&model.Knowledge{}).Where("shared=?", true).Limit(page.PageSize).Offset((page.Page - 1) * page.PageSize).Order("updateTime desc").Find(&sknow).Error; err != nil {
		return nil, 0, err
	}

	var total int64
	if err := db2.Orm.Model(&model.Knowledge{}).Where("shared=?", true).Count(&total).Error; err != nil {
		return nil, 0, err
	}
	return sknow, total, nil
}

func SaveEditKnowledge(tags model.Tag, title string, content string, key string) (*model.Category, error) {

	var err error
	var cate = &model.Category{}

	err = db2.Orm.Model(&model.Category{}).Where("ckey=?", key).Find(&cate).Error
	if err != nil {
		return nil, err
	}

	//更新一下title返回前端
	cate.Title = title

	now := time.Unix(time.Now().Unix(), 0).Format("2006-01-02 15:04:05")
	var tKnow model.Knowledge
	tKnow.Content = content

	tKnow.Title = title
	tKnow.Tags = tags
	tKnow.UpdateTime = now
	err = db2.Orm.Model(&model.Knowledge{}).Where("ckey=?", key).Updates(&tKnow).Error
	if err != nil {
		return nil, err
	}
	return cate, nil
}

func SaveNewKnowledge(ParentId int, tags model.Tag, username string, parentId int, title string, content string) (*model.Knowledge, *model.Category, error) {
	var err error
	var tKnow = &model.Knowledge{}
	var cate = &model.Category{}

	now := time.Unix(time.Now().Unix(), 0).Format("2006-01-02 15:04:05")

	tms := false

	tKnow.ParentId = ParentId
	tKnow.Author = username
	tKnow.CreateTime = now
	tKnow.Content = content
	tKnow.Title = title
	tKnow.CKey = helper2.Md5(now)
	tKnow.UpdateTime = now
	tKnow.Shared = &tms
	tKnow.Tags = tags

	cate.ParentId = parentId
	cate.Title = title
	cate.CKey = tKnow.CKey
	flg := true
	cate.IsLeaf = &flg

	err = db2.Orm.Model(&model.Knowledge{}).Create(&tKnow).Error
	if err != nil {
		return nil, nil, err
	}

	err = db2.Orm.Model(&model.Category{}).Create(&cate).Error
	if err != nil {
		return nil, nil, err
	}

	return tKnow, cate, nil
}

func UpdateCatelogTItle(title string, key string) error {
	ct := &model.Category{}
	if err := db2.Orm.Model(&model.Category{}).Where("ckey=?", key).Find(&ct).Error; err != nil {
		return err
	}

	ct.Title = title
	if err := db2.Orm.Where("ckey=?", key).Updates(&ct).Error; err != nil {
		return err
	}
	return nil
}

// TD:递归删除目录下的所有节点
func DelTreeNode(isLeaf bool, id int) error {
	var err error
	if isLeaf {
		err = db2.Orm.Model(&model.Knowledge{}).Where("id=?", id).Delete(&model.Knowledge{}).Error
		if err != nil {
			return err
		}
		err := db2.Orm.Model(&model.Category{}).Where("id=?", id).Delete(&model.Category{}).Error
		if err != nil {
			return err
		}
	} else {

		err := db2.Orm.Model(&model.Category{}).Where("id=?", id).Delete(&model.Category{}).Error
		if err != nil {
			return err
		}

		err = db2.Orm.Model(&model.Category{}).Where("parentId=?", id).Delete(&model.Category{}).Error
		if err != nil {
			return err
		}

		err = db2.Orm.Model(&model.Knowledge{}).Where("parentId=?", id).Delete(&model.Knowledge{}).Error
		if err != nil {
			return err
		}
	}

	return nil
}

func ModifyTreeNode(isLeaf bool, id int, keyword string) error {

	if !isLeaf {
		err := db2.Orm.Model(&model.Category{}).Where("id=?", id).Update("title", keyword).Error
		if err != nil {
			return err
		}
	} else {
		return errors.New("非目录节点无法修改")
	}

	return nil
}

// func DelTag(id int) error {
// 	err := db2.Orm.Model(&model.Category{}).Where("id=?", id).Delete(&model.Tag{}).Error
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// func AddTag(key string) error {
// 	tag := &model.Tag{}
// 	tag.Key = key

// 	if err := db2.Orm.Create(&tag).Error; err != nil {
// 		return err
// 	}
// 	return nil
// }
