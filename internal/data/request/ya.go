package request

type YaRequest struct{}

type SendNotificationReq struct {
	Type  string `json:"type"`
	Email string `json:"email"`
	Phone string `json:"phone"`
	Data  string `json:"data"`
}

func (ya *YaRequest) ToSendNotificationReq() SendNotificationReq {
	return SendNotificationReq{}
}
