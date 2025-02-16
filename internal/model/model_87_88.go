package model

//message 87 object InsuredPersonResponse
//message 88 object InsuredPersonMismatch

type Protocol struct {
	item []string
}

type InsuredPersonResponse struct {
	ResponseOn string
	Snils      string
	Status     string
	Protocol
}

type InsuredPersonMismatch struct {
	Snils string
	Protocol
}
