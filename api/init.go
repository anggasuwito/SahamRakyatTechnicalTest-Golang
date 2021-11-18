package api

import (
	"SahamRakyatTechnicalTest-Golang/api/binary"
	"SahamRakyatTechnicalTest-Golang/api/palindrome"
	"github.com/gin-gonic/gin"
	"os"
)

func Init() {
	r := gin.Default()

	binaryUsecase := binary.NewBinaryUsecase()
	binaryRoute := binary.BinaryController{Usecase: binaryUsecase}
	binaryRoute.Binary(r)

	palindromeUsecase := palindrome.NewPalindromeUsecase()
	palindromeRoute := palindrome.PalindromeController{Usecase: palindromeUsecase}
	palindromeRoute.Palindrome(r)

	r.Run(os.Getenv("PORT"))
}
