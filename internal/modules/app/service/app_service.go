package app_service

import "time"

type Service struct{}

var startTime = time.Now()

func Init() *Service {
	return &Service{}
}

type healthCheckResp struct {
	Uptime int
}

func (s *Service) HealthCheck() healthCheckResp {
	return healthCheckResp{
		Uptime: int(time.Since(startTime)),
	}
}
