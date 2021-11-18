package palindrome

import (
	"SahamRakyatTechnicalTest-Golang/api/model"
	"fmt"
	"strings"
)

func (l PalindromeUsecase) LongestPalindrome(body model.Palindrome) string {
	maxLen := len(body.Word)

	for i := maxLen; i > 0; i-- {
		var wordA string
		var wordB string
		for j := 0; j+i <= maxLen; j++ {
			tempA := j
			tempB := j + i - 1
			wordA = strings.ToLower(fmt.Sprintf("%s", body.Word[tempA]))
			wordB = strings.ToLower(fmt.Sprintf("%s", body.Word[tempB]))

			for tempA < tempB && wordA == wordB {
				tempA++
				tempB--
			}

			if tempA >= tempB {
				return body.Word[j : j+i]
			}
		}
		if wordA == " " || wordB == " " {
			break
		}
	}

	return body.Word[0:1]
}
