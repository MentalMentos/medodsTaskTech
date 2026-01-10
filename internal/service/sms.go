package service

import (
	"errors"
	"fmt"

	"github.com/MentalMentos/medodsTaskTech/internal/config"
	"github.com/MentalMentos/medodsTaskTech/internal/data/request"
	"github.com/MentalMentos/medodsTaskTech/internal/model"
	"github.com/MentalMentos/medodsTaskTech/internal/repository/sms"
)

func (s *AuthService) SendNotify(req request.SendNotificationReq) error {
	switch model.SendType(req.Type) {
	case model.SmsType:
		if isSend, id, err := s.sendSMS(req.Phone, req.Data); err != nil || !isSend {
			if !isSend {
				s.logger.Fatal(
					"Message not sent",
					fmt.Sprintf("id: %v, phone: %s, data: %s", id, req.Phone, req.Data),
				)
			}
			return err
		}
		return nil
	case model.EmailType:
		return nil
	default:
		return errors.New("invalid type")
	}
}

func (s *AuthService) sendSMS(phone, message string) (isSend bool, id string, err error) {
	// Получаем конфигурацию для отправки СМС
	config, err := config.GetSmsConfig()

	// Если в конфиге намеренно указанно, что смс не нужно отправлять, то выходим без отправки
	if config.User == "disable" {
		return false, id, errors.New("SmsDisableError")
	}

	if err != nil {
		return false, id, err
	}

	// Отправляем сообщение
	beeSms := sms.NewBeeSms(s.logger, config.User, config.Pass, config.Sender)
	id, err = beeSms.SendSms(phone, message)
	if err == nil {
		isSend = true
	}
	return
}
