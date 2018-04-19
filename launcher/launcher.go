package launcher

import (
	"fmt"
	"github.com/mediocregopher/radix.v2/cluster"
	"log"
	"redis_benchmark/commands"
	"time"
)

type Launcher struct {
	radix *cluster.Cluster
}

// Create launcher
func New(addr string) Launcher {
	radixConn, err := cluster.New(addr)
	if err != nil {
		log.Fatal(err)
	}
	return Launcher{radix: radixConn}
}

// Run benchmark
func (launcher Launcher) Run(cmd *commands.Commands, times int) string {
	before := time.Now()
	for i := 0; i < times; i++ {
		launcher.runOnly(cmd)
	}
	duration := time.Since(before)
	return fmt.Sprintln(cmd, " Ans: ", duration, "/", times, "Times")
}

// Run command
func (launcher Launcher) runOnly(cmd *commands.Commands) {
	launcher.radix.Cmd(cmd.Cmd(), cmd.Args()...)
}

// Close launcher
func (launcher Launcher) Close() {
	launcher.radix.Close()
}
