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
	ledPins = []rpio.Pin{
		rpio.Pin(17),
		rpio.Pin(18),
		rpio.Pin(27),
		rpio.Pin(22),
		rpio.Pin(23),
		rpio.Pin(24),
		rpio.Pin(25),
		rpio.Pin(2),
		rpio.Pin(3),
		rpio.Pin(8),
	}
	animationTime = 50 * time.Millisecond
)

func main() {
	if err := rpio.Open(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Unmap gpio memory when done
	defer rpio.Close()
	fmt.Println("Initializing pins")
	ledPinAnimation := ledPins

	for _, pin := range ledPins {
		pin.Output()
		pin.High()
	}

	fmt.Println("Program starting")
	for {
		for index := len(ledPins) - 1; index >= 0; index-- {
			for i, ledPin := range ledPinAnimation {
				ledPin.Low()
				time.Sleep(animationTime)
				if i != len(ledPinAnimation)-1 {
					ledPin.High()
				}
			}

			if len(ledPinAnimation) > 0 {
				ledPinAnimation = ledPinAnimation[:len(ledPinAnimation)-1]
			}
		}

		for _, pin := range ledPins {
			ledPinAnimation = append(ledPinAnimation, pin)
			for index := len(ledPinAnimation) - 1; index >= 0; index-- {
				ledPinAnimation[index].Low()
				time.Sleep(animationTime)
				ledPinAnimation[index].High()
			}
		}
	}
}
