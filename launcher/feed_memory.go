package launcher

import (
	"log"
	"math/rand"
	"redis_benchmark/commands"
	"strconv"
	"time"
)

// Configure the count of feedMemory-operators
const operators int = 2

// Using for control operators
var feedMemorySwitch bool
var feedMemoryStopFlag chan bool

func init() {
	feedMemoryStopFlag = make(chan bool, operators)
}

func (launcher Launcher) StartFeedMemory() {
	if feedMemorySwitch {
		log.Fatal("FeedMemory has already started.")
	}
	feedMemorySwitch = true
	for i := 0; i < operators; i++ {
		// Ensure each operator has the different key
		head := strconv.Itoa(rand.Int())
		go feedMemory(launcher, head, 0)
	}
}

func (launcher Launcher) StopFeedMemory() {
	if !feedMemorySwitch {
		log.Fatal("FeedMemory has not started.")
	}
	feedMemorySwitch = false
	// Hold here, waiting for all operators are stopped
	for i := 0; i < operators; i++ {
		<-feedMemoryStopFlag
	}
}

func feedMemory(launcher Launcher, head string, step int) {
	uniqueKey := head + "&" + strconv.Itoa(step) + "&" + strconv.Itoa(rand.Int())
	cmd := commands.New(commands.SET, uniqueKey, commands.MaxSize)
	launcher.runOnly(cmd)
	// Sleep 10 milliseconds
	duration, err := time.ParseDuration("10ms")
	if err != nil {
		log.Fatal(err)
	}
	time.Sleep(duration)
	if !feedMemorySwitch {
		// Release the stop-caller
		feedMemoryStopFlag <- true
		return
	}
	feedMemory(launcher, head, step+1)
}
