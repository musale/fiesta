package utils

import (
    "log"
    "net/smtp"

    "bytes"
)

type EmailUser struct {
    Username    string
    Password    string
    EmailServer string
    Port        int
}

type SmtpTemplateData struct {
    From    string
    To      string
    Subject string
    Body    string
}

const emailTemplate = `From: &#123;&#123;.From&#125;&#125;
To: &#123;&#123;.To&#125;&#125;
Subject: &#123;&#123;.Subject&#125;&#125;

&#123;&#123;.Body&#125;&#125;

Sincerely,

&#123;&#123;.From&#125;&#125;
`

const (
    emailUser = "noreply@smsleopard.com"
    emailPass = "autocook25#"
)

func SendMail(subject string, body string, destination string) {

    emailUser := &EmailUser{emailUser, emailPass, "smtp.gmail.com", 587}

    auth := smtp.PlainAuth("",
        emailUser.Username,
        emailUser.Password,
        emailUser.EmailServer,
    )

    var err error
    var doc bytes.Buffer

    context := &SmtpTemplateData{
        &emailUser, destination, subject, body,
    }
    t := template.New("emailTemplate")
    t, err = t.Parse(emailTemplate)
    if err != nil {
        CheckError("error trying to parse mail template", err)
    }
    err = t.Execute(&doc, context)
    if err != nil {
        CheckError("error trying to execute mail template", err)
    }

    err = smtp.SendMail(emailUser.EmailServer+":"+strconv.Itoa(emailUser.Port), // in our case, "smtp.google.com:587"
        auth,
        emailUser.Username,
        []string{"nathanleclaire@gmail.com"},
        doc.Bytes())
    if err != nil {
        log.Print("ERROR: attempting to send a mail ", err)
    }
}
