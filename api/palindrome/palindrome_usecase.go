package palindrome

import "SahamRakyatTechnicalTest-Golang/api/model"

type IPalindromeUsecase interface {
	LongestPalindrome(body model.Palindrome) string
}

type PalindromeUsecase struct{}

func NewPalindromeUsecase() IPalindromeUsecase {
	return &PalindromeUsecase{}
}
