package server

import (
	"context"
	rms_users "github.com/RacoonMediaServer/rms-packages/pkg/service/rms-users"
	"github.com/RacoonMediaServer/rms-post/internal/server/models"
	"github.com/RacoonMediaServer/rms-post/internal/server/restapi/operations"
	"github.com/go-openapi/errors"
	"net/http"
)

func (s *Server) configureAPI(api *operations.ServerAPI) {
	api.KeyAuth = func(token string) (*models.Principal, error) {
		resp, err := s.Users.GetPermissions(context.Background(), &rms_users.GetPermissionsRequest{Token: token})
		if err != nil {
			s.log.Errorf("Cannot retrieve permissions: %s", err)
			return nil, errors.New(http.StatusForbidden, "Forbidden")
		}
		notifyAllowed := false
		for _, p := range resp.Perms {
			// отдельных прав не делал на уведомления, но нужно
			if p == rms_users.Permissions_SendNotifications {
				notifyAllowed = true
				break
			}
		}
		if !notifyAllowed {
			return nil, errors.New(http.StatusForbidden, "Forbidden")
		}
		return &models.Principal{Token: token}, nil
	}
}
