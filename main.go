package main

import (
	"log"
	"redis_benchmark/commands"
	"redis_benchmark/launcher"
	"time"
)

const addr = "127.0.0.1:9000"

func main() {
	launcherOne := launcher.New(addr)
	defer launcherOne.Close()
	launcherOne.StartFeedMemory()
	defer launcherOne.StopFeedMemory()
	for {
		// Does it have finished all tests?
		var leftFlag []uint64
		for k, v := range memoryFlag {
			if !v {
				leftFlag = append(leftFlag, k)
			}
		}
		if len(leftFlag) == 0 {
			break
		}
		// Is it arrived enough memory?
		memorySize, memorySizeHuman := launcherOne.MemorySize()
		log.Println("Memory arrived: ", memorySizeHuman)
		var arrived bool
		for _, v := range leftFlag {
			if memorySize >= v {
				arrived = true
				memoryFlag[v] = true
			}
		}
		if arrived {
			launcherAllRules(launcherOne)
		}
		// Sleep 3 seconds
		duration, err := time.ParseDuration("3s")
		if err != nil {
			log.Fatal(err)
		}
		time.Sleep(duration)
	}
}

func launcherAllRules(launcherOne launcher.Launcher) {
	for _, rule := range rules {
		command := commands.New(rule.cmd, rule.key, rule.sizeBytes)
		output := launcherOne.Run(command, rule.times)
		log.Println(output)
	}
}
