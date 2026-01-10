package sms

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/MentalMentos/medodsTaskTech/internal/repository/models"
	"github.com/MentalMentos/medodsTaskTech/pkg/helpers/phone"
	"github.com/MentalMentos/medodsTaskTech/pkg/logger"
)

const (
	HOST   = "https://a2p-sms-https.beeline.ru/proto/http/rest"
	ACTION = "post_sms"
)

type BeeSms struct {
	user    string
	pass    string
	action  string
	message string
	sender  string
	phone   string
	client  *http.Client
	logger  logger.Logger
}

func NewBeeSms(logger logger.Logger, user, pass, sender string) *BeeSms {
	return &BeeSms{
		logger: logger,
		action: ACTION,
		user:   user,
		pass:   pass,
		sender: sender,
		client: &http.Client{
			Timeout: time.Second * 8,
		},
	}
}

func (b *BeeSms) SendSms(tel, message string) (id string, err error) {
	phoneLib, _ := phone.NewPhone(tel)
	if !phoneLib.IsMobile() {
		return id, errors.New("PhoneNotMobileError")
	}
	b.message = strings.Trim(message, " ")
	b.phone = phoneLib.E164()
	output, err := b.request()
	if err != nil {
		return "", err
	}

	if len(output.Actions) == 0 {
		return id, errors.New("SendMessageFailedError")
	}

	id = "sms_" + output.Actions[0].Id
	return
}

func (b *BeeSms) request() (output models.Output, err error) {
	data := models.SendSMSReq{
		User:    b.user,
		Pass:    b.pass,
		Action:  b.action,
		Sender:  b.sender,
		Message: b.message,
		Target:  b.phone,
	}

	marshal, err := json.Marshal(data)
	if err != nil {
		return output, err
	}
	reqBody := bytes.NewBuffer(marshal)

	resp, err := b.client.Post(HOST, "application/json", reqBody)
	if err != nil {
		return output, err
	}
	defer resp.Body.Close()

	byt, err := io.ReadAll(resp.Body)
	if err != nil {
		return output, err
	}
	err = json.Unmarshal(byt, &output)

	var body []byte
	if output.Error.Message != "" {
		body, err = json.Marshal(map[string]interface{}{
			"phone": b.phone,
			"data":  b.message,
		})
		if err != nil {
			return output, err
		}
		b.logger.Info(fmt.Sprintf(
			"Ответ из Биллайн: %s",
			output.Error.Message),
			string(body),
		)

		return output, fmt.Errorf("отправка СМС: %s", output.Error.Message)
	}
	return
}
