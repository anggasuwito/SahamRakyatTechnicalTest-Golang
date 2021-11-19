package palindrome

import (
	"SahamRakyatTechnicalTest-Golang/api/model"
	"strings"
)

func (l PalindromeUsecase) LongestPalindrome(body model.Palindrome) string {
	var result string
	var current string
	removedSpace := strings.ReplaceAll(body.Word, " ", "")
	for i := 0; i < len(removedSpace); i++ {
		odd := expandFromCenter(removedSpace, i, i)
		even := expandFromCenter(removedSpace, i, i+1)

		if len(odd) > len(even) {
			current = odd
		} else {
			current = even
		}

		if len(current) > len(result) {
			result = current
		}
	}

	return result
}

func expandFromCenter(word string, left int, right int) string {
	for left >= 0 && right < len(word) && word[left] == word[right] {
		left -= 1
		right += 1
	}

	return word[left+1 : right]
}
