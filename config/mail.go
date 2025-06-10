package config

import (
    
    "net/smtp"
)

// Cấu hình SMTP thật
var (
    SMTPHost = "smtp.gmail.com"
    SMTPPort = "587"
    SMTPUser = "tominhnhat04@gmail.com"       
    SMTPPass = "your-app-password-here"     
)

func SendMail(to string, subject string, body string) error {
    from := SMTPUser
    pass := SMTPPass

    msg := "From: " + from + "\n" +
        "To: " + to + "\n" +
        "Subject: " + subject + "\n\n" +
        body

    auth := smtp.PlainAuth("", from, pass, SMTPHost)
    return smtp.SendMail(SMTPHost+":"+SMTPPort, auth, from, []string{to}, []byte(msg))
}
