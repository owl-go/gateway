package mail

import (
	"testing"
)

func TestSendMail(t *testing.T) {
	cfg := &MailConfig{
		Hostname: "smtp.office365.com",
		Port:     "587",
		Username: "xxxxxx",
		Password: "xxxxxx",
	}
	mail := NewMail(cfg)
	err := mail.SendMail("from@hotmail.com", "to@hostmail.com", "hello", "test")
	if err != nil {
		t.Error(err)
	}
}
