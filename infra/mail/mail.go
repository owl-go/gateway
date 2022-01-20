package mail

import (
	"crypto/tls"
	"errors"
	"fmt"
	"net"
	"net/smtp"
	"time"
)

type MailConfig struct {
	Hostname string
	Port     string
	Username string
	Password string
}

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
			return nil, errors.New("Unknown from server")
		}
	}
	return nil, nil
}

type Mail struct {
	Host string
	Port string
	Auth smtp.Auth
}

func NewMail(cfg *MailConfig) *Mail {
	auth := LoginAuth(cfg.Username, cfg.Password)
	//auth := smtp.PlainAuth("", cfg.Username, cfg.Password, host)
	return &Mail{
		Host: cfg.Hostname,
		Port: cfg.Port,
		Auth: auth,
	}
}

func (m *Mail) SendMail(from, to, subj, msg string) error {
	t := time.Now()
	year, month, day := t.Date()
	curtime := fmt.Sprintf("%d-%d-%d %d:%d", year, month, day, t.Hour(), t.Minute())
	header := make(map[string]string)
	header["From"] = from
	header["To"] = to
	header["Date"] = curtime
	header["Content-Type"] = "text/html;charset=UTF-8"
	header["Subject"] = subj

	body := ""
	for k, v := range header {
		body += fmt.Sprintf("%s:%s\r\n", k, v)
	}
	body += "\r\n" + msg

	host := net.JoinHostPort(m.Host, m.Port)
	conn, err := net.Dial("tcp", host)
	if err != nil {
		return err
	}

	smtpCli, err := smtp.NewClient(conn, m.Host)
	if err != nil {
		return err
	}
	tlsConfig := &tls.Config{ServerName: "smtp.office365.com"}
	if err = smtpCli.StartTLS(tlsConfig); err != nil {
		return err
	}

	defer smtpCli.Close()
	if m.Auth != nil {
		if ok, _ := smtpCli.Extension("AUTH"); ok {
			if err = smtpCli.Auth(m.Auth); err != nil {
				return err
			}
		}
	}
	if err = smtpCli.Mail(from); err != nil {
		return err
	}
	if err = smtpCli.Rcpt(to); err != nil {
		return err
	}
	w, err := smtpCli.Data()
	if err != nil {
		return err
	}
	_, err = w.Write([]byte(body))
	if err != nil {
		return err
	}
	err = w.Close()
	if err != nil {
		return err
	}
	smtpCli.Quit()
	return nil
}
