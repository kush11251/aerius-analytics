package models
import (
	"database/sql"
	_ "github.com/lib/pq"
)
type User struct {
	ID int
	Name string
}
func InitDatabase() {
	db, err := sql.Open("postgres", "user=aerius password=aerius dbname=aerius host=localhost port=5432")
	if err != nil {
		panic(err)
	}
}