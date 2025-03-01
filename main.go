package main

import (
	"encoding/xml"
	"fmt"
	"main/internal/model"
)

func main() {
	insuredperson := model.NewInsuredPerson("02126176506", "231101593917")

	insuredperson.NameSpaceN1 = "http://www.fss.ru/integration/types/person/v02"
	insuredperson.NameSpaceN2 = "http://www.fss.ru/integration/types/common/v01"

	insuredperson.FullName.FirstName = "Светлана"
	insuredperson.FullName.LastName = "Иванова"
	insuredperson.FullName.MiddleName = "Григорьевна"
	insuredperson.BirthDate = "1976-01-18"
	insuredperson.Gender = model.GenderTypeFEMALE

	insuredperson.IdentityDocument.PassportRF.Series = "0321"
	insuredperson.IdentityDocument.PassportRF.Number = "866496"
	insuredperson.IdentityDocument.PassportRF.IssueDate = "2021-05-20"
	insuredperson.IdentityDocument.PassportRF.WhoIssued = "ГУ МВД России по Краснодарскому краю"

	insuredperson.RegAddress.AddressFias.HouseGuid = "9220a25f-d590-4216-9bfa-170a43afa32d"
	insuredperson.RegAddress.AddressFias.Flat = 65
	insuredperson.RegAddress.PostalCode = "350028"

	insuredperson.MethodReceivePayment.BankInfo.BankName = "КРАСНОДАРСКОЕ ОТДЕЛЕНИЕ N 8619"
	insuredperson.MethodReceivePayment.BankInfo.Bik = "040349602"
	insuredperson.MethodReceivePayment.BankInfo.AccountNum = "40817810430851481443"

	xmlText, err := xml.MarshalIndent(insuredperson, " ", " ")
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}
	fmt.Printf("\n--- Marshal ---\n\n")
	fmt.Printf("%s\n", string(xmlText))
}
