package binary

import (
	"SahamRakyatTechnicalTest-Golang/api/model"
	"math"
	"strconv"
	"strings"
)

func (b BinaryUsecase) DecimalToBinary(decimal model.Binary) int {
	tempInt := []int{}
	var tempString string
	var result int
	for {
		tempInt = append(tempInt, decimal.Input%2)
		decimal.Input = decimal.Input / 2
		if decimal.Input <= 0 {
			break
		}
	}

	for i := len(tempInt) - 1; i >= 0; i-- {
		tempString += strconv.Itoa(tempInt[i])
	}

	result, _ = strconv.Atoi(tempString)

	return result
}

func (b BinaryUsecase) BinaryToDecimal(binary model.Binary) int {
	stringBinary := strconv.Itoa(binary.Input)
	splitBinary := strings.Split(stringBinary, "")
	var result int
	var counter int
	for i := len(splitBinary) - 1; i >= 0; i-- {
		if splitBinary[i] == "1" {
			result += int(math.Pow(2, float64(counter)))
		}
		counter++
	}

	return result
}
