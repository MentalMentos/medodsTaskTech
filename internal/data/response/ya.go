package response

// GetDevices - получение девайсов
type Device struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

type Payload struct {
	UserID  string   `json:"user_id"`
	Devices []Device `json:"devices"`
}

type YaResponse struct {
	RequestID string  `json:"request_id"`
	Payload   Payload `json:"payload"`
}
