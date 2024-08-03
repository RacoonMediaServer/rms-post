package main

import (
	"fmt"
	"net/http"

	"github.com/RacoonMediaServer/rms-packages/pkg/service/servicemgr"
	"github.com/RacoonMediaServer/rms-post/internal/config"
	"github.com/RacoonMediaServer/rms-post/internal/notifier"
	"github.com/RacoonMediaServer/rms-post/internal/server"
	"github.com/apex/log"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/urfave/cli/v2"
	"go-micro.dev/v4"

	// Plugins
	_ "github.com/go-micro/plugins/v4/registry/etcd"
)

var Version = "0.0.0"

func main() {
	log.Infof("rms-post %s", Version)
	defer log.Info("DONE.")

	useDebug := false

	service := micro.NewService(
		micro.Name("rms-post"),
		micro.Version(Version),
		micro.Flags(
			&cli.BoolFlag{
				Name:        "verbose",
				Aliases:     []string{"debug"},
				Usage:       "debug log level",
				Value:       false,
				Destination: &useDebug,
			},
		),
	)

	service.Init(
		micro.Action(func(context *cli.Context) error {
			configFile := "/etc/rms/rms-post.json"
			if context.IsSet("config") {
				configFile = context.String("config")
			}
			return config.Load(configFile)
		}),
	)

	if useDebug {
		log.SetLevel(log.DebugLevel)
	}

	cfg := config.Config()

	srv := server.Server{}
	srv.Users = servicemgr.NewServiceFactory(service).NewUsers()
	srv.Notifier = notifier.New(cfg.Delivery)

	go func() {
		http.Handle("/metrics", promhttp.Handler())
		if err := http.ListenAndServe(fmt.Sprintf("%s:%d", cfg.Monitor.Host, cfg.Monitor.Port), nil); err != nil {
			log.Fatalf("Cannot bind monitoring endpoint: %s", err)
		}
	}()

	if err := srv.ListenAndServer(cfg.Http.Host, cfg.Http.Port); err != nil {
		log.Fatalf("Cannot start web server: %+s", err)
	}
}
