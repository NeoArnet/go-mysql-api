package province

import (
	"encoding/json"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/your-username/go-mysql-api/database"
)

func GetProvince(c *gin.Context) {
	db := database.Connect()
	var provinces []database.Province
	result, err := db.Query("SELECT ProvinceId, NameProvince FROM Province")
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		var tag database.Province
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
	var provinces []database.ProvinceAll
	result, err := db.Query("call SP_GetProvinceByAll();")
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		var tag database.ProvinceAll
		err = result.Scan(&tag.Name)
		if err != nil {
			panic(err.Error())
		}
		log.Printf(tag.Name)
		provinces = append(provinces, tag)
	}
	json.NewEncoder(c.Writer).Encode(provinces)

}
