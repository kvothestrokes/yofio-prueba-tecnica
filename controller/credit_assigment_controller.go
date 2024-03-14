package controller

import (
	"yofio-prueba-tecnica/dto"
	"yofio-prueba-tecnica/service"

	"github.com/gin-gonic/gin"
)

type CreditAssignerController interface {
	AssignCredit(ctx *gin.Context)
	CreditsStatistics(ctx *gin.Context)
}

type creditAssignerController struct {
	creditAssignerService service.CreditAssigner
	saveCreditAssignation service.CreditAssignerDB
}

func NewCreditAssignerController(creditAssignerService service.CreditAssigner) CreditAssignerController {
	return &creditAssignerController{
		creditAssignerService: creditAssignerService,
		saveCreditAssignation: service.NewCreditAssignerDB(),
	}
}

func (c *creditAssignerController) AssignCredit(ctx *gin.Context) {

	//leer request y parsear a dto
	var request = dto.CreditAssigmentRequest{}
	err := ctx.BindJSON(&request)
	if err != nil {
		ctx.JSON(400, gin.H{"error": "petición inválida"})
		return
	}

	//validar el parametro investment
	if request.Investment <= 0 {
		ctx.JSON(400, gin.H{"error": "el monto de inversión es inválido"})
		return
	}

	//Calcular asignación de créditos
	credit_700, credit_500, credit_300, err := c.creditAssignerService.Assign(request.Investment)

	if err != nil {

		//Guardar error en asignación de créditos
		go c.saveCreditAssignation.SaveCreditAssignation(request.Investment, 0)
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	//Guardar asignación de créditos
	go c.saveCreditAssignation.SaveCreditAssignation(request.Investment, 1)

	//Responder con la asignación de créditos
	response := dto.CreditAssigmentResponse{
		CreditType300: credit_300,
		CreditType500: credit_500,
		CreditType700: credit_700,
	}

	ctx.JSON(200, response)

}

func (c *creditAssignerController) CreditsStatistics(ctx *gin.Context) {

	total, err := c.saveCreditAssignation.GetTotalAssignations()
	if err != nil {
		ctx.JSON(400, gin.H{"error": "error obteniendo estadísticas"})
		return
	}

	totalSuccess, err := c.saveCreditAssignation.GetTotalSuccessAssignations()
	if err != nil {
		ctx.JSON(400, gin.H{"error": "error obteniendo estadísticas"})
		return
	}

	totalFailed, err := c.saveCreditAssignation.GetTotalFailedAssignations()
	if err != nil {
		ctx.JSON(400, gin.H{"error": "error obteniendo estadísticas"})
		return
	}

	averageSucceded, err := c.saveCreditAssignation.GetAverageSuccessAssignations()
	if err != nil {
		ctx.JSON(400, gin.H{"error": "error obteniendo estadísticas"})
		return
	}

	averageFailed, err := c.saveCreditAssignation.GetAverageFailedAssignations()
	if err != nil {
		ctx.JSON(400, gin.H{"error": "error obteniendo estadísticas"})
		return
	}

	response := dto.CreditStatisticsResponse{
		TotalAssignations:         total,
		TotalSuccessAssignation:   totalSuccess,
		TotalFailedAssignation:    totalFailed,
		AverageSuccessAssignation: averageSucceded,
		AverageFailedAssignation:  averageFailed,
	}

	ctx.JSON(200, response)

}
