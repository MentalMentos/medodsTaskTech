package model

type YaUrl string

// GetDevices GET
// Body: Authorization
const GetDevices YaUrl = "https://example.com/v1.0/user/devices"

//

func (y *YaUrl) GetDevices() string {
	return string(GetDevices)
}
