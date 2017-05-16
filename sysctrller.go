package main

import (
    "fmt"
    "os"
    "os/signal"
    "time"
    VolCtrller  "smartconn.cc/sibolwolf/volumecontroller"
    SysSW       "smartconn.cc/sibolwolf/syssleepwake"
)

func SysSWTest() {
    // Nothing to do
    SysSW.UpdateLockStatus("audiolock", 0)
    time.Sleep(time.Second * 5)
    SysSW.UpdateLockStatus("storysynclock", 1)
    time.Sleep(time.Second * 5)
    SysSW.UpdateLockStatus("storysynclock", 0)
    time.Sleep(time.Second * 1)
    SysSW.UpdateLockStatus("storysynclock", 1)
    time.Sleep(time.Second * 1)
    SysSW.UpdateLockStatus("storydecompresslock", 1)
    time.Sleep(time.Second * 5)
    SysSW.UpdateLockStatus("storysynclock", 0)
    time.Sleep(time.Second * 5)
    SysSW.UpdateLockStatus("storydecompresslock", 0)
    SysSW.UpdateLockStatus("storydecompresslock", 1)
    SysSW.UpdateLockStatus("storydecompresslock", 0)
    SysSW.UpdateLockStatus("storydecompresslock", 1)
    SysSW.UpdateLockStatus("storydecompresslock", 0)
}

func main(){
    // Init syssleepwake package
    fmt.Println("Hello, SysSleepWake")
    SysSW.Init()
    go SysSWTest()

    // Init volumectroller package
    fmt.Println("Hello, VolumeController")
    VolCtrller.Init()

    signalChanel := make(chan os.Signal, 1)
    signal.Notify(signalChanel, os.Interrupt)
    for {
        select {
        case <-signalChanel:
            return
        }
    }
}
