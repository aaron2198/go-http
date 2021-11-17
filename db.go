package main

import (
	"fmt"
	"time"

	"github.com/aaron2198/vts_broker/model"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func (e *Env) DBconnect() {
	conn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		e.Conf.admin_db_user,
		e.Conf.admin_db_pass,
		e.Conf.admin_db_host,
		e.Conf.admin_db_port,
		e.Conf.admin_db_name,
	)
	db, err := gorm.Open(mysql.Open(conn), &gorm.Config{})
	for err != nil {
		e.Logger.Log.WithFields(logrus.Fields{
			"host": e.Conf.admin_db_host,
			"port": e.Conf.admin_db_port,
		}).Error("Failed to connect to database ... retrying in 3 seconds")
		time.Sleep(time.Second * 3)
		db, err = gorm.Open(mysql.Open(conn), &gorm.Config{})
	}
	e.Logger.Log.WithFields(logrus.Fields{
		"host": e.Conf.admin_db_host,
		"port": e.Conf.admin_db_port,
	}).Info("Connected to database")
	e.Db = db
}

// Migrate apply the applications current schema to the DB
func (e *Env) Migrate() {
	err := e.Db.AutoMigrate(
		&model.Community{},
		&model.User{},
	)
	if err != nil {
		e.Logger.Log.WithFields(logrus.Fields{
			"error": err,
		}).Fatal("Failed to migrate database")
	}
}
