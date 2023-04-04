package server

import (
	"fmt"
	rms_users "github.com/RacoonMediaServer/rms-packages/pkg/service/rms-users"
	"github.com/RacoonMediaServer/rms-post/internal/server/restapi"
	"github.com/RacoonMediaServer/rms-post/internal/server/restapi/operations"
	"github.com/apex/log"
	"github.com/go-openapi/loads"
)

type Server struct {
	srv *restapi.Server
	log *log.Entry

	Users    rms_users.RmsUsersService
	Notifier Notifier
}

func (s *Server) ListenAndServer(host string, port int) error {
	s.log = log.WithField("from", "api")

	if s.srv == nil {
		swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
		if err != nil {
			return err
		}

		// создаем хендлеры API по умолчанию
		api := operations.NewServerAPI(swaggerSpec)
		s.configureAPI(api)

		// устанавливаем свой логгер
		api.Logger = func(content string, i ...interface{}) {
			s.log.Infof(content, i...)
		}

		// создаем и настраиваем сервер
		s.srv = restapi.NewServer(api)

		// устанавливаем middleware
		s.srv.SetHandler(setupGlobalMiddleware(api.Serve(nil)))
	}

	s.srv.Host = host
	s.srv.Port = port

	if err := s.srv.Listen(); err != nil {
		return fmt.Errorf("cannot start server: %w", err)
	}

	return s.srv.Serve()
}

func (s *Server) Shutdown() error {
	if s.srv != nil {
		return s.srv.Shutdown()
	}

	return nil
}
