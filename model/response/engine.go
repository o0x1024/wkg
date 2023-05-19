package response

type ResBase struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
}

type AgentRegisterRes struct {
	ResBase
	Data struct {
		Id            string `json:"id"`
		CoreAutoStart int    `json:"coreAutoStart,omitempty"`
	} `json:"data"`
}

type Rule struct {
	BlackURL             []string `json:"blackURL,omitempty"`
	BlackIP              []string `json:"blackIP,omitempty"`
	InWafStatus          *bool    `json:"inWafStatus,omitempty"`
	RaspBlackStatus      *bool    `json:"raspBlackStatus,omitempty"`
	InWafXssStatus       *bool    `json:"inWafXssStatus,omitempty" `
	RaspSqlStatus        *bool    `json:"raspSqlStatus,omitempty" `
	RaspExecStatus       *bool    `json:"raspExecStatus,omitempty" `
	RaspXssStatus        *bool    `json:"raspXssStatus,omitempty" `
	RaspFileUploadStatus *bool    `json:"raspFileUploadStatus,omitempty" `
	RaspFileReadStatus   *bool    `json:"raspFileReadStatus,omitempty" `
	RaspSSRFStatus       *bool    `json:"raspSSRFStatus,omitempty" `
}
