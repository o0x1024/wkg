package conf

import (
	"io/ioutil"
	"os"
	"wkg/model"

	"go.uber.org/zap"
	"gopkg.in/yaml.v2"
)

var Cfg *Config

type Config struct {
	Redis   model.Redis
	Mysql   model.Mysql
	BaseUrl string `yaml:"base_url"`
}

type Scan struct {
	VulnScanRate    int
	DomainScanRate  int
	WebSiteScanRate int
	IPsScanRate     int
}

func LoadConfig(configFile string) (config *Config, err error) {
	f, err := os.Open(configFile)
	if err != nil {
		zap.S().Error(err)
		return
	}
	defer f.Close()
	data, err := ioutil.ReadAll(f)

	if err != nil {
		zap.S().Error(err)
		return
	}
	config = new(Config)
	err = yaml.Unmarshal(data, config)
	if err != nil {
		zap.S().Error(err)
		return
	}
	//Glog.Println("load config, time:", time.Now().Format("2006-01-02 15:04:05"), "info:", config, "error:", err)
	return
}

func (c *Config) GetRedis() model.Redis {
	return c.Redis
}
