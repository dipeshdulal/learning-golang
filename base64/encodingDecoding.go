package base64encdec

import (
	b64 "encoding/base64"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

// EncodeString encodes base64 string
func EncodeString(str string) string {
	sEnc := b64.StdEncoding.EncodeToString([]byte(str))
	return sEnc
}

// DecodeString decodes base64 string
func DecodeString(str string) (string, error) {
	b, err := b64.StdEncoding.DecodeString(str)
	return string(b), err
}

// FileUpload input struct
type FileUpload struct {
	Name   string `json:"name" binding:"required"`
	Base64 string `json:"base64" binding:"required"`
}

// HandleBase64FileUpload handles file uploads
// done in string base 63 format
func HandleBase64FileUpload(c *gin.Context) {
	var input FileUpload

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
			"error":   true,
		})
		return
	}

	byteArray, err := b64.StdEncoding.DecodeString(input.Base64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"error":   true,
		})
		return
	}

	err = ioutil.WriteFile(input.Name, byteArray, 0777)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"error":   true,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":    "file uploaded successfully",
		"byteLength": len(byteArray),
	})
}

// HandleBase64OctetStream handles sending octet stream to client
func HandleBase64OctetStream(c *gin.Context) {

	c.Header("Content-Type", "application/octet-stream")

	c.Stream(func(w io.Writer) bool {
		b, err := ioutil.ReadFile("sample")
		if err != nil {
			fmt.Println("Error: ", err.Error())
			return true
		}
		var n int
		n, err = w.Write(b)
		if err != nil {
			fmt.Println("Error: ", err.Error())
			return true
		}
		fmt.Printf("Written %v bytes", n)
		return false
	})
}
