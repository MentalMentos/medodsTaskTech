package request

import "github.com/MentalMentos/medodsTaskTech/internal/data/response"

// -----------------------
type Request struct {
	Meta    *RequestMeta    `json:"meta"`
	Request *RequestRequest `json:"request"`
	Session *RequestSession `json:"session"`
	State   *RequestState   `json:"state"`
	Version string          `json:"version"`
}

type RequestMeta struct {
	Locale     string                `json:"locale"`
	Timezone   string                `json:"timezone"`
	ClientId   string                `json:"client_id"`
	Interfaces RequestMetaInterfaces `json:"interfaces"`
}

type RequestMetaInterfaces struct {
	Screen         struct{} `json:"screen"`
	AccountLinking struct{} `json:"account_linking"`
	AudioPlayer    struct{} `json:"audio_player"`
}

type RequestSession struct {
	MessageId   int                       `json:"message_id"`
	SessionId   string                    `json:"session_id"`
	SkillId     string                    `json:"skill_id"`
	UserId      string                    `json:"user_id"`
	User        RequestSessionUser        `json:"user"`
	Application RequestSessionApplication `json:"application"`
	New         bool                      `json:"new"`
}
type RequestState struct {
	Session     response.IntValue `json:"session"`
	User        response.IntValue `json:"user"`
	Application response.IntValue `json:"application"`
}
type RequestSessionUser struct {
	UserId      string `json:"user_id"`
	AccessToken string `json:"access_token"`
}
type RequestSessionApplication struct {
	ApplicationId string `json:"application_id"`
}

type RequestRequest struct {
	Command           string               `json:"command"`
	OriginalUtterance string               `json:"original_utterance"`
	Markup            RequestRequestMarkup `json:"markup"`
	Payload           struct{}             `json:"payload"`
	Nlu               *RequestRequestNlu   `json:"nlu"`
	Type              string               `json:"type"`
}

type RequestRequestMarkup struct {
	DangerousContext bool `json:"dangerous_context"`
}

type RequestRequestNlu struct {
	Tokens   []string                     `json:"tokens"`
	Entities []*RequestRequestNluEntities `json:"entities"`
	Intents  struct{}                     `json:"intents"`
}

type RequestRequestNluEntities struct {
	Tokens *RequestRequestNluEntitiesToken `json:"tokens"`
	Type   string                          `json:"type"`
	Value  interface{}                     `json:"value"`
}

type RequestRequestNluEntitiesToken struct {
	Start int `json:"start"`
	End   int `json:"end"`
}

// -----------------------
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
