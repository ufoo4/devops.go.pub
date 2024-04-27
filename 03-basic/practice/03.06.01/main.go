package main

import "fmt"

type HealthCheck struct {
	ServiceID int
	status    string
}

const (
	PassStatus = "pass"
	FailStatus = "fail"
)

func GenerateCheck() []HealthCheck {
	var checks []HealthCheck
	for i := 0; i <= 5; i++ {
		if i%2 == 0 && i != 0 {
			checks = append(checks, HealthCheck{i, PassStatus})
		} else {
			checks = append(checks, HealthCheck{i, FailStatus})
		}
	}
	return checks
}
func main() {
	for _, v := range GenerateCheck() {
		if v.status == PassStatus {
			fmt.Println("Проверка пройдена у:", v.ServiceID)
		}
	}
}
