package main

import (
    "log"
    "os"
    "os/signal"
    "time"
    VolCtrller  "smartconn.cc/sibolwolf/volumecontroller"
    SysSW       "smartconn.cc/sibolwolf/syssleepwake"
    BatteryM    "smartconn.cc/sibolwolf/batterymonitor"
)

func Init() {
    // Init volumectroller package
    log.Println("Hello, Init VolumeController ...")
    VolCtrller.Init()

    // Init syssleepwake package
    log.Println("Hello, Init SysSleepWake ...")
    SysSW.Init()

    // Init batterymonitor package which must after syssleepwake
    // Batterymonitor is running as a daemon program
    log.Println("Hello, Init batterymonitor ...")
    go BatteryM.Init()
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
    SysLockUpdate("audiolock", 1)
    time.Sleep(time.Second * 5)
    SysLockUpdate("audiolock", 0)

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
