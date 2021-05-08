package log

import (
    log "github.com/sirupsen/logrus"
)

func a() {
    log.WithFields(log.Fields{
        "json": "json1",
    }).Info("foo bar")

    log.WithField("json", "json1").Infoln("foo bar")
}
