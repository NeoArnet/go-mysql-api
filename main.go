package main

import (
	"fmt"

	insertupdateall "go-mysql-api/controller/insertupdateall"
	province "go-mysql-api/controller/province"

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

	r.GET("/Getcategories", province.Getcategories)
	r.GET("/Getproducts", province.Getproducts)
	r.GET("/Getproducts/:id", province.GetproductsID)
	r.POST("/CreateCategories", insertupdateall.CreateCategories)
	//r.GET("/getProvince", province.GetProvince)
	r.GET("/GetProvinceAll", province.GetProvinceAll)
	r.Run()
}
