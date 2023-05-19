package portService

import (
	"context"
	"fmt"
	"strings"
	"wkg/model"
	db2 "wkg/pkg/db"
	helper2 "wkg/pkg/libs/helper"

	"github.com/Ullaakut/nmap/v2"
	"go.uber.org/zap"
)

func PortScan(cmp model.Company) error {
	Tbool := true
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	ipsList := []model.Ips{}
	if err := db2.Orm.Model(&model.Ips{}).Where("cid=?", cmp.Id).Find(&ipsList).Error; err != nil {
		return err
	}
	scanner, err := nmap.NewScanner(
		nmap.WithPorts("80,81,82,83,84,85,86,87,88,89,8000,8001,8002,8007,8008,8009,8010,8011,8021,8022,8031,8042,8045,8080,8081,8082,8083,8084,8085,8086,8087,8088,8089,8090,8093,8099,8100,8180,8181,8192,8193,8194,8200,8222,8254,8290,8291,8292,8300,8333,8383,8400,8402,8443,8500,8600,8649,8651,8652,8654,8701,8800,8873,8888,8899,8994,9000,9001,9002,9003,9009,9010,9011,9040,9050,9071,9080,9081,9090,9091,9099,9100,9101,9102,9103,9110,9111,9200,1099,873,5984,6379,7001,7002,9300,11211,27017,27018,50000,50060,50070,50030,2375,3128,2601,2604,4440,4848,9000,9043,21,22,23,161,389,445,1433,1521,3306,3389,5432,5900"),
		nmap.WithContext(ctx),
	)
	if err != nil {
		zap.S().Errorf("%s", err.Error())
		return err
	}
	for _, v := range ipsList {
		scanner.AddOptions(nmap.WithTargets(v.Ip))
	}

	result, _, err := scanner.Run()
	if err != nil {
		zap.S().Errorf("%s", err.Error())
		return err
	}

	for _, host := range result.Hosts {
		if len(host.Ports) == 0 || len(host.Addresses) == 0 {
			continue
		}
		if len(host.Ports) > 15 {
			continue
		}
		//fmt.Printf("Host %q:\n", host.Addresses[0])

		for _, port := range host.Ports {
			//fmt.Printf("\tPort %d/%s %s %s\n", port.ID, port.Protocol, port.State, port.Service.Name)
			if strings.Contains(port.State.State, "open") {
				service := &model.Services{}
				service.Cid = cmp.Id
				service.IsNew = &Tbool
				service.Service = port.Service.Name
				service.Host = host.Addresses[0].Addr
				service.Port = fmt.Sprintf("%d", port.ID)
				service.UpdateTime = helper2.GetCurTime()
				var count int64
				if err := db2.Orm.Model(&model.Services{}).Where("host=? and port=?", host.Addresses[0].Addr, port.ID).Count(&count).Error; err != nil {
					zap.S().Errorf("%s", err.Error())
					continue
				}
				if count == 0 {
					if err := db2.Orm.Model(&model.Services{}).Create(&service).Error; err != nil {
						zap.S().Errorf("%s", err.Error())
						continue
					}
				}
			}
		}

	}

	return nil
}
