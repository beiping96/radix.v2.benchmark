package commands

// GET Key Value
func get(key string) *Commands {
	return &Commands{cmd: "GET", args: []interface{}{key}}
}

// SET Key Value
func set(key string, sizeBytes uint64) *Commands {
	value := generateString(sizeBytes)
	return &Commands{cmd: "SET", args: []interface{}{key, value}}
}

// DEL Key
func del(key string) *Commands {
	return &Commands{cmd: "DEL", args: []interface{}{key}}
}
