package connection

import(
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Env struct{
	db *gorm.DB
}

func CreateConnection() *Env {
	fmt.Println("Going to create connection")
	db, err := gorm.Open("postgres", "test.db")
	if err!=nil{
		fmt.Println("Error occured connecting to db",err)
		return err
	}
	defer db.Close()
	env:=&Env{db:db}
	fmt.Println("Connection...",*env)
	return env
}