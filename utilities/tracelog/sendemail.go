package tracelog

import (
	"bytes"
	"fmt"
	"net/smtp"
	"text/template"
)

//** GLOBAL VARIABLES

var (
	EmailHost         string // Host address to the email server
	EmailPort         int    // Host port to the email
	EmailUserName     string // The email user for authentication
	EmailPassword     string // The password for authentication
	EmailTo           string // Address to send messages
	EmailAlertSubject string // The subject for email alerts
)

var (
	emailTemplate *template.Template // A template for sending emails
)

// SendEmail will send an email
func SendEmail(sessionId string, subject string, message string) (err error) {
	STARTEDf(sessionId, "SendEmail", "Subject[%s]", subject)

	if emailTemplate == nil {
		emailTemplate = template.Must(template.New("emailTemplate").Parse(emailScript()))
	}

	parameters := &struct {
		From    string
		To      string
		Subject string
		Message string
	}{
		EmailUserName,
		EmailTo,
		subject,
		message,
	}

	emailMessage := new(bytes.Buffer)
	emailTemplate.Execute(emailMessage, parameters)

	auth := smtp.PlainAuth("", EmailUserName, EmailPassword, EmailHost)

	err = smtp.SendMail(fmt.Sprintf("%s:%d", EmailHost, EmailPort), auth, EmailUserName, []string{EmailTo}, emailMessage.Bytes())

	if err != nil {
		COMPLETED_ERROR(err, sessionId, "SendEmail")
		return err
	}

	COMPLETED(sessionId, "SendEmail")
	return err
}

// emailScript returns a template for the email message to be sent
func emailScript() (script string) {
	return `From: {{.From}}
To: {{.To}}
Subject: {{.Subject}}
MIME-version: 1.0
Content-Type: text/html; charset="UTF-8"

<html><body>{{.Message}}</body></html>`
}

// SendProblemEmail sends an email with the slice of problems
func SendProblemEmail(sessionId string, subject string, problems []string) (err error) {
	STARTED(sessionId, "SendProblemEmail")

	// Create a buffer to build the message
	message := new(bytes.Buffer)

	// Build the message
	for _, problem := range problems {
		message.WriteString(fmt.Sprintf("%s<br />", problem))
	}

	// Send the email
	SendEmail(sessionId, subject, message.String())

	COMPLETED(sessionId, "SendProblemEmail")
	return err
}
