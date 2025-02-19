package main

import (
	"encoding/xml"
	"fmt"
	"main/internal/model"
)

func main() {
	insuredperson := model.NewInsuredPerson("02126176506", "231101593917")
	xmlText, err := xml.MarshalIndent(insuredperson, " ", " ")
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}
	fmt.Printf("\n--- Marshal ---\n\n")
	fmt.Printf("xml: %s\n", string(xmlText))
}
