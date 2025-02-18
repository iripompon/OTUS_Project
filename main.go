package main

import (
	"fmt"
	"person"
)

func main() {
	insuredperson := person.NewInsuredPerson("02126176506", "231101593917")
	fmt.Println(insuredperson)
}
