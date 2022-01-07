package service

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type Service struct {
	AdminPort uint
	// Admin *AdminServer
	Logger *zap.Logger
	// MetricsFactory metrics.Factory
	signalsChannel chan os.Signal
	hcStatusChannel chan Status
	// Mux Router
	Router mux.Router
}

func NewService(adminPort uint) *Service {
	signalsChannel := make(chan os.Signal, 1)
	hcStatusChannel := make(chan Status)
	signal.Notify(signalsChannel, os.Interrupt, syscall.SIGTERM)

	return &Service{
		AdminPort:       adminPort,
		Logger:          &zap.Logger{},
		signalsChannel:  signalsChannel,
		hcStatusChannel: hcStatusChannel,
		Router:          InitializeRouter(),
	}

}

func (s *Service) SetHealthCheckStatus(status Status){
	s.hcStatusChannel <- status
}

func (s *Service) Start(v *viper.Viper) error {
	return nil
}