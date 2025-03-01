package model

import (
	"encoding/xml"
)

type InsuredPersonStatusType string

const (

	// Статус говорит о том, что началась обработка сообщения
	InsuredPersonStatusTypeRECEIVED InsuredPersonStatusType = "RECEIVED"

	// Статус говорит о том, что в процессе обработки обнаружены ошибки
	InsuredPersonStatusTypeERROR InsuredPersonStatusType = "ERROR"

	// Сообщение успешно обработано
	InsuredPersonStatusTypePROCESSED InsuredPersonStatusType = "PROCESSED"
)

type GenderType string

const (

	// Пол: мужской
	GenderTypeMALE GenderType = "MALE"

	// Пол: женский
	GenderTypeFEMALE GenderType = "FEMALE"
)

type Protocol struct {
	Item []struct {

		// Тип записи протокола
		Mnemonic string `xml:"mnemonic,omitempty"`

		// Уровень записи протокола
		Level string `xml:"level,omitempty"`

		// Текст записи протокола
		Text string `xml:"text,omitempty"`

		// Дополнительное описание
		Description string `xml:"description,omitempty"`
	} `xml:"item,omitempty"`
}

type InsuredPersonResponseType struct {

	// СНИЛС застрахованного информация о котором подавалась ранее
	Snils string `xml:"snils,omitempty"`

	// Статус обработки сообщения карточки застрахованного
	Status InsuredPersonStatusType `xml:"status,omitempty"`

	// Протокол несоответствия
	Protocol `xml:"protocol,omitempty"`
}

// структура 87 сообщения
type InsuredPersonResponse struct {
	XMLName xml.Name `xml:"urn:ru:fss:integration:types:rpu:InsuredPerson:v01 insuredPersonResponse"`

	InsuredPersonResponseType

	ResponseOn string `xml:"responseOn,attr,omitempty"`
}

type MismatchProtocol struct {
	Item []struct {
		Mnemonic string `xml:"mnemonic,omitempty"`

		Message string `xml:"message,omitempty"`

		Description string `xml:"description,omitempty"`
	} `xml:"item,omitempty"`
}

// структура 88 сообщения
type InsuredPersonMismatchType struct {
	XMLName xml.Name `xml:"urn:ru:fss:integration:types:rpu:InsuredPerson:v01 mismatchInfo"`

	// СНИЛС застрахованного информация о котором подавалась ранее
	Snils string `xml:"snils,omitempty"`

	// Протокол несоответствия
	MismatchProtocol `xml:"protocol,omitempty"`
}

type FullName struct {

	// Имя
	FirstName string `xml:"fullName>firstName"`

	// Фамилия
	LastName string `xml:"fullName>lastName"`

	// Отчество
	MiddleName string `xml:"fullName>middleName"`
}

type PassportRF struct {

	// Серия паспорта РФ

	Series string `xml:"identityDocument>passport>series"`

	// Номер паспорта РФ

	Number string `xml:"identityDocument>passport>number"`

	// Дата выдачи
	IssueDate string `xml:"identityDocument>passport>issueDate"`

	// Кем выдан
	WhoIssued string `xml:"identityDocument>passport>whoIssued"`
}

type IdentityDocument struct {
	PassportRF
}

type AddressFias struct {

	// Квартира
	Flat int `xml:"regAddress>fiasAddress>flat"`

	// Соответствие ГАР: Если указан дом: HOUSEGUID – Глобальный уникальный идентификатор дома

	HouseGuid string `xml:"regAddress>fiasAddress>houseGuid"`
}

type RegAddress struct {
	AddressFias

	PostalCode string `xml:"postalCode,omitempty"`
}

type BankInfo struct {

	// Наименование банка

	BankName string `xml:"methodReceivePayment>bankInfo>bankName"`

	// БИК банка

	Bik string `xml:"methodReceivePayment>bankInfo>bik"`

	// Номер расчетного счета

	AccountNum string `xml:"methodReceivePayment>bankInfo>accountNum"`
}

type MethodReceivePayment struct {
	BankInfo
}

// структура 86 сообщения
type InsuredPerson struct {
	XMLName xml.Name `xml:"urn:ru:fss:integration:types:rpu:InsuredPerson:v01 insuredPerson"`

	FullName

	BirthDate string `xml:"birthDate"`

	Gender GenderType `xml:"gender"`

	Snils string `xml:"snils"`

	Inn string `xml:"inn"`

	IdentityDocument

	RegAddress

	MethodReceivePayment
}

func NewInsuredPerson(snils string, inn string) *InsuredPerson {
	return &InsuredPerson{
		Snils: snils,
		Inn:   inn,
	}
}
