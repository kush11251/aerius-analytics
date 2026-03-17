package controllers
import (
	"aerius-analytics/src/models"
	"aerius-analytics/src/utils"
	"fmt"
)
func GetMetrics() {
	models.InitDatabase()
	utils.AnalyzeData()
	fmt.Println("Metrics retrieved")
}