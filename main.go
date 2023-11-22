package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/your-username/go-mysql-api/database"
)

type MemberType struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func main() {

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	r := gin.Default()
	r.Use(cors.Default())
	r.GET("/getProvince", getProvince)
	r.Run("localhost:3000")

}

func getProvince(c *gin.Context) {
	db := database.Connect()
	var posts []MemberType
	result, err := db.Query("SELECT ProvinceId, NameProvince FROM Province")
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		var tag MemberType
		err = result.Scan(&tag.ID, &tag.Name)
		if err != nil {
			panic(err.Error())
		}
		log.Printf(tag.Name)
		posts = append(posts, tag)
	}
	json.NewEncoder(c.Writer).Encode(posts)

}
