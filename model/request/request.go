package request

type Query struct {
	Page        int    `json:"page"`
	PageSize    int    `json:"pagesize"`
	Type        string `json:"type"`
	Keyword     string `json:"keyword"`
	Cid         int    `json:"cid"`
	Id          int    `json:"id"`
	AssetOption string `json:"assetOption"`
}

type ParamIds struct {
	Id string `json:"id"`
}

type InputDataStruct struct {
	CompanyId int      `json:"companyId"`
	DataType  string   `json:"dataType"`
	Content   []string `json:"content"`
}

type ChatGPT struct {
	ApiKey    string `json:"apiKey"`
	SecretKey string `json:"secretKey"`
	Question  string `json:"question"`
}
