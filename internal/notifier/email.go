package notifier

import (
	"fmt"
	"gopkg.in/gomail.v2"
	"io"
)

func (s Service) SendEmail(mailTo, subject, body string, images [][]byte) error {
	m := gomail.NewMessage()
	m.SetHeader("From", s.cfg.Email.User)
	m.SetHeader("To", mailTo)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)
	for i, img := range images {
		id := fmt.Sprintf("image_%d", i)
		headers := map[string][]string{
			"Content-ID": {id},
		}
		m.Attach(id+".jpg", gomail.SetCopyFunc(func(w io.Writer) error {
			_, err := w.Write(img)
			return err
		}), gomail.SetHeader(headers))
	}

	d := gomail.NewDialer(s.cfg.Email.Host, s.cfg.Email.Port, s.cfg.Email.User, s.cfg.Email.Password)
	return d.DialAndSend(m)
}
