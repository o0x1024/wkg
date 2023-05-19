package knowledge

import (
	"net/http"
	"strconv"
	"strings"
	"time"
	"wkg/model"
	"wkg/model/request"
	db2 "wkg/pkg/db"
	helper2 "wkg/pkg/libs/helper"
	Gjwt "wkg/pkg/libs/jwt"
	"wkg/services/knowledgeService"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type KnowStrut struct {
	Tags     []string `json:"tags"`
	Category string   `json:"category"`
	Content  string   `json:"content"`
	Title    string   `json:"title"`
	Key      string   `json:"key"`
	ParentId int      `json:"parentId"`
}

func SaveNewKnowledge(c *gin.Context) {
	var err error
	var ks = &KnowStrut{}
	if err = c.ShouldBindJSON(ks); err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": err.Error()})
		return
	}

	if ks.Title == "" || ks.Content == "" {
		zap.S().Errorf("param is null")
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": "param is null"})
		return
	}
	token := c.Request.Header.Get("token")
	username, err := Gjwt.GetUserNameByToken(token)
	if err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": err.Error()})
		return
	}

	know, cate, err := knowledgeService.SaveNewKnowledge(ks.ParentId, ks.Tags, username, ks.ParentId, ks.Title, ks.Content)
	if err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": err.Error()})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "成功", "data": gin.H{"knowledge": know, "category": cate}})
	}
}

func SaveEditKnowledge(c *gin.Context) {
	var err error
	var ks = &KnowStrut{}
	if err = c.ShouldBindJSON(ks); err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": err.Error()})
		return
	}

	if ks.Title == "" || ks.Content == "" {
		zap.S().Errorf("%s", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": err.Error()})
		return
	}

	if err = knowledgeService.UpdateCatelogTItle(ks.Title, ks.Key); err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": err.Error()})
		return
	}

	cate, err := knowledgeService.SaveEditKnowledge(ks.Tags, ks.Title, ks.Content, ks.Key)
	if err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": err.Error()})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "保存成功.", "data": gin.H{"category": cate}})
		return
	}
}

func BatchCancelShare(c *gin.Context) {
	var err error
	type IList struct {
		IdList string `json:"idList"`
	}
	IL := &IList{}
	if err = c.ShouldBindJSON(IL); err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": err.Error()})
		return
	}

	Idlists := strings.Split(IL.IdList, ",")

	for _, v := range Idlists {
		if err := db2.Orm.Where("id=?", v).Delete(&model.Knowledge{}).Error; err != nil {
			zap.S().Errorf("%s", err.Error())
			c.JSON(http.StatusOK, gin.H{"code": 400, "msg": err.Error()})
			return
		}

	}

	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "成功"})

}

func CancelShareByKey(c *gin.Context) {
	id := c.Query("key")
	if err := db2.Orm.Model(&model.Knowledge{}).Where("ckey=?", id).Update("shared", false).Error; err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "成功"})

}

func CancelShare(c *gin.Context) {
	id := c.Query("id")

	if err := db2.Orm.Where("id=?", id).Delete(&model.Knowledge{}).Error; err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "成功"})

}

func GetShareknowledgeInBackend(c *gin.Context) {
	var err error
	page := &request.PageParam{}
	if err = c.ShouldBindJSON(page); err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": err.Error()})
		return
	}

	data, total, err := knowledgeService.GetKnowlegeInBackend(page)
	if err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": "failed", "data": ""})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "成功", "data": data, "total": int(total)})
}

func SearchShareKnowledgeByWord(c *gin.Context) {
	page := &request.PageParam{}
	if err := c.ShouldBindJSON(page); err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": err.Error()})
		return
	}

	dataListString, total, err := knowledgeService.SearchShareKnowledgeByWord(page)
	if err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "成功", "data": dataListString, "total": int(total)})
}

func GetHotList(c *gin.Context) {
	data, err := knowledgeService.GetHotList()
	if err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "成功", "data": data})
}

func Shareknowledge(c *gin.Context) {
	key := c.Query("key")
	shareUrl, err := knowledgeService.Shareknowledge(key)
	if err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "成功", "data": shareUrl})

}

