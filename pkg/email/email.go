package email

import (
	"bytes"
	"crypto/tls"
	"errors"
	"fmt"
	"net"
	"net/smtp"
	"strconv"
	"text/template"

	"github.com/Augenblick-tech/bilibot/pkg/services/email"
)

type loginAuth struct {
	username, password string
}

func LoginAuth(username, password string) smtp.Auth {
	return &loginAuth{username, password}
}

func (a *loginAuth) Start(server *smtp.ServerInfo) (string, []byte, error) {
	return "LOGIN", []byte(a.username), nil
}

func (a *loginAuth) Next(fromServer []byte, more bool) ([]byte, error) {
	if more {
		switch string(fromServer) {
		case "Username:":
			return []byte(a.username), nil
		case "Password:":
			return []byte(a.password), nil
		default:
			return nil, errors.New("unknown from server")
		}
	}
	return nil, nil
}

func SendEmail(UserID uint, subject string, content interface{}) error {
	emailConfig, err := email.GetConfig(UserID)
	if err != nil {
		return err
	}

	conn, err := net.Dial("tcp", emailConfig.Host+":"+strconv.Itoa(emailConfig.Port))
	if err != nil {
		return err
	}

	c, err := smtp.NewClient(conn, emailConfig.Host)
	if err != nil {
		return err
	}

	tlsconfig := &tls.Config{
		ServerName: emailConfig.Host,
	}

	if err := c.StartTLS(tlsconfig); err != nil {
		return err
	}

	auth := LoginAuth(emailConfig.From, emailConfig.Pass)

	if err := c.Auth(auth); err != nil {
		return nil
	}

	t, _ := template.New("email").Parse(`<!DOCTYPE html>
	<html>
	
	<body>
		<div>{{.}}</div>
	</body>
	
	</html>`)

	var body bytes.Buffer

	mimeHeaders := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body.Write([]byte(fmt.Sprintf("Subject: "+subject+" \n%s\n\n", mimeHeaders)))

	t.Execute(&body, content)

	// Sending email.
	err = smtp.SendMail(
		emailConfig.Host+":"+strconv.Itoa(emailConfig.Port),
		auth,
		emailConfig.From,
		[]string{emailConfig.To},
		body.Bytes(),
	)
	if err != nil {
		return err
	}

	return nil
}
