package service

import (
	"sync/atomic"
	"time"
	"go.uber.org/zap"
	"net/http"
	"encoding/json"
	"fmt"
)

type Status int

const (
	Unavailable Status = iota
	Ready
	Broken
)

func (s Status) String() string {
	switch s {
	case Unavailable:
		return "unavailable"
	case Ready:
		return "ready"
	case Broken:
		return "broken"
	default:
		return "unknown"
	}
}

type healthCheckResponse struct {
	statusCode int
	StatusMsg string `json:"status"`
	UpSince time.Time `json:"upSinuce"`
	Uptime	string `json:"uptime"`
}

type HealthCheck struct {
	state atomic.Value
	logger *zap.Logger
	responses map[Status]healthCheckResponse
}

type state struct {
	status  Status
	upSince time.Time
}

// New creates a HealthCheck with the specified initial state.
func New() *HealthCheck {
	hc := &HealthCheck{
		logger: zap.NewNop(),
		responses: map[Status]healthCheckResponse{
			Unavailable: {
				statusCode: http.StatusServiceUnavailable,
				StatusMsg:  "Server not available",
			},
			Ready: {
				statusCode: http.StatusOK,
				StatusMsg:  "Server available",
			},
		},
	}
	hc.state.Store(state{status: Unavailable})
	return hc
}

// SetLogger initializes a logger.
func (hc *HealthCheck) SetLogger(logger *zap.Logger) {
	hc.logger = logger
}

// Handler creates a new HTTP handler.
func (hc *HealthCheck) Handler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		state := hc.getState()
		template := hc.responses[state.status]

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(template.statusCode)

		w.Write(hc.createRespBody(state, template))
	})
}

func (hc *HealthCheck) createRespBody(state state, template healthCheckResponse) []byte {
	resp := template // clone
	if state.status == Ready {
		resp.UpSince = state.upSince
		resp.Uptime = fmt.Sprintf("%v", time.Since(state.upSince))
	}
	healthCheckStatus, _ := json.Marshal(resp)
	return healthCheckStatus
}

// Set a new health check status
func (hc *HealthCheck) Set(status Status) {
	oldState := hc.getState()
	newState := state{status: status}
	if status == Ready {
		if oldState.status != Ready {
			newState.upSince = time.Now()
		}
	}
	hc.state.Store(newState)
	hc.logger.Info("Health Check state change", zap.Stringer("status", status))
}

// Get the current status of this health check
func (hc *HealthCheck) Get() Status {
	return hc.getState().status
}

func (hc *HealthCheck) getState() state {
	return hc.state.Load().(state)
}

// Ready is a shortcut for Set(Ready) (kept for backwards compatibility)
func (hc *HealthCheck) Ready() {
	hc.Set(Ready)
}