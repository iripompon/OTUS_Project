package model

import "encoding/xml"

// Идентификатор абонента. Для Страхователя - регномер, для МО - ОГРН
type RecipientId string

const (
	RecipientIdGPI RecipientId = "7713050625"
	MessageType86  string      = "86"
)

type SystemInfo struct {

	// Версия спецификации, по которой происходит взаимодействие

	SpecVersion string `xml:"systemInfo>specVersion"`

	// Наименование используемого программного обеспечения

	Software string `xml:"systemInfo>software"`

	// Версия используемого программного обеспечения

	SoftwareVersion string `xml:"systemInfo>softwareVersion"`
}

type PutMessageRequest struct {
	XMLName xml.Name `xml:"n0:putMessageRequest"`

	NameSpaceN0 string `xml:"xmlns:n0,attr"`

	// Информация о системе, направляющей запрос
	SystemInfo

	// Тип взаимодействия
	InteractionType int32 `xml:"interactionType"`

	RecipientId RecipientId `xml:"recipientId"`

	// Идентификатор типа сообщения
	MessageType string `xml:"messageType"`

	// Сообщение в кодировке base64
	Message string `xml:"message"`
}

func NewPutMessage(Version string, Software string, SoftwareVersion string, InterType int32) *PutMessageRequest {
	return &PutMessageRequest{
		SystemInfo: SystemInfo{
			SpecVersion:     Version,
			Software:        Software,
			SoftwareVersion: SoftwareVersion,
		},
		InteractionType: InterType,
		RecipientId:     RecipientIdGPI,
		MessageType:     MessageType86,
	}
}

// здесь еще добавить структуру для GetMessage и GenCntMessage

// type CommonMessageRequest struct {

// 	// Информация о системе, направляющей запрос
// 	SystemInfo *SystemInfo `xml:"systemInfo,omitempty" json:"systemInfo,omitempty"`

// 	// Тип взаимодействия
// 	InteractionType int32 `xml:"interactionType,omitempty" json:"interactionType,omitempty"`
// }

// type GetCntMessageRequest struct {
// 	XMLName xml.Name `xml:"http://www.fss.ru/integration/types/sedo/v01 getCntMessageRequest"`

// 	*CommonMessageRequest

// 	// Дата сообщения
// 	MessageDate soap.XSDDate `xml:"messageDate,omitempty" json:"messageDate,omitempty"`

// 	// Дата сообщения (за период ОТ)
// 	MessageFrom soap.XSDTime `xml:"messageFrom,omitempty" json:"messageFrom,omitempty"`

// 	// Дата сообщения (за период ДО)
// 	MessageTo soap.XSDTime `xml:"messageTo,omitempty" json:"messageTo,omitempty"`
// }

// type GetMessageStatusRequest struct {
// 	*CommonMessageRequest

// 	// Идентификатор абонента
// 	RecipientId *RecipientId `xml:"recipientId,omitempty" json:"recipientId,omitempty"`

// 	// Идентификатор сообщения
// 	Uuids *UuidList `xml:"uuids,omitempty" json:"uuids,omitempty"`
// }
