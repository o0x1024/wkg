package model

type HostInfoType struct {
	Host      string
	Ports     string
	Domain    string
	Url       string
	Timeout   int64
	Scantype  string
	Command   string
	Username  string
	Password  string
	Usernames []string
	Passwords []string
}

type SystemConfigType struct {
	GenConfig bool
	//ThreadNum int
}

type WebDirScanType struct {
	Target         string
	TargetDirPath  string
	Payload        string
	PayloadDirPath string
	BuildInDir     []string
	ThreadNum      int
	Proxy          string
	UserAgent      string
	Timeout        int
}

type WebAliveScanType struct {
	Target  string
	DirPath string
	Proxy   string
	Out     string
}

type EncodeDecodeType struct {
	Target  string
	DirPath string
	Proxy   string
}

type FofaType struct {
	Rule   string
	Doamin bool
	IP     bool
	Title  bool
	Host   bool
	Out    string
}
