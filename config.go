package main

import (
	"os"
	"strconv"

	"github.com/sirupsen/logrus"
)

// Config struct
type Config struct {
	port          int
	logLevel      logrus.Level
	logLoc        string
	admin_db_host string
	admin_db_user string
	admin_db_pass string
	admin_db_port int
	admin_db_name string
}

// Read config from environment variables
func ReadConfig() *Config {
	c := new(Config)

	// The listening port
	c.port = getPort("VTSB_PORT", 8080)

	// Logging
	c.logLevel = getLogLevel("VTSB_LOGLEVEL")

	// defaults to stdout in logrus
	c.logLoc = os.Getenv("VTSB_LOGLOCATION")

	// Database
	c.admin_db_host = getHostName("VTSB_ADMIN_DB_HOST")
	c.admin_db_port = getPort("VTSB_ADMIN_DB_PORT", 3306)
	c.admin_db_user = os.Getenv("VTSB_ADMIN_DB_USER")
	c.admin_db_pass = os.Getenv("VTSB_ADMIN_DB_PASS")
	c.admin_db_name = os.Getenv("VTSB_ADMIN_DB_NAME")

	return c
}

// getPort from env key, or assign a default value if not set
func getPort(key string, def int) int {
	port := os.Getenv(key)
	portnum, err := strconv.Atoi(port)
	if err != nil {
		portnum = def
		// print key is not a number
		logrus.Warnf("%s is not a number, using default value %d", key, portnum)
	}
	if portnum >= 65535 || portnum <= 0 {
		portnum = def
		logrus.Warnf("Error: %s is not within 0-65535, using default port %d.", key, portnum)
	}
	return portnum
}

// getLogLevel from env key, default: INFO
func getLogLevel(key string) logrus.Level {
	level := os.Getenv(key)
	switch level {
	case "TRACE":
		return logrus.TraceLevel
	case "DEBUG":
		return logrus.DebugLevel
	case "INFO":
		return logrus.InfoLevel
	case "WARN":
		return logrus.WarnLevel
	case "ERROR":
		return logrus.ErrorLevel
	default:
		return logrus.InfoLevel
	}
}

// getHostName from env key, default localhost
func getHostName(key string) string {
	hostname := os.Getenv(key)
	if hostname == "" {
		return "localhost"
	}
	return hostname
}
