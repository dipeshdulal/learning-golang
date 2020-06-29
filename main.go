package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
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

	router.Run(":5000")

}
