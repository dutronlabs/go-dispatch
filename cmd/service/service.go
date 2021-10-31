package service

import (
	"os"
	"syscall"
	"go.uber.org/zap"
	"os/signal"

)

type Service struct {
	AdminPort uint
	// Admin *AdminServer
	Logger *zap.Logger
	// MetricsFactory metrics.Factory
	signalsChannel chan os.Signal
	hcStatusChannel chan Status
}

func NewService(adminPort uint) *Service {
	signalsChannel := make(chan os.Signal, 1)
	hcStatusChannel := make(chan Status)
	signal.Notify(signalsChannel, os.Interrupt, syscall.SIGTERM)

	return &Service{
		// Admin: NewAdminServer(ports.PortToHostPort(adminPort)),
		signalsChannel: signalsChannel,
		hcStatusChannel: hcStatusChannel,
	}
}

