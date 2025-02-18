package person

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
	FirstName string `xml:"firstName,omitempty"`

	// Фамилия
	LastName string `xml:"lastName,omitempty"`

	// Отчество
	MiddleName string `xml:"middleName,omitempty"`
}

type PassportRF struct {

	// Серия паспорта РФ

	Series string `xml:"series,omitempty"`

	// Номер паспорта РФ

	Number string `xml:"number,omitempty"`

	// Дата выдачи
	IssueDate string `xml:"issueDate,omitempty"`

	// Кем выдан
	WhoIssued string `xml:"whoIssued,omitempty"`
}

type IdentityDocument struct {
	PassportRF `xml:"passport,omitempty"`
}

type AddressFias struct {

	// Квартира
	Flat int `xml:"flat,omitempty"`

	// Соответствие ГАР: Если указан дом: HOUSEGUID – Глобальный уникальный идентификатор дома

	HouseGuid string `xml:"houseGuid,omitempty"`
}

type RegAddress struct {
	AddressFias `xml:"fiasAddress,omitempty"`

	PostalCode string `xml:"postalCode,omitempty"`
}

type BankInfo struct {

	// Наименование банка

	BankName string `xml:"bankName,omitempty"`

	// БИК банка

	Bik string `xml:"bik,omitempty"`

	// Номер расчетного счета

	AccountNum string `xml:"accountNum,omitempty"`
}

type MethodReceivePayment struct {
	BankInfo `xml:"bankInfo,omitempty"`
}

// структура 86 сообщения
type InsuredPerson struct {
	XMLName xml.Name `xml:"urn:ru:fss:integration:types:rpu:InsuredPerson:v01 insuredPerson"`

	FullName `xml:"fullName,omitempty"`

	BirthDate string `xml:"birthDate,omitempty"`

	Gender GenderType `xml:"gender,omitempty"`

	Snils string `xml:"snils,omitempty"`

	Inn string `xml:"inn,omitempty"`

	IdentityDocument `xml:"identityDocument,omitempty"`

	RegAddress `xml:"regAddress,omitempty"`

	MethodReceivePayment `xml:"methodReceivePayment,omitempty"`
}
