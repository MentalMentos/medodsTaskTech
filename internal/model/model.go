package model

import (
	"time"
)

type User struct {
	ID        int64     `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"type:string" json:"name"`
	Email     string    `gorm:"unique" json:"email"`
	Password  string    `gorm:"type:string" json:"-"`
	Role      string    `gorm:"type:string" json:"role"`
	IP        string    `gorm:"type:string" json:"-"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Task struct {
	Id               int    `json:"id"`
	Type             string `json:"type"`
	DestinationPhone string `json:"destination_phone"`
	Data             string `json:"data"`
}

type SendType string

const SmsType SendType = "sms"

const EmailType SendType = "email"

//func(s SendType) ValidateType() (SendType, error) {
//	switch s {
//	case "sms":
//		return SmsType, nil
//	case "email":
//		return EmailType, nil
//	default:
//		return "", errors.New("unknown type")
//	}
//}
