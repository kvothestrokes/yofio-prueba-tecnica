package main

// Creado por Carlos André Sánchez M. - 14/03/2024
import (
	"fmt"
	"yofio-prueba-tecnica/controller"
	"yofio-prueba-tecnica/service"

	"github.com/gin-gonic/gin"
)

func main() {

	AssignCreditController := controller.NewCreditAssignerController(service.NewCreditAssignerService())

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	//Ruta para calcular la asignación de créditos
	r.POST("/credit-assignment", AssignCreditController.AssignCredit)

	//Ruta para obtener las estadísticas de asignación de créditos
	r.POST("/statistics", AssignCreditController.CreditsStatistics)

	fmt.Println("Servidor corriendo en el puerto 8000...")
	r.Run(":8000")

}
