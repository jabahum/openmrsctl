package logs

import (
	"fmt"
	"strings"
)

func filterLogsByGrep(logs string, pattern string) string {
	var filteredLines []string
	for _, line := range strings.Split(logs, "\n") {
		if strings.Contains(line, pattern) {
			filteredLines = append(filteredLines, line)
		}
	}
	return strings.Join(filteredLines, "\n")
}
func filterLogsByLevel(logs string, level string) string {
	minLevel := strings.ToUpper(level)
	var filteredLines []string
	for _, line := range strings.Split(logs, "\n") {
		if shouldInclude(line, minLevel) {
			filteredLines = append(filteredLines, line)
		}
	}
	return strings.Join(filteredLines, "\n")
}

func shouldInclude(line string, minLevel string) bool {
	levelMap := map[string]int{"DEBUG": 1, "INFO": 2, "WARN": 3, "ERROR": 4}

	currentLevel := "DEBUG"
	for l := range levelMap {
		if strings.Contains(line, "["+l+"]") {
			currentLevel = l
			break
		}
	}

	return levelMap[currentLevel] >= levelMap[minLevel]
}

func (d *DockerLogManager) findContainerID(component string) (string, error) {
	switch strings.ToLower(component) {
	case "app", "":
		return "openmrs_app_1", nil
	case "db":
		return "openmrs_db_1", nil
	case "solr":
		return "openmrs_solr_1", nil
	default:
		return "", fmt.Errorf("unknown component: %s", component)
	}
}

func (s *SystemdLogManager) getServiceName(component string) (string, error) {
	switch strings.ToLower(component) {
	case "app", "", "tomcat":
		return "tomcat.service", nil
	case "db", "mysql":
		return "mysql.service", nil
	case "postgresql":
		return "postgresql.service", nil
	default:
		return component + ".service", nil
	}
}
