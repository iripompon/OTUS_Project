package model

//message 86 object InsuredPerson

type FullName struct {
	FirstName  string
	LastName   string
	MiddleName string
}

type Fired struct {
	Date               string
	ContractCancelDate string
}

type Passport struct {
	Series    string
	Number    string
	IssueDate string
	WhoIssued string
}

type TempIdentDocument struct {
	Number         string
	IssueDate      string
	WhoIssued      string
	ExpirationDate string
}

type OtherIdentDocType struct {
	Type           string
	Series         string
	Number         string
	ExpirationDate string
}

type IdentityDocument struct {
	Passport
	TempIdentDocument
	OtherIdentDocType
}
