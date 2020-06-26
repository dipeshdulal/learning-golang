package sendmail

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

// SendMailSyncController sends email by blocking program execution
func SendMailSyncController(c *gin.Context) {
	start := time.Now()
	SendEmail("Ola", "Ola Comoestas")
	end := time.Now()
	timeTakenStr := fmt.Sprintf("%v", end.Sub(start))
	c.JSON(http.StatusOK, gin.H{
		"message":    "message sent",
		"time-taken": timeTakenStr,
	})
}

// SendMailAsyncController send email by non-blocking program execution
func SendMailAsyncController(c *gin.Context) {
	start := time.Now()
	go SendEmail("Ola", "Ola Comoestas")
	end := time.Now()
	timeTakenStr := fmt.Sprintf("%v", end.Sub(start))
	c.JSON(http.StatusOK, gin.H{
		"message":    "mail putout to goroutines sent. (mail will be sent later)",
		"time-taken": timeTakenStr,
	})
}

// SendMailCheckConnection send mail check connection
func SendMailCheckConnection(c *gin.Context) {
	portHost := os.Getenv("SMTP_HOST") + ":" + os.Getenv("SMTP_PORT")
	start := time.Now()
	conn, err := net.Dial("tcp", portHost)
	if err != nil {
		fmt.Println(err)
	}
	conn.Close()
	end := time.Now()
	timeTakenStr := fmt.Sprintf("%v", end.Sub(start))
	c.JSON(http.StatusOK, gin.H{
		"message":    "check tcp connection",
		"time-taken": timeTakenStr,
	})
}
