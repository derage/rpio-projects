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

	"github.com/stianeikeland/go-rpio/v4"
)

var (
	// Use mcu pin 22, corresponds to GPIO 3 on the pi
	buttonPin = rpio.Pin(18)
	ledPin    = rpio.Pin(17)
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
	buttonPin.Detect(rpio.FallEdge) // enable falling edge event detection

	fmt.Println("press a button")

	for i := 0; i < 2; {
		if buttonPin.EdgeDetected() { // check if event occured
			fmt.Println("button pressed")
			i++
			ledPin.Toggle()
		}
		fmt.Println(buttonPin.Read())
		time.Sleep(time.Second / 2)
	}
	buttonPin.Detect(rpio.NoEdge) // disable edge event detection
}
