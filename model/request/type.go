package request


type PageParam struct {
	Page int 		`json:"page,omitempty"`
	PageSize int 	`json:"pagesize,omitempty"`
	Type 	string	`json:"type,omitempty"`
	KeyWord	string	`json:"keyword,omitempty"`
}

type SearchStrut struct {
	Type string `json:"type"`
	KeyWord string `json:"keyWord"`
}