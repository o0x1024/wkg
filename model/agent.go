package model

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
)

type Agent struct {
	Records
	AgentId      string `json:"agent_id,omitempty" gorm:"primary_key;column:agent_id"`
	IntranetIpv4 Lis    `json:"IntranetIpv4,omitempty" gorm:"column:IntranetIpv4"`
	IntranetIpv6 Lis    `json:"IntranetIpv6,omitempty" gorm:"column:IntranetIpv6"`
	ExtranetIpv4 Lis    `json:"ExtranetIpv4,omitempty" gorm:"column:ExtranetIpv4"`
	ExtranetIpv6 Lis    `json:"ExtranetIpv6,omitempty" gorm:"column:ExtranetIpv6"`
	Hostname     string `json:"host_name,omitempty" gorm:"column:host_name"`
	Version      string `json:"version,omitempty" gorm:"column:version"`
	Product      string `json:"product,omitempty" gorm:"column:product"`
}

type Records struct {
	KernelVersion   string `json:"kernel_version,omitempty" gorm:"column:kernel_version"`
	Arch            string `json:"arch,omitempty" gorm:"column:arch"`
	Platform        string `json:"platform,omitempty" gorm:"column:platform"`
	PlatformFamily  string `json:"platform_family,omitempty" gorm:"column:platform_family"`
	PlatformVersion string `json:"platform_version,omitempty" gorm:"column:platform_version"`

	TotalMem   string `json:"total_mem,omitempty" gorm:"column:total_mem"`
	Cpu        string `json:"cpu,omitempty" gorm:"column:cpu"`
	Rss        string `json:"rss,omitempty" gorm:"column:rss"`
	ReadSpeed  string `json:"read_speed,omitempty" gorm:"column:read_speed"`
	WriteSpeed string `json:"write_speed,omitempty" gorm:"column:write_speed"`
	Pid        string `json:"pid,omitempty" gorm:"column:pid"`
	Nfd        string `json:"nfd,omitempty" gorm:"column:nfd"`
	StartTime  string `json:"start_time,omitempty" gorm:"column:start_time"`

	HostSerial string `json:"host_serial,omitempty" gorm:"column:host_serial"`
	HostId     string `json:"host_id,omitempty" gorm:"column:host_id"`
	HostModel  string `json:"host_model,omitempty" gorm:"column:host_model"`
	HostVendor string `json:"host_vendor,omitempty" gorm:"column:host_vendor"`

	CpuName    string `json:"cpu_name,omitempty" gorm:"column:cpu_name"`
	BootTime   string `json:"boot_time,omitempty" gorm:"column:boot_time"`
	CpuUsage   string `json:"cpu_usage,omitempty" gorm:"column:cpu_usage"`
	MemUsage   string `json:"mem_usage,omitempty" gorm:"column:mem_usage"`
	UpdateTime string `json:"update_time,omitempty" gorm:"column:update_time"`
}

func (d *Agent) TableName() string {
	return "agent"
}

type Lis []string

func (t *Lis) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}
	return json.Unmarshal(bytes, &t)
}

func (t Lis) Value() (driver.Value, error) {
	str, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return string(str), nil
}
