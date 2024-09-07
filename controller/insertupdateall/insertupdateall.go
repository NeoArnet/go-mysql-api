package insertupdateall

import (
	"go-mysql-api/database"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func CreateCategories(c *gin.Context) {
	db := database.Connect()
	category_name := c.Query("category_name")

	result, err := db.Query("insert into categories(category_name) values (?)", category_name)

	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		panic("Seccess to create the order.\r\n,")
	}

}
