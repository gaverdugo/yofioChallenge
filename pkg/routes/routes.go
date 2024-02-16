package routes

import (
	"net/http"

	creditassigner "github.com/gaverdugo/yofioChallenge/pkg/creditAssigner"
	"github.com/gin-gonic/gin"
)

type CreditAssigner interface {
	Assign(investment int32) (int32, int32, int32, error)
}

type InvestmentReq struct {
	Investment int32
}

func Routes(router *gin.Engine) {
	router.POST("/credit-assignment", creditAssignmentPOST)
}

func creditAssignmentPOST(c *gin.Context) {
	var investment InvestmentReq

	if err := c.BindJSON(&investment); err != nil {
		return
	}

	if investment.Investment <= 0 || investment.Investment%100 != 0 {
		c.JSON(http.StatusBadRequest, map[string]string{"error": "input not valid"})
		return
	}

	var creditAssignerM CreditAssigner = creditassigner.CreditAssigner{}

	assigned1, assigned2, assigned3, err := creditAssignerM.Assign(investment.Investment)

	if err != nil {
		c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	response := map[string]int32{
		"credit_type_300": assigned1,
		"credit_type_500": assigned2,
		"credit_type_700": assigned3,
	}

	c.JSON(http.StatusOK, response)
}
