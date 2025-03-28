package serv

import (
	"main/internal/model"
	repo "main/internal/repository"
)

func CreateItem() {
	insuredperson := model.NewInsuredPerson("02126176506", "231101593917")
	putMessageReg := model.NewPutMessage("1.0", "go", "1.23.4", 2)

	repo.CreateSliceOfItem(insuredperson)
	repo.CreateSliceOfItem(putMessageReg)

}
