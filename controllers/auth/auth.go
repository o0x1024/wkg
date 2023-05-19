package auth

import (
	"net/http"
	"wkg/model"
	response2 "wkg/model/response"
	db2 "wkg/pkg/db"
	Gjwt2 "wkg/pkg/libs/jwt"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/mailru/easyjson/jlexer"
)

type Result struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data string `json:"data"`
}

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (u User) UnmarshalEasyJSON(w *jlexer.Lexer) {
	panic("implement me")
}

func ModyfiPasswd(c *gin.Context) {
	type MUser struct {
		OldPasswd       string `json:"oldPasswd"`
		NewPasswd       string `json:"newPasswd"`
		ConfimNewPasswd string `json:"confimNewPasswd"`
	}

	mu := &MUser{}
	if err := c.ShouldBindJSON(mu); err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": "ShouldBindJSON error"})
		return
	}

	if mu.NewPasswd != mu.ConfimNewPasswd {
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": "password not same"})
		return
	}

	var count int64
	if err := db2.Orm.Model(&model.Users{}).Where("password=? and id=1", mu.OldPasswd).Count(&count).Error; err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": "password error"})
		return
	}

	u := &User{}
	u.Password = mu.NewPasswd

	if count > 0 {
		if err := db2.Orm.Model(&User{}).Where("id=1").Update("password", mu.NewPasswd).Error; err != nil {
			zap.S().Errorf("%s", err.Error())
			c.JSON(http.StatusOK, gin.H{"code": 400, "msg": "db error"})
			return
		} else {
			c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "成功"})
			return
		}

	} else {
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": "old password error"})
		return
	}

}

func Login(c *gin.Context) {

	user := &User{}
	err := c.ShouldBindJSON(user)
	if err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": "ShouldBindJSON error"})
		return
	}

	var count int64
	err = db2.Orm.Model(&model.Users{}).Where("username=? and password=?", user.Username, user.Password).Count(&count).Error
	if err != nil {
		zap.S().Errorf("%s", err.Error())
		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": "query error"})
		return
	}

	if count > 0 {
		token := Gjwt2.GenToken()
		rd := response2.Rdata{Token: token}
		c.JSON(http.StatusOK, gin.H{"code": 200, "msg": "成功", "data": rd})
		return
	} else {

		c.JSON(http.StatusOK, gin.H{"code": 400, "msg": "login failed"})
		return
	}
}
