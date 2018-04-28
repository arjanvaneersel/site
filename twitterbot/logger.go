package twitterbot

import "github.com/Sirupsen/logrus"

type twitterLogger struct {
	*logrus.Logger
}

func (log *twitterLogger) Critical(args ...interface{})                 { log.Error(args...) }
func (log *twitterLogger) Criticalf(format string, args ...interface{}) { log.Errorf(format, args...) }
func (log *twitterLogger) Notice(args ...interface{})                   { log.Info(args...) }
func (log *twitterLogger) Noticef(format string, args ...interface{})   { log.Infof(format, args...) }
