package loginit

import (
	logrustash "github.com/bshuster-repo/logrus-logstash-hook"
	"github.com/sirupsen/logrus"
	"log"
	"net"
)

func LogInit(connType string, logstashAdress string, appName string) *logrus.Logger {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetLevel(logrus.InfoLevel)
	conn, err := net.Dial(connType, logstashAdress)
	if err != nil {
		log.Fatal(err)
	}
	hook := logrustash.New(conn, logrustash.DefaultFormatter(logrus.Fields{"type": appName}))

	logger.Hooks.Add(hook)
	return logger
}
