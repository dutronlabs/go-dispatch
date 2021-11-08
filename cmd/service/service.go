package service

import (
	"os"
	"syscall"
	"go.uber.org/zap"
	"os/signal"
	"github.com/spf13/viper"
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

func (s *Service) SetHealthCheckStatus(status Status){
	s.hcStatusChannel <- status
}

func (s *Service) Start(v *viper.Viper) error {
	return nil
}