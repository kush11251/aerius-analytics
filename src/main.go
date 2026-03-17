package main
import (
	"aerius-analytics/src/models"
	"aerius-analytics/src/utils"
	"fmt"
)
func main() {
	models.InitDatabase()
	utils.AnalyzeData()
	fmt.Println("Aerius Analytics started")
}