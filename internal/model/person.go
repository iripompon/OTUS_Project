package model

import (
	"encoding/base64"
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
	FirstName string `xml:"n0:fullName>n1:firstName"`

	// Фамилия
	LastName string `xml:"n0:fullName>n1:lastName"`

	// Отчество
	MiddleName string `xml:"n0:fullName>n1:middleName"`
}

type PassportRF struct {

	// Серия паспорта РФ

	Series string `xml:"n0:identityDocument>n0:passport>n0:series"`

	// Номер паспорта РФ

	Number string `xml:"n0:identityDocument>n0:passport>n0:number"`

	// Дата выдачи
	IssueDate string `xml:"n0:identityDocument>n0:passport>n0:issueDate"`

	// Кем выдан
	WhoIssued string `xml:"n0:identityDocument>n0:passport>n0:whoIssued"`
}

type IdentityDocument struct {
	PassportRF
}

type AddressFias struct {

	// Квартира
	Flat int `xml:"n0:regAddress>n0:fiasAddress>n2:flat"`

	// Соответствие ГАР: Если указан дом: HOUSEGUID – Глобальный уникальный идентификатор дома

	HouseGuid string `xml:"n0:regAddress>n0:fiasAddress>n2:houseGuid"`
}

type RegAddress struct {
	AddressFias

	PostalCode string `xml:"n0:regAddress>n0:postalCode"`
}

type BankInfo struct {

	// Наименование банка

	BankName string `xml:"n0:methodReceivePayment>n2:bankInfo>n2:bankName"`

	// БИК банка

	Bik string `xml:"n0:methodReceivePayment>n2:bankInfo>n2:bik"`

	// Номер расчетного счета

	AccountNum string `xml:"n0:methodReceivePayment>n2:bankInfo>n2:accountNum"`
}

type MethodReceivePayment struct {
	BankInfo
}

// структура 86 сообщения
type InsuredPerson struct {
	XMLName xml.Name `xml:"n0:insuredPerson"`

	NameSpaceN0 string `xml:"xmlns:n0,attr"`
	NameSpaceN1 string `xml:"xmlns:n1,attr"`
	NameSpaceN2 string `xml:"xmlns:n2,attr"`

	FullName

	BirthDate string `xml:"n0:birthDate"`

	Gender GenderType `xml:"n0:gender"`

	Snils string `xml:"n0:snils"`

	Inn string `xml:"n0:inn"`

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

func MessageToBase64(msg []byte) string {
	return base64.StdEncoding.EncodeToString(msg)
}
