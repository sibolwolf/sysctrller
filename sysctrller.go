package main

import (
    "log"
    "os"
    "os/signal"
    "time"
    VolCtrller  "smartconn.cc/sibolwolf/volumecontroller"
    SysSW       "smartconn.cc/sibolwolf/syssleepwake"
)

func Init() {
    // Init syssleepwake package
    log.Println("Hello, Init SysSleepWake ...")
    SysSW.Init()

    // Init volumectroller package
    log.Println("Hello, Init VolumeController ...")
    VolCtrller.Init()
}

func SysLockUpdate(lockname string, value int) {
    /*
    cameralock
    audiolock
    storydownloadlock
    storydecompresslock
    storysynclock
    */
    SysSW.UpdateLockStatus(lockname, value)
}

func SysSWTest() {
    // Nothing to do
    SysLockUpdate("audiolock", 0)
    time.Sleep(time.Second * 5)
    SysLockUpdate("storysynclock", 1)
    time.Sleep(time.Second * 5)
    SysLockUpdate("storysynclock", 0)
    time.Sleep(time.Second * 1)
    SysLockUpdate("storysynclock", 1)
    time.Sleep(time.Second * 1)
    SysLockUpdate("storydecompresslock", 1)
    time.Sleep(time.Second * 5)
    SysLockUpdate("storysynclock", 0)
    time.Sleep(time.Second * 5)
    SysLockUpdate("storydecompresslock", 0)
    SysLockUpdate("storydecompresslock", 1)
    SysLockUpdate("storydecompresslock", 0)
    SysLockUpdate("storydecompresslock", 1)
    SysLockUpdate("storydecompresslock", 0)
}

func main() {
    log.Println("Hello, this is sysctrller module ...")
    Init()

    // Test module
    go SysSWTest()

    signalChanel := make(chan os.Signal, 1)
    signal.Notify(signalChanel, os.Interrupt)
    for {
        select {
        case <-signalChanel:
            return
        }
    }
}
