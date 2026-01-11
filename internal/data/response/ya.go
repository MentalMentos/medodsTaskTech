package response

//--------
type ResponseYa struct {
	Response         *ResponseResponse  `json:"response"`
	SessionState     *IntValue          `json:"session_state"`
	UserStateUpdate  *IntValue          `json:"user_state_update"`
	ApplicationState *IntValue          `json:"application_state"`
	Analytics        *ResponseAnalytics `json:"analytics"`
	Version          string             `json:"version"`
}

type ResponseResponse struct {
	Text       string            `json:"text"`
	Tts        string            `json:"tts"`
	Card       *ResponseCard     `json:"card"`
	Buttons    []*ResponseButton `json:"buttons"`
	EndSession bool              `json:"end_session"`
	Directives struct{}          `json:"directives"`
}
type ResponseCard struct {
	Type string `json:"type"`
}

type ResponseButton struct {
	Title   string `json:"title"`
	Payload struct {
	} `json:"payload"`
	Url  string `json:"url"`
	Hide bool   `json:"hide"`
}

type IntValue struct {
	Value int `json:"value"`
}

type ResponseAnalytics struct {
	Events []*ResponseAnalyticsEvent `json:"events"`
}

type ResponseAnalyticsEvent struct {
	Name  string                       `json:"name"`
	Value *ResponseAnalyticsEventValue `json:"value,omitempty"`
}

type ResponseAnalyticsEventValue struct {
	Field       string                                  `json:"field"`
	SecondField *ResponseAnalyticsEventValueSecondField `json:"second field"`
}

type ResponseAnalyticsEventValueSecondField struct {
	ThirdField string `json:"third field"`
}

//--------

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
