package main

import (
	"github.com/aaron2198/vts_broker/logger"

	"gorm.io/gorm"
)

type Env struct {
	Conf   *Config
	Logger *logger.VTSBLogger
	Db     *gorm.DB
}
