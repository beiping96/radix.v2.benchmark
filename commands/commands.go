package commands

import (
	"code.cloudfoundry.org/bytefmt"
	"fmt"
	"log"
	"math/rand"
)

type Commands struct {
	cmd       string
	args      []interface{}
	sizeBytes uint64
}

// Limit the max size of generateMaster
const MaxSize = 1024 * 1024 // 1MB

// All of generating strings are slices in this array
var generateMaster []byte

// Enum of commands
type CommandsEnum int

const (
	_ CommandsEnum = iota
	GET
	SET
	DEL
	HGET
	HSET
	HDEL
)

// Create Commands
func New(commandEnum CommandsEnum, key string, sizeBytes uint64) (command *Commands) {
	switch commandEnum {
	case GET:
		command = get(key)
	case SET:
		command = set(key, sizeBytes)
	case DEL:
		command = del(key)
	case HGET:
		command = hget(key)
	case HSET:
		command = hset(key, sizeBytes)
	case HDEL:
		command = hdel(key)
	}
	command.sizeBytes = sizeBytes
	return
}

func init() {
	// Initialize generateMaster and fill up
	generateMaster = make([]byte, MaxSize)
	rand.Read(generateMaster)
}

// Implement fmt.Stringer
func (commands Commands) String() string {
	return fmt.Sprintln("Commands: ", commands.cmd, " Size: ", bytefmt.ByteSize(commands.sizeBytes))
}

func (commands *Commands) Cmd() string {
	return commands.cmd
}

func (commands *Commands) Args() []interface{} {
	return commands.args
}

// Make the slices of generateMaster
func generateString(sizeBytes uint64) string {
	if sizeBytes > MaxSize {
		log.Fatal("Can not generate the string is bigger than 1MB.")
	}
	return string(generateMaster[:sizeBytes])
}