func GetKnowledgeCategories(c *gin.Context) {
	key := c.Query("key")

	KnowledgeCategories, err := knowledgeService.GetKnowledgeCategoriesList(key)
	if err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "成功", "data": KnowledgeCategories})

}

type Res struct {
	Code  int    `json:"code"`
	Msg   string `json:"msg"`
	Data  string `json:"data,omitempty"`
	Total int    `json:"total,omitempty"`
}

func GetKnowledgeByKey(c *gin.Context) {
	var err error
	key := c.Query("key")

	Knowledge, err := knowledgeService.GetKnowledgeByKey(key)
	if err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "成功", "data": Knowledge})

}

func GetSummary(c *gin.Context) {

	knowTitleList, err := knowledgeService.GetSummary()
	if err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "成功", "data": knowTitleList})

}

func AddNode(c *gin.Context) {
	nodeName := c.Query("nodeName")
	parentId := c.Query("parentId")
	if nodeName == "" || parentId == "" {
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": "参数错误"})
		return
	}
	pid, _ := strconv.Atoi(parentId)
	cate, err := knowledgeService.AddNode(pid, nodeName)
	if err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "成功", "data": cate})
}

func GetTree(c *gin.Context) {

	tree, err := knowledgeService.GetTree()
	if err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "成功", "data": tree})

}

func DelTreeNode(c *gin.Context) {
	var err error
	type delOption struct {
		IsLeaf bool `json:"isLeaf"`
		Id     int  `json:"id"`
	}

	var op = &delOption{}

	if err = c.ShouldBindJSON(op); err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": err.Error()})
		return
	}
	err = knowledgeService.DelTreeNode(op.IsLeaf, op.Id)
	if err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "成功"})

}

func ModifyTreeNode(c *gin.Context) {
	var err error
	type delOption struct {
		IsLeaf  bool   `json:"isLeaf"`
		Id      int    `json:"id"`
		Keyword string `json:"keyword"`
	}

	var op = &delOption{}

	if err = c.ShouldBindJSON(op); err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": err.Error()})
		return
	}
	err = knowledgeService.ModifyTreeNode(op.IsLeaf, op.Id, op.Keyword)
	if err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "成功"})

}

func UploadImage(c *gin.Context) {

	imageBaseDir := "upload/img/"
	f, err := c.FormFile("mkimg")
	if err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": err.Error()})
		return
	}

	ext := strings.Split(f.Filename, ".")[len(strings.Split(f.Filename, "."))-1]
	if strings.Contains(ext, ".jpg") || strings.Contains(ext, ".png") {
		zap.S().Errorf("fileOp type error")
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": "fileOp type error"})
		return
	}
	filename := helper2.Md5(f.Filename+time.Now().String()) + "." + ext
	path := imageBaseDir + filename
	//关闭上传的文件，不然的话会出现临时文件不能清除的情况
	err = c.SaveUploadedFile(f, path)
	if err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": err.Error()})
		return
	}
	rpath := "/img/" + filename
	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "成功", "data": rpath})
}

func DelKnowledgeByKey(c *gin.Context) {

	tree, err := knowledgeService.GetTree()
	if err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "success", "data": tree})

}

// func DelTag(c *gin.Context) {
// 	id := c.Query("id")
// 	if id == "" {
// 		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": "param is null"})
// 		return
// 	}
// 	iid, err := strconv.Atoi(id)
// 	if err != nil {
// 		zap.S().Errorf("%s", err.Error())
// 		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": err.Error()})
// 		return
// 	}
// 	err = knowledgeService.DelTag(iid)
// 	if err != nil {
// 		zap.S().Errorf("%s", err.Error())
// 		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "success"})

// }

// func AddTag(c *gin.Context) {
// 	key := c.Query("key")
// 	if key == "" {
// 		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": "param is null"})
// 		return
// 	}
// 	err := knowledgeService.AddTag(key)
// 	if err != nil {
// 		zap.S().Errorf("%s", err.Error())
// 		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "success"})

// }
