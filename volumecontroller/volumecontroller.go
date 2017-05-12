package main

// package name is smartconn.cc/sibolwolf/volumecontroller
import (
    "fmt"
    "os"
    "os/signal"
    LED "smartconn.cc/sibolwolf/sysctrller/volumecontroller/led"
    VOL "smartconn.cc/sibolwolf/sysctrller/volumecontroller/volume"
    KEY "smartconn.cc/liugen/input"
)

var curvolclass string

func Init() {
    curvolclass = VOL.GetVolumeClass()
    // fmt.Println(curvolclass)
    switch curvolclass {
    case "class0":
        LED.UpdateLed(0)
    case "class1":
        LED.UpdateLed(1)
    case "class2":
        LED.UpdateLed(2)
    case "class3":
        LED.UpdateLed(3)
    default:
        LED.UpdateLed(0)
    }

}

func UpdateVol() {
    if curvolclass == "class0" {
        curvolclass = "class1"
        VOL.SetVolumeClass(curvolclass)
        LED.UpdateLed(1)
    } else if curvolclass == "class1" {
        curvolclass = "class2"
        VOL.SetVolumeClass(curvolclass)
        LED.UpdateLed(2)
    } else if curvolclass == "class2" {
        curvolclass = "class3"
        VOL.SetVolumeClass(curvolclass)
        LED.UpdateLed(3)
    } else if curvolclass == "class3" {
        curvolclass = "class0"
        VOL.SetVolumeClass(curvolclass)
        LED.UpdateLed(0)
    } else {
        curvolclass = "class0"
        VOL.SetVolumeClass(curvolclass)
        LED.UpdateLed(0)
    }
}

func main() {
    fmt.Println("Hello, VolumeController")

    Init()

    KEY.Connect("readingangel")
    KEY.GetButton("volume").OnPress(func() {
            fmt.Println("RA got a short key press event for volume")
            UpdateVol()
        })

    signalChanel := make(chan os.Signal, 1)
	signal.Notify(signalChanel, os.Interrupt)
	for {
		select {
		case <-signalChanel:
			return
		}
	}
}
