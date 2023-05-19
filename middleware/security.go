package middleware

import (
	"net/http"
	"strings"
	"sync"
	"time"
	"wkg/global"
	"wkg/model"
	db2 "wkg/pkg/db"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var (
	once  = sync.Once{}
	rlock = sync.RWMutex{}
)

func BruteBlock() gin.HandlerFunc {
	updateBlackList()
	return func(c *gin.Context) {
		ip := strings.Split(c.Request.RemoteAddr, ":")[0]
		if strings.Contains(c.Request.RequestURI, "login") && ip != "" {
			if count, ok := global.IPReqRecord[ip]; ok {
				if count > 10 {
					rlock.Lock()
					global.IPReqRecord[ip] = count + 1
					zap.S().Info("有人爆破密码,IP:" + ip)
					rlock.Unlock()
					c.JSON(http.StatusOK, gin.H{"code": 500, "msg": "就这."})
					c.Abort()
				} else {
					rlock.Lock()
					global.IPReqRecord[ip] = count + 1
					rlock.Unlock()
				}
			} else {
				rlock.Lock()
				global.IPReqRecord[ip] = count + 1
				rlock.Unlock()
			}
		}

		c.Next()
		return
	}
}

func updateBlackList() {
	go once.Do(func() {
		tick := time.NewTicker(1 * time.Hour)
		for {
			select {
			case <-tick.C:
				//存储一下攻击IP
				for k, v := range global.IPReqRecord {
					bip := &model.BlackIpList{}
					bip.Ip = k
					bip.Count = v
					if err := db2.Orm.Model(&model.BlackIpList{}).Create(&bip).Error; err != nil {
						zap.S().Errorf("%s", err.Error())
						continue
					}
				}
				global.IPReqRecord = make(map[string]int)
				// zap.S().Info("%s", "更新黑名单IP缓存信息")
			}
		}
	})
}
