package alert

import (
	"fmt"
	"net/smtp"
	"os"
	"strings"
)

func Send(email string,message string,sendEmail chan bool) {
	from := os.Getenv("MAIL")
	password := os.Getenv("PASSWD")


	to := []string{email}

	smtpHost := "smtp.gmail.com"
	smtpPort := 587
	body := ""
	if strings.Contains(message,"warning"){
		body = "<html><body><p style='color:black;font-size:46px;'>Warning!!</p></body></html>"
	} else if strings.Contains(message,"threat"){
		body = "<html><body><p style='color:red;font-size:46px;'>THREAT!!!!!!!!!!</p></body></html>"
	} else {
		body = "<html><body><p style='color:yellow;font-size:46px;'>Alert!!</p></body></html>"
	}

	emailMessage := fmt.Sprintf("To: %s\r\n", to)
    emailMessage += fmt.Sprintf("Subject: %s\r\n", message)
	

	emailMessage += "MIME-Version: 1.0\r\n"
    emailMessage += "Content-Type: text/html; charset=utf-8\r\n"

	emailMessage += "\r\n" + body

	auth := smtp.PlainAuth("",from,password,smtpHost)

	err := smtp.SendMail(fmt.Sprintf("%s:%d", smtpHost, smtpPort), auth, from,to, []byte(emailMessage))

	if err != nil {
		sendEmail <- false
	}

	sendEmail <- true

}