package binary

import "SahamRakyatTechnicalTest-Golang/api/model"

type IBinaryUsecase interface {
	DecimalToBinary(decimal model.Binary) int
	BinaryToDecimal(binary model.Binary) int
}

type BinaryUsecase struct{}

func NewBinaryUsecase() IBinaryUsecase {
	return &BinaryUsecase{}
}
