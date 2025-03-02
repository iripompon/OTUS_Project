package model

type SedoEnvelope struct {

	// блок Header – содержит служебную информацию

	SedoHeader

	//блок Body – содержит смысловые данные сообщения

	SedoBody []byte
}
type SedoHeader struct {

	//содержит информацию об электронной подписи сообщения

	Signature

	//содержит публичный сертификат пользователя

	BinarySecurityToken
}

type Signature struct {

	//содержит информацию о методе каноникализации, алгоритме хэширования,
	// алгоритме генерации ЭЦП и ссылку на подписываемый блок данных

	SignedInfo

	//содержит рассчитанное значение ЭП

	SignatureValue string

	//содержит ссылку на сертификат пользователя, который содержится в
	//  BinarySecurityToken и с помощью которого была рассчитана ЭП

	KeyInfo

	//блок для встраивания машиночитаемой доверенности (МЧД)

	Object
}

type SignedInfo struct {
	CanonicalizationMethod string
	SignatureMethod        string
	//Ссылка на подписываемые данные
	Reference
}

type Reference struct {
	Transform string
	//алгоритм вычисления хэш суммы
	DigestMethod string
	//вычисленное значение хэш суммы от подписываемых данных
	DigestValue string
}

type KeyInfo struct {
	SecurityTokenRef string
}

type Object struct {
	PowerOfAttorneyLink string
}

type BinarySecurityToken struct {
	EncodingType string
	ValueType    string
}
