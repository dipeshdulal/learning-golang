package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	base64encdec "wesionary.team/dipeshdulal/golang-test/base64"
	"wesionary.team/dipeshdulal/golang-test/fileupload"
	"wesionary.team/dipeshdulal/golang-test/sendmail"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading env file error: ", err.Error())
	}

	fmt.Println("Hello World")

	router := gin.Default()

	router.GET("/send-email-sync", sendmail.SendMailSyncController)

	router.GET("/send-email-async", sendmail.SendMailAsyncController)

	router.GET("/check-smtp-server", sendmail.SendMailCheckConnection)

	router.POST("/single-file-upload", fileupload.HandleSingleFileUpload)

	router.POST("/multiple-file-upload", fileupload.HandleMultipleFileUpload)

	router.POST("/base64-file-upload", base64encdec.HandleBase64FileUpload)

	router.GET("/base64-octet-stream", base64encdec.HandleBase64OctetStream)

	router.Run(":5000")

}
