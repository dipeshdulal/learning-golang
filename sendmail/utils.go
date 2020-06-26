package sendmail

import (
	"fmt"
	"net/smtp"
	"os"
)

// SendEmail sends email
func SendEmail(title string, body string) {
	auth := smtp.PlainAuth("", os.Getenv("SMTP_USERNAME"), os.Getenv("SMTP_PASSWORD"), os.Getenv("SMTP_HOST"))

	to := []string{"nilkantha.dipesh@gmail.com"}
	msg := []byte("To: test@gmail.com\r\n" +
		"From: dipesh.dulal@wesionary.team \r\n" +
		"Subject: " + title + "\r\n" +
		"\r\n" +
		"Hello random people of the smtp world.\r\n" +
		body + "\r\n" +
		"नेपालिमा पनि मेल पठाउन मिल्ने हुन्छ कि हुन्न भन्दिन पर्यो \r\n")

	portHost := os.Getenv("SMTP_HOST") + ":" + os.Getenv("SMTP_PORT")
	err := smtp.SendMail(portHost, auth, "test@gmail.com", to, msg)
	if err != nil {
		fmt.Println("Error: ", err.Error())
	}
}

func hello(user string) string {
	if len(user) == 0 {
		return "Hello Dude!"
	}
	return fmt.Sprintf("Hello %v!", user)
}
