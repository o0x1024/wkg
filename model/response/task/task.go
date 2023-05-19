package task

// taskid   1000 信息搜集针对域名和端口扫描    1001 漏洞扫描,漏洞扫描只针对website
type ResTask struct {
	TaskId   string    `json:"taskId,omitempty"`
	TaskNo   int       `json:"taskNo,omitempty"`
	Domains  []Domain  `json:"domains,omitempty"`
	IPs      []IP      `json:"IPs,omitempty"`
	WebSites []WebSite `json:"webSites,omitempty"`
	VSAssets VSAsset   `json:"VSAssets,omitempty"`
}

// 需要扫描的漏洞交道
type VSAsset struct {
	Targets    []string `json:"targets,omitempty"`
	PocContent string   `json:"pocContent,omitempty"`
}

type Domain struct {
	CId         int      `json:"cid,omitempty"`
	RootDomains []string `json:"rootDomains,omitempty"`
}

type IP struct {
	CId int      `json:"cid,omitempty"`
	Ips []string `json:"ips,omitempty"`
}

type WebSite struct {
	CId     int      `json:"cid,omitempty"`
	WebSite []string `json:"webSite,omitempty"`
}
