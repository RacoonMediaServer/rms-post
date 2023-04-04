package notifier

import "github.com/RacoonMediaServer/rms-post/internal/config"

type Service struct {
	cfg config.Delivery
}

func New(cfg config.Delivery) Service {
	return Service{cfg: cfg}
}
