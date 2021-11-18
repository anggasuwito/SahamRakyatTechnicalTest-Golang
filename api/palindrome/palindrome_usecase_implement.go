package palindrome

import (
	"SahamRakyatTechnicalTest-Golang/api/model"
	"strings"
)

//Belum sesuai soal
func (l PalindromeUsecase) LongestPalindrome(body model.Palindrome) string {
	splittedWord := strings.Split(body.Word, "")
	for i := 0; i <= len(splittedWord)-1; i++ {
		if splittedWord[i] != splittedWord[len(splittedWord)-(i+1)] {
			return "Its Not Palindrome"
		}
	}

	return "Palindrome"
}
