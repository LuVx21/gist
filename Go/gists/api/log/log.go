package log

import (
    log "github.com/sirupsen/logrus"
    l "log"
)

func a() {
    log.WithFields(log.Fields{
        "json": "json1",
    }).Info("foo bar")

    log.WithField("json", "json1").Infoln("foo bar")

    l.Println("foo...bar...")
}
