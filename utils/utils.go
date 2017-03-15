package utils

import (
    "log"
    "net/mail"
    "net/smtp"

    "github.com/scorredoira/email"
)

func CheckError(message string, err error) {
    if err != nil {
        log.Fatal(message, err)
    }
}

func SendMail(subj, body string, dest []string, fileLoc string) {
    // m := email.NewMessage(subl, body)
    m := email.NewHTMLMessage(subj, body)
    m.From = mail.Address{
        Name: "SMSLeopard NoReply", Address: "noreply@smsleopard.com",
    }
    m.To = dest

    err := m.Attach(fileLoc)
    CheckError("Cannot attach file", err)

    auth := smtp.PlainAuth("", "noreply@smsleopard.com", "autocook25#", "smtp.gmail.com")
    err = email.Send("smtp.gmail.com:587", auth, m)
    CheckError("Cannot send mail:", err)
    return
}
