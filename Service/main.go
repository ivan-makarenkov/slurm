package main

type HealthCheck struct {
	ServiceID string `json:"service_id"`
	Status    string `json:"status"`
}

const (
	PassStatus = "pass"
	FailStatus = "fail"
)
