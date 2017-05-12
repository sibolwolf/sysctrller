package led

import (
    "fmt"
    "os/exec"
)

/*  Tina
1. The led name for RA in tina is below:
fun_led0/  fun_led1/  net_led/   vol_led0/  vol_led1/  vol_led2/

2. The attribution for led are blow:
brightness max_brightness
*/

type ledatt struct {
    brightness      int
    max_brightness  int
}


var ledname = []string{
    "fun_led0",
    "fun_led1",
    "net_led",
    "vol_led0",
    "vol_led1",
    "vol_led2",
}

func GenFilename(led string) string {
    fmt.Println("Gen filename ...")
    switch led {
        case ledname[0]:
            return "/sys/class/leds/fun_led0/brightness"
        case ledname[1]:
            return "/sys/class/leds/fun_led1/brightness"
        case ledname[2]:
            return "/sys/class/leds/net_led/brightness"
        case ledname[3]:
            return "/sys/class/leds/vol_led0/brightness"
        case ledname[4]:
            return "/sys/class/leds/vol_led1/brightness"
        case ledname[5]:
            return "/sys/class/leds/vol_led2/brightness"
        default:
            return "Err"
    }
}

func GetLedStatus(filename string) {
    fmt.Println("Get the vol_ledx status ...")
    cmdget := "cat" + " " + filename
    cmd := exec.Command("/bin/sh", "-c", cmdget)
    bytes, err := cmd.Output()
    if err != nil {
        fmt.Println("outputerr: " + err.Error())
        return
    }

    fmt.Println(filename +" status is: " + string(bytes))
}

func SetLed(filename string, value string) {
    fmt.Println("Set ledx with new value ...")
    cmdset := "echo" + " " + value + " " + ">" + filename
    fmt.Println(cmdset)
    cmd := exec.Command("/bin/sh", "-c", cmdset)
    _, err := cmd.Output()
    if err != nil {
        fmt.Println("outputerr: " + err.Error())
        return
    }
}

func UpdateLed(volclass int) {
    fmt.Println("Update led status according volume class")
    vol0_led_filename := GenFilename("vol_led0")
    vol1_led_filename := GenFilename("vol_led1")
    vol2_led_filename := GenFilename("vol_led2")
    switch volclass {
    case 0:
        SetLed(vol0_led_filename, "0")
        SetLed(vol1_led_filename, "0")
        SetLed(vol2_led_filename, "0")
    case 1:
        SetLed(vol0_led_filename, "1")
        SetLed(vol1_led_filename, "0")
        SetLed(vol2_led_filename, "0")
    case 2:
        SetLed(vol0_led_filename, "1")
        SetLed(vol1_led_filename, "1")
        SetLed(vol2_led_filename, "0")
    case 3:
        SetLed(vol0_led_filename, "1")
        SetLed(vol1_led_filename, "1")
        SetLed(vol2_led_filename, "1")
    default:
        SetLed(vol0_led_filename, "0")
        SetLed(vol1_led_filename, "0")
        SetLed(vol2_led_filename, "0")
    }
}
