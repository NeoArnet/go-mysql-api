package main

import (
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/your-username/go-mysql-api/controller/province"
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
	r.Run("localhost:3000")

}
