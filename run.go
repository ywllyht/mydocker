package main

import (
    "github.com/ywllyht/mydocker/container"
    "github.com/ywllyht/mydocker/cgroups/subsystems"
    "github.com/ywllyht/mydocker/cgroups"
    log "github.com/sirupsen/logrus"
    "os"
    "strings"
)


func Run(tty bool, comArray []string, res *subsystems.ResourceConfig) {
    parent, writePipe := container.NewParentProcess(tty)
    if parent == nil {
        log.Errorf("New parent process error")
        return
    }
    if err := parent.Start(); err != nil {
        log.Error(err)
    }
    // use mydocker-cgroup as cgroup name
    cgroupManager := cgroups.NewCgroupManager("mydocker-cgroup")
    defer cgroupManager.Destroy()
    cgroupManager.Set(res)
    cgroupManager.Apply(parent.Process.Pid)

    sendInitCommand(comArray, writePipe)
    parent.Wait()
    mntURL := "/home/liangjie/myproject/golang/projects/mnt/"    //"/root/mnt/"
    rootURL := "/home/liangjie/myproject/golang/projects/"       //"/root/"
    container.DeleteWorkSpace(rootURL, mntURL)
    os.Exit(0)
}

func sendInitCommand(comArray []string, writePipe *os.File) {
    command := strings.Join(comArray, " ")
    log.Infof("command all is %s", command)
    writePipe.WriteString(command)
    writePipe.Close()
}
