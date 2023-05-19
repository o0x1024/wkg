package response

type Rdata struct {
	Token string `json:"token"`
}
type Result struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data Rdata  `json:"data"`
}

type BugFindRes struct {
	Code int      `json:"code"`
	Msg  string   `json:"msg"`
	Data []string `json:"data"`
}

type Option struct {
	Value string `json:"value"`
	Lable string `json:"label"`
}
