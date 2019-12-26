package main

import (
    "github.com/ywllyht/mydocker/container"
    "github.com/ywllyht/mydocker/cgroups/subsystems"
    "github.com/ywllyht/mydocker/cgroups"
    log "github.com/sirupsen/logrus"
    "os"
    "strings"
)


func Run(tty bool, comArray []string, res *subsystems.ResourceConfig, volume string) {
    parent, writePipe := container.NewParentProcess(tty, volume)
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
    //mntURL := "/root/mnt"
    //rootURL := "/root"
    mntURL := "/home/liangjie/myproject/golang/projects/mnt/"
    rootURL := "/home/liangjie/myproject/golang/projects/"   
    container.DeleteWorkSpace(rootURL, mntURL, volume)
    os.Exit(0)
}

func sendInitCommand(comArray []string, writePipe *os.File) {
    command := strings.Join(comArray, " ")
    log.Infof("command all is %s", command)
    writePipe.WriteString(command)
    writePipe.Close()
}
