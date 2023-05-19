package request

type AgentRegisterReq struct {
	AgentId string `json:"agentId"`
	Name    string `json:"name,omitempty"`
	Version string `json:"version,omitempty"`

	PrivateIPv4 string `json:"privateIPv4,omitempty"`
	PublicIPv4  string `json:"publicIPv4,omitempty"`
	PrivateIPv6 string `json:"privateIPv6,omitempty"`
	PublicIPv6  string `json:"publicIPv6,omitempty"`

	ProjectName string `json:"projectName,omitempty"`
	Hostname    string `json:"hostname,omitempty"`
	Network     string `json:"network,omitempty"`
	ServerPath  string `json:"serverPath,omitempty"`
	ServerEnv   string `json:"serverEnv,omitempty"`
	Pid         string `json:"pid,omitempty"`
}

type Ids struct {
	Id string `json:"id"`
}

type ListData struct {
	Href        string `json:"href"`
	Title       string `json:"title"`
	Hit         int    `json:"hit"`
	Description string `json:"description"`
}

// UploadReq
/* Type
数据类型，可选择：
1 - Agent心跳数据
2 - 攻击信息
3 - 异常信息
*/
type UploadReq struct {
	Type   int    `json:"type,omitempty"`
	Detail Detail `json:"detail,omitempty"`
}

type Detail struct {
	AgentId string `json:"agentId"`
	Pant
}

type Pant struct {
	Disk   string `json:"disk,omitempty"`
	Memory string `json:"memory,omitempty"`
	Cpu    string `json:"cpu,omitempty"`
}

type SearchParam struct {
	Page     int    `json:"page"`
	PageSize int    `json:"pageSize"`
	Type     string `json:"type"`
	Keyword  string `json:"keyword"`
}

/*
攻击类型
1

*/

type DomainScanRes struct {
	Cid     int    `json:"cid"`
	Domains string `json:"domains"`
}
