package server

import (
	"context"
	"io"
	"net/http"

	rms_users "github.com/RacoonMediaServer/rms-packages/pkg/service/rms-users"
	"github.com/RacoonMediaServer/rms-post/internal/server/models"
	"github.com/RacoonMediaServer/rms-post/internal/server/restapi/operations"
	"github.com/RacoonMediaServer/rms-post/internal/server/restapi/operations/notify"
	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
)

type Notifier interface {
	SendSMS(ctx context.Context, tel, text string) error
	SendEmail(mailTo, subject, body string, images [][]byte) error
}

func (s *Server) sendSMS(params notify.NotifySMSParams, key *models.Principal) middleware.Responder {
	l := s.log.WithField("user", key.Token).WithField("tel", params.To)
	l.Debug("Request")

	err := s.Notifier.SendSMS(params.HTTPRequest.Context(), params.To, params.Text)
	if err != nil {
		l.Errorf("Send SMS failed: %s", err)
		return notify.NewNotifySMSInternalServerError()
	}
	return notify.NewNotifySMSOK()
}

func (s *Server) sendEmail(params notify.NotifyEmailParams, key *models.Principal) middleware.Responder {
	l := s.log.WithField("user", key.Token).WithField("mailTo", params.To).WithField("subject", params.Subject)
	l.Debug("Request")

	var attachments [][]byte

	if params.Attachment != nil {
		defer params.Attachment.Close()
		data, err := io.ReadAll(params.Attachment)
		if err != nil {
			l.Errorf("Cannot read attachment: %s", err)
			return notify.NewNotifyEmailBadRequest()
		}
		attachments = append(attachments, data)
	}

	err := s.Notifier.SendEmail(params.To, params.Subject, params.Text, attachments)
	if err != nil {
		l.Errorf("Send E-Mail failed: %s", err)
		return notify.NewNotifyEmailInternalServerError()
	}
	return notify.NewNotifyEmailOK()
}

func (s *Server) configureAPI(api *operations.ServerAPI) {

	api.NotifyNotifySMSHandler = notify.NotifySMSHandlerFunc(s.sendSMS)
	api.NotifyNotifyEmailHandler = notify.NotifyEmailHandlerFunc(s.sendEmail)

	api.KeyAuth = func(token string) (*models.Principal, error) {
		req := rms_users.CheckPermissionsRequest{
			Token: token,
			Perms: []rms_users.Permissions{rms_users.Permissions_SendNotifications},
		}
		resp, err := s.Users.CheckPermissions(context.Background(), &req)
		if err != nil {
			s.log.Errorf("Cannot retrieve permissions: %s", err)
			return nil, errors.New(http.StatusForbidden, "Forbidden")
		}
		if !resp.Allowed {
			return nil, errors.New(http.StatusForbidden, "Forbidden")
		}
		return &models.Principal{Token: resp.UserId}, nil
	}
}
