package model

//message 87 object InsuredPersonResponse

type Protocol struct {
	item []string
}

type InsuredPersonResponse struct {
	ResponseOn string
	Snils      string
	Status     string
	Protocol
}
