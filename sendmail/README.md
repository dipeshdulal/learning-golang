### Sending E-Mail using `net/smtp` package

`net/smtp` is smtp client library for sending email. The code was tested using mail-trap. 

```go
// SendEmail sends email
func SendEmail(title string, body string) {
	auth := smtp.PlainAuth("", os.Getenv("SMTP_USERNAME"), os.Getenv("SMTP_PASSWORD"), os.Getenv("SMTP_HOST"))

	to := []string{"nilkantha.dipesh@gmail.com"}
	msg := []byte("To: test@gmail.com\r\n" +
		"From: dipesh.dulal@wesionary.team \r\n" +
		"Subject: " + title + "\r\n" +
		"\r\n" +
		"Hello random people of the smtp world.\r\n" +
		body + "\r\n")

	portHost := os.Getenv("SMTP_HOST") + ":" + os.Getenv("SMTP_PORT")
	err := smtp.SendMail(portHost, auth, "test@gmail.com", to, msg)
	if err != nil {
		fmt.Println("Error: ", err.Error())
	}
}
```

The function `SendEmail` send the email. 
There were some of the problems that I came across while implementing smtp;
- One of the problem that I came across is that `port 25` was closed for mailtrap which I checked using tcp library and dialing. After changing the port mail was sending just fine but It didn't show it in mailtrap.

- After looking at it the message `msg := ...` part of the code needs to have a certain format for it to show in mailtrap. `To`, `From` , `Subject` are all necessary part of smtp body. Because all other programming languages libraries makes it easier to always send mail using simple functions but in `golang` instead of higher level abstraction, lower level functionality were provided due to which I needed to be more carefull when constructing body of email. 