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
	memoryFlag[1024*1024*1024*1] = false //  1GB
	memoryFlag[1024*1024*1024*3] = false //  3GB
	memoryFlag[1024*1024*1024*5] = false //  5GB
	memoryFlag[1024*1024*1024*7] = false //  7GB
	memoryFlag[1024*1024*1024*9] = false //  9GB
	// Do this commands in order
	rules = []rule{
		// 1KB
		rule{commands.SET, "set1", 1024, 1000},
		rule{commands.GET, "set1", 1024, 1000},
		rule{commands.DEL, "set1", 1024, 1000},
		rule{commands.HSET, "hset1", 1024, 1000},
		rule{commands.HGET, "hset1", 1024, 1000},
		rule{commands.HDEL, "hset1", 1024, 1000},
		// 5KB
		rule{commands.SET, "set5", 1024 * 5, 1000},
		rule{commands.GET, "set5", 1024 * 5, 1000},
		rule{commands.DEL, "set5", 1024 * 5, 1000},
		rule{commands.HSET, "hset5", 1024 * 5, 1000},
		rule{commands.HGET, "hset5", 1024 * 5, 1000},
		rule{commands.HDEL, "hset5", 1024 * 5, 1000},
		// 10KB
		rule{commands.SET, "set10", 1024 * 10, 1000},
		rule{commands.GET, "set10", 1024 * 10, 1000},
		rule{commands.DEL, "set10", 1024 * 10, 1000},
		rule{commands.HSET, "hset10", 1024 * 10, 1000},
		rule{commands.HGET, "hset10", 1024 * 10, 1000},
		rule{commands.HDEL, "hset10", 1024 * 10, 1000},
		// 30KB
		rule{commands.SET, "set30", 1024 * 30, 1000},
		rule{commands.GET, "set30", 1024 * 30, 1000},
		rule{commands.DEL, "set30", 1024 * 30, 1000},
		rule{commands.HSET, "hset30", 1024 * 30, 1000},
		rule{commands.HGET, "hset30", 1024 * 30, 1000},
		rule{commands.HDEL, "hset30", 1024 * 30, 1000},
		// 50KB
		rule{commands.SET, "set50", 1024 * 50, 1000},
		rule{commands.GET, "set50", 1024 * 50, 1000},
		rule{commands.DEL, "set50", 1024 * 50, 1000},
		rule{commands.HSET, "hset50", 1024 * 50, 1000},
		rule{commands.HGET, "hset50", 1024 * 50, 1000},
		rule{commands.HDEL, "hset50", 1024 * 50, 1000},
		// 70KB
		rule{commands.SET, "set70", 1024 * 70, 1000},
		rule{commands.GET, "set70", 1024 * 70, 1000},
		rule{commands.DEL, "set70", 1024 * 70, 1000},
		rule{commands.HSET, "hset70", 1024 * 70, 1000},
		rule{commands.HGET, "hset70", 1024 * 70, 1000},
		rule{commands.HDEL, "hset70", 1024 * 70, 1000},
		// 90KB
		rule{commands.SET, "set90", 1024 * 90, 1000},
		rule{commands.GET, "set90", 1024 * 90, 1000},
		rule{commands.DEL, "set90", 1024 * 90, 1000},
		rule{commands.HSET, "hset90", 1024 * 90, 1000},
		rule{commands.HGET, "hset90", 1024 * 90, 1000},
		rule{commands.HDEL, "hset90", 1024 * 90, 1000},
	}
}
