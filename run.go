package main

import (
    "github.com/ywllyht/mydocker/container"
    //"github.com/xianlubird/mydocker/container"
    log "github.com/sirupsen/logrus"
    "os"
)


func Run(tty bool, command string) {
    parent := container.NewParentProcess(tty, command)
    if err := parent.Start(); err != nil {
        log.Error(err)
    }
    parent.Wait()
    os.Exit(-1)
}
