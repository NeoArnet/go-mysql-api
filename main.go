package main

import (
	"fmt"

	"go-mysql-api/controller/province"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	r := gin.Default()
	r.Use(cors.Default())

	r.GET("/getProvince", province.GetProvince)
	r.GET("/getProvinceAll", province.GetProvinceAll)
	r.Run()

}
