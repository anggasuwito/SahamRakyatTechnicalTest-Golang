package palindrome

import (
	"SahamRakyatTechnicalTest-Golang/api/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

type PalindromeController struct {
	Usecase IPalindromeUsecase
}

func (p PalindromeController) Palindrome(r *gin.Engine) {
	r.POST("/palindrome", p.LongestPalindrome)
}

func (p PalindromeController) LongestPalindrome(g *gin.Context) {
	var body model.Palindrome
	err := g.ShouldBindJSON(&body)
	if err != nil {
		g.JSON(http.StatusBadRequest, model.Response{Message: "Error = " + err.Error()})
		return
	}
	result := p.Usecase.LongestPalindrome(body)
	g.JSON(http.StatusOK, model.Response{
		Message: "Success",
		Result:  result,
	})
}
