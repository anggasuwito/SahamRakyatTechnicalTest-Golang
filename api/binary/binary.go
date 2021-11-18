package binary

import (
	"SahamRakyatTechnicalTest-Golang/api/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

type BinaryController struct {
	Usecase IBinaryUsecase
}

func (b BinaryController) Binary(r *gin.Engine) {
	r.POST("/binary", b.BinaryConverter)
}

func (b BinaryController) BinaryConverter(g *gin.Context) {
	var body model.Binary
	var result interface{}
	output := g.Query("output")
	err := g.ShouldBindJSON(&body)
	if err != nil {
		g.JSON(http.StatusBadRequest, model.Response{Message: "Error = " + err.Error()})
		return
	}
	if output=="binary"{
		result = b.Usecase.DecimalToBinary(body)
	} else {
		result = b.Usecase.BinaryToDecimal(body)
	}

	g.JSON(http.StatusOK, model.Response{
		Message: "Success",
		Result:  result,
	})
}
