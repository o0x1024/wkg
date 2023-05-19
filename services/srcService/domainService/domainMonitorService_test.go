package domainService

import (
	"log"
	"testing"
	"wkg/model"
	db2 "wkg/pkg/db"
)

func TestDomainMoniter_UpdateDomainInfo(t *testing.T) {
	var DomainCache []model.Domain

	err := db2.Orm.Model(&model.Domain{}).Find(&DomainCache).Error
	if err != nil {
		log.Println("[!] domainService.go sql domain cache failed. line:83,   [", err, "]")
		return
	}

	err = db2.Orm.Model(&DomainCache).Where("1=1").Save(&DomainCache).Error
	if err != nil {
		log.Println("[!] domainService.go save domain cache failed. line:75.  [", err, "]")
		return
	}
}

func addrtram(str *string) {
	*str = "abc"
}

func TestDetDomaininfo(t *testing.T) {

}
