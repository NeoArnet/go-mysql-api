package province

import (
	"encoding/json"
	//"log"

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
		//log.Printf(tag.Name)
		provinces = append(provinces, tag)
	}
	json.NewEncoder(c.Writer).Encode(provinces)

}

func GetproductsID(c *gin.Context) {
	db := database.Connect()
	var Product []model.Products
	result, err := db.Query("SELECT product_id, product_name, category_id, price, description, image_path FROM products where product_id=?", c.Param("id"))

	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		var tag model.Products
		err = result.Scan(&tag.PID, &tag.PName, &tag.CID, &tag.Price, &tag.Description, &tag.Image_path)
		if err != nil {
			panic(err.Error())
		}
		//log.Printf(tag.PName)
		Product = append(Product, tag)
	}
	json.NewEncoder(c.Writer).Encode(Product)
}

func Getproducts(c *gin.Context) {
	db := database.Connect()
	var Product []model.Products
	result, err := db.Query("SELECT product_id, product_name, category_id, price, description, image_path FROM products")

	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		var tag model.Products
		err = result.Scan(&tag.PID, &tag.PName, &tag.CID, &tag.Price, &tag.Description, &tag.Image_path)
		if err != nil {
			panic(err.Error())
		}
		//log.Printf(tag.PName)
		Product = append(Product, tag)
	}
	json.NewEncoder(c.Writer).Encode(Product)
}

func Getcategories(c *gin.Context) {
	db := database.Connect()
	var Categorie []model.Categories
	result, err := db.Query("SELECT category_id, category_name FROM categories")
	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		var tag model.Categories
		err = result.Scan(&tag.CID, &tag.CName)
		if err != nil {
			panic(err.Error())
		}
		//log.Printf(tag.CName)
		Categorie = append(Categorie, tag)
	}
	json.NewEncoder(c.Writer).Encode(Categorie)

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
		//log.Printf(tag.Name)
		provinces = append(provinces, tag)
	}
	json.NewEncoder(c.Writer).Encode(provinces)

}
