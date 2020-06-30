package structparser

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// JSONType json type
type JSONType struct {
	Name string `json:"name"`
	ID   int    `json:"id"`
}

// GetRequestInput sample model
type GetRequestInput struct {
	Name string   `getParser:"name" json:"name"`
	ID   int      `getParser:"id"  json:"id"`
	JSON JSONType `getParser:"json" json:"json"`
}

// HandleGetRequestParsing is used to check parsing request
func HandleGetRequestParsing(c *gin.Context) {
	var getRequestInput GetRequestInput

	if err := GetShouldBindToStruct(c, &getRequestInput); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   true,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "succesfully bind to getRequestInput",
		"data":    getRequestInput,
	})
}
