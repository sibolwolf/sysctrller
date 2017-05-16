package main

import (
    "fmt"
    "os"
    "os/signal"
    VolCtrller  "smartconn.cc/sibolwolf/volumecontroller"
    SysSw       "smartconn.cc/sibolwolf/syssleepwake"
)

func main(){
    // Init syssleepwake package
    fmt.Println("Hello, SysSleepWake")
    SysSW.Init()
    SysSw.testLockStatus2()

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
