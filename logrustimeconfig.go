package main

import (
	"time"

	log "github.com/sirupsen/logrus"
)

type EasternTimeFormatter struct {
	log.Formatter
}

func (u EasternTimeFormatter) Format(e *log.Entry) ([]byte, error) {
	newYork, err := time.LoadLocation("America/New_York")
	if err != nil {
		return []byte{}, err
	}
	e.Time = e.Time.In(newYork)
	return u.Formatter.Format(e)
}

func main() {
	log.SetFormatter(EasternTimeFormatter{&log.JSONFormatter{}})
	log.Info("the current time in NYC")
}
