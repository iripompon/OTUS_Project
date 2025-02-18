package main

import (
	"fmt"
	"main/internal/model"
)

func main() {
	insuredperson := model.NewInsuredPerson("02126176506", "231101593917")
	fmt.Println(insuredperson)
}
