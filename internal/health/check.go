package health

import (
	"net/http"
	"os/exec"
)

type HealthResult struct {
	Name   string
	Status string
}

func RunChecks() []HealthResult {
	var results []HealthResult

	// 1. Check Tomcat
	tomcatStatus := "Unknown"
	if err := exec.Command("systemctl", "is-active", "tomcat9").Run(); err == nil {
		tomcatStatus = "Running"
	} else {
		tomcatStatus = "Stopped"
	}
	results = append(results, HealthResult{Name: "Tomcat9", Status: tomcatStatus})

	// 2. Check MySQL
	mysqlStatus := "Unknown"
	if err := exec.Command("mysqladmin", "ping").Run(); err == nil {
		mysqlStatus = "Running"
	} else {
		mysqlStatus = "Stopped"
	}
	results = append(results, HealthResult{Name: "MySQL", Status: mysqlStatus})

	// 3. Optional: Check OpenMRS API
	apiStatus := "Unknown"
	resp, err := http.Get("http://localhost:8080/openmrs/ws/rest/v1/system/version")
	if err == nil && resp.StatusCode == 200 {
		apiStatus = "Healthy"
	} else {
		apiStatus = "Unreachable"
	}
	results = append(results, HealthResult{Name: "OpenMRS API", Status: apiStatus})

	return results
}
