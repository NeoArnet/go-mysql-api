package province

import (
	"encoding/json"
	"log"

	database "go-mysql-api/database"
	model "go-mysql-api/model"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func GetProvince(c *gin.Context) {
	db := database.Connect()
	var provinces []model.Province
	result, err := db.Query("SELECT ProvinceId, NameProvince FROM Province")
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		var tag model.Province
		err = result.Scan(&tag.ID, &tag.Name)
		if err != nil {
			panic(err.Error())
		}
		log.Printf(tag.Name)
		provinces = append(provinces, tag)
	}
	json.NewEncoder(c.Writer).Encode(provinces)

}

func GetProvinceAll(c *gin.Context) {
	db := database.Connect()
	var provinces []model.ProvinceAll
	result, err := db.Query("call SP_GetProvinceByAll();")
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		var tag model.ProvinceAll
		err = result.Scan(&tag.Name)
		if err != nil {
			panic(err.Error())
		}
		log.Printf(tag.Name)
		provinces = append(provinces, tag)
	}
	json.NewEncoder(c.Writer).Encode(provinces)

}
