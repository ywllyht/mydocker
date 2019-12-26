package main


import (
    log "github.com/sirupsen/logrus"
    "fmt"
    "os/exec"
)

func commitContainer(imageName string){
    //mntURL := "/root/mnt"
    //imageTar := "/root/" + imageName + ".tar"
    mntURL := "/home/liangjie/myproject/golang/projects/mnt"
    imageTar := "/home/liangjie/myproject/golang/projects/" + imageName + ".tar"
    
    fmt.Printf("%s",imageTar)
    if _, err := exec.Command("tar", "-czf", imageTar, "-C", mntURL, ".").CombinedOutput(); err != nil {
        log.Errorf("Tar folder %s error %v", mntURL, err)
    }
}
