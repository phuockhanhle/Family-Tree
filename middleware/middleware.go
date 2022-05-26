package middleware

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	model "github.com/phuockhanhle/familytree/model"
)

func GetTreeByID(driver model.Neo4jDriver) func (c *gin.Context) {	
	return func (c *gin.Context) {
		IDTree := c.Param("id")
		fmt.Println(IDTree)
		tree := driver.RunTransaction(model.GetTreeByID, IDTree)
		c.IndentedJSON(http.StatusOK, tree.(model.Tree).ToJson())
	}
}