package main

import (
	"redis_benchmark/commands"
)

type rule struct {
	cmd       commands.CommandsEnum
	key       string
	sizeBytes uint64
	times     int
}

var rules []rule

var memoryFlag map[uint64]bool

func init() {
	// When memory arrive memoryFlag
	memoryFlag = make(map[uint64]bool)
	memoryFlag[1024*1024*1024*1] = false  //  1GB
	memoryFlag[1024*1024*1024*2] = false  //  2GB
	memoryFlag[1024*1024*1024*3] = false  //  3GB
	memoryFlag[1024*1024*1024*4] = false  //  4GB
	memoryFlag[1024*1024*1024*5] = false  //  5GB
	memoryFlag[1024*1024*1024*6] = false  //  6GB
	memoryFlag[1024*1024*1024*7] = false  //  7GB
	memoryFlag[1024*1024*1024*8] = false  //  8GB
	memoryFlag[1024*1024*1024*9] = false  //  9GB
	memoryFlag[1024*1024*1024*10] = false // 10GB
	// Do this commands in order
	rules = []rule{
		// 1KB
		rule{cmd: commands.SET, key: "set1", sizeBytes: 1024, times: 1000},
		rule{cmd: commands.GET, key: "set1", sizeBytes: 1024, times: 1000},
		rule{cmd: commands.DEL, key: "set1", sizeBytes: 1024, times: 1000},
		// 3KB
		rule{cmd: commands.SET, key: "set3", sizeBytes: 1024 * 3, times: 1000},
		rule{cmd: commands.GET, key: "set3", sizeBytes: 1024 * 3, times: 1000},
		rule{cmd: commands.DEL, key: "set3", sizeBytes: 1024 * 3, times: 1000},
		// 5KB
		rule{cmd: commands.SET, key: "set5", sizeBytes: 1024 * 5, times: 1000},
		rule{cmd: commands.GET, key: "set5", sizeBytes: 1024 * 5, times: 1000},
		rule{cmd: commands.DEL, key: "set5", sizeBytes: 1024 * 5, times: 1000},
		// 10KB
		rule{cmd: commands.SET, key: "set10", sizeBytes: 1024 * 10, times: 1000},
		rule{cmd: commands.GET, key: "set10", sizeBytes: 1024 * 10, times: 1000},
		rule{cmd: commands.DEL, key: "set10", sizeBytes: 1024 * 10, times: 1000},
	}
}
