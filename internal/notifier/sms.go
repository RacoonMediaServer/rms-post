package notifier

import (
	"context"
	"fmt"
	"github.com/apex/log"
	"io/ioutil"
	"net/http"
	"net/url"
)

func (s Service) SendSMS(ctx context.Context, tel, text string) error {
	const sign = "SMS Aero"

	u := fmt.Sprintf("http://%s:%s@gate.smsaero.ru/v2/sms/send?number=%s&text=%s&sign=%s",
		url.QueryEscape(s.cfg.Sms.Account), s.cfg.Sms.Key, url.QueryEscape(tel), url.QueryEscape(text), url.QueryEscape(sign))
	resp, err := http.Get(u)
	if err != nil {
		return err
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}
	raw, _ := ioutil.ReadAll(resp.Body)
	log.Debugf("send SMS response: %s", string(raw))
	return nil
}
