package models

type Output struct {
	Actions    []Action `json:"actions"`
	AgtId      string   `json:"agt_id"`
	DateReport string   `json:"date_report"`
	Error      Error    `json:"error"`
}

type Action struct {
	Id          string `json:"id"`
	Phone       string `json:"phone"`
	Message     string `json:"message"`
	Action      string `json:"action"`
	SmsType     string `json:"sms_type"`
	SmsResCount string `json:"sms_res_count"`
	SmsGroupId  string `json:"sms_group_id"`
}

type SendSMSReq struct {
	User    string `json:"user"`
	Pass    string `json:"pass"`
	Action  string `json:"action"`
	Sender  string `json:"sender"`
	Message string `json:"message"`
	Target  string `json:"target"`
}

type Error struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}
