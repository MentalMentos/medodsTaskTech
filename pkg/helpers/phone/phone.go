package phone

import (
	"github.com/MentalMentos/medodsTaskTech/pkg/helpers/chars"
	"github.com/ttacon/libphonenumber"
)

type Phone struct {
	phoneStr string
	phoneLib *libphonenumber.PhoneNumber
}

func NewPhone(phone string) (*Phone, error) {
	var num *libphonenumber.PhoneNumber

	num, err := libphonenumber.Parse(phone, "RU")
	if err != nil {
		return &Phone{}, err
	}
	return &Phone{
		phoneStr: phone,
		phoneLib: num,
	}, nil
}

func (p *Phone) IsMobile() bool {
	numberType := libphonenumber.GetNumberType(p.phoneLib)
	return numberType == libphonenumber.MOBILE
}

func (p *Phone) IsValid() bool {
	return libphonenumber.IsValidNumber(p.phoneLib)
}

// E164 Формат телефона +79141555900
func (p *Phone) E164() string {
	if p.IsValid() {
		return libphonenumber.Format(p.phoneLib, libphonenumber.E164)
	}
	return ""
}

// E164numbers Формат телефона 79141555900
func (p *Phone) E164numbers() string {
	if p.IsValid() {
		phone := libphonenumber.Format(p.phoneLib, libphonenumber.E164)
		return chars.OnlyNumbers(phone)
	}
	return ""
}

// National Формат телефона 8 (914) 155-59-00
func (p *Phone) National() string {
	if p.IsValid() {
		return libphonenumber.Format(p.phoneLib, libphonenumber.NATIONAL)
	}
	return ""
}

// NationalOnlyNumbers Формат телефона 89141555900
func (p *Phone) NationalOnlyNumbers() string {
	if p.IsValid() {
		phone := libphonenumber.Format(p.phoneLib, libphonenumber.NATIONAL)
		return chars.OnlyNumbers(phone)
	}
	return ""
}

// International Формат телефона +7 914 155-59-00
func (p *Phone) International() string {
	if p.IsValid() {
		return libphonenumber.Format(p.phoneLib, libphonenumber.INTERNATIONAL)
	}
	return ""
}

// RFC3966 Формат телефона tel:+7-914-155-59-00
func (p *Phone) RFC3966() string {
	if p.IsValid() {
		return libphonenumber.Format(p.phoneLib, libphonenumber.RFC3966)
	}
	return ""
}

// WithoutRegion Формат телефона 9141555900
func (p *Phone) WithoutRegion() string {
	if p.IsValid() {
		phone := libphonenumber.Format(p.phoneLib, 5)
		return chars.OnlyNumbers(phone)
	}
	return ""
}
