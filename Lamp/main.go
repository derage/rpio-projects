/*
An example of edge event handling by @Drahoslav7, using the go-rpio library
Waits for button to be pressed twice before exit.
Connect a button between pin 22 and some GND pin.
*/

package main

import (
	"fmt"
	"os"
	"time"

	rpio "github.com/stianeikeland/go-rpio/v4"
)

var (
	// Use mcu pin 22, corresponds to GPIO 3 on the pi
	buttonPin       = rpio.Pin(18)
	ledPin          = rpio.Pin(17)
	buttonState     = rpio.High
	lastButtonState = rpio.High
	lastChangeTime  = timeMilli()
	ledState        = rpio.Low
	captureTime     = int64(50)
)

func main() {
	// Open and map memory to access gpio, check for errors
	if err := rpio.Open(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	// Unmap gpio memory when done
	defer rpio.Close()

	// Set pin to output mode
	ledPin.Output()
	ledPin.Low()

	buttonPin.Input()
	buttonPin.PullUp()

	fmt.Println("press a button")

	for {
		reading := buttonPin.Read()
		if reading != lastButtonState {
			lastChangeTime = timeMilli()
			fmt.Printf("taking time %v\n", lastChangeTime)
			fmt.Printf("taking time %v - %v > %v\n", timeMilli(), lastChangeTime, captureTime)
		}
		if timeMilli()-lastChangeTime > captureTime {
			if reading != buttonState {
				buttonState = reading
				if buttonState == rpio.Low {
					fmt.Println("button pressed")
					if ledState == rpio.Low {
						fmt.Println("turn on")
						ledState = rpio.High
					} else {
						fmt.Println("turn off")
						ledState = rpio.Low
					}
					ledPin.Toggle()
				} else {
					fmt.Println("button released")
				}
			}
		}
		ledPin.Write(ledState)
		lastButtonState = reading
	}
}

func timeMilli() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}
