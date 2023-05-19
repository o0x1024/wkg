package Gjwt

import (
	"fmt"
	"reflect"
	"time"

	"github.com/dgrijalva/jwt-go"
	"go.uber.org/zap"
)

var (
	key []byte = []byte("src-monitor-by-gelen")
)

// 产生json web token
func GenToken() string {
	claims := &jwt.StandardClaims{
		NotBefore: int64(time.Now().Unix()),
		ExpiresAt: int64(time.Now().Unix() + 8640),
		Issuer:    "gelen",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(key)
	if err != nil {
		zap.S().Errorf("%s", err.Error())
		return ""
	}
	return ss
}

// 校验token是否有效
func CheckToken(token string) bool {
	curtoken, err := jwt.Parse(token, func(*jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err != nil {
		zap.S().Errorf("%s", err.Error())
		return false
	}
	if curtoken.Valid {
		return true
	}
	return false
}

func GetUserNameByToken(token string) (string, error) {

	curtoken, err := jwt.Parse(token, func(*jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err != nil {
		zap.S().Errorf("%s", err.Error())
		return "", err
	}
	username := getValFormTokenClaims("iss", curtoken.Claims)
	return username, nil
}

func getValFormTokenClaims(key string, claims jwt.Claims) string {

	v := reflect.ValueOf(claims)
	if v.Kind() == reflect.Map {
		for _, k := range v.MapKeys() {
			value := v.MapIndex(k)

			if fmt.Sprintf("%s", k.Interface()) == key {
				return fmt.Sprintf("%v", value.Interface())
			}
		}
	}
	return ""
}
