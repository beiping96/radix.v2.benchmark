package launcher

import (
	"code.cloudfoundry.org/bytefmt"
	"github.com/mediocregopher/radix.v2/redis"
	"log"
	"strconv"
	"strings"
)

// Get the memory of all nodes used
func (launcher Launcher) MemorySize() (sizeBytes uint64, sizeHuman string) {
	// Find the nodes list
	rsp, err := launcher.radix.Cmd("CLUSTER", "NODES").Str()
	if err != nil {
		log.Fatal(err)
	}
	for _, nodeString := range strings.Split(rsp, "\n") {
		if nodeString != "" {
			// Find the memory of each node used
			nodeUsedMemoryBytes, err := nodeUsedMemory(nodeString)
			if err != nil {
				log.Println(err)
				continue
			}
			sizeBytes += nodeUsedMemoryBytes
		}
	}
	// Make the human-size
	sizeHuman = bytefmt.ByteSize(sizeBytes)
	return
}

func nodeUsedMemory(nodeInfo string) (nodeUsedMemoryBytes uint64, err error) {
	nodeAddrLong := strings.Split(nodeInfo, " ")[1]
	nodeAddr := strings.Split(nodeAddrLong, "@")[0]
	// Replace with `cluster.Cluster` of `redis.Client`
	// `cluster.Cluster` always automatic find the node by key
	nodeClient, err := redis.Dial("tcp", nodeAddr)
	if err != nil {
		return
	}
	defer nodeClient.Close()
	// Just find the `used_memory` in reply
	rsp, err := nodeClient.Cmd("INFO", "MEMORY").Str()
	if err != nil {
		return
	}
	nodeUsedMemoryLong := strings.Split(rsp, "\r\n")[1]
	nodeUsedMemory := strings.Split(nodeUsedMemoryLong, ":")[1]
	nodeUsedMemoryBytes, err = strconv.ParseUint(nodeUsedMemory, 10, 64)
	return
}
