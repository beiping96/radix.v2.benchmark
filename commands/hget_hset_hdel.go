package commands

// HGET Key Field Value
func hget(key string) *Commands {
	return &Commands{cmd: "HGET", args: []interface{}{key, "field"}}
}

// HSET Key Field Value
func hset(key string, sizeBytes uint64) *Commands {
	value := generateString(sizeBytes)
	return &Commands{cmd: "HSET", args: []interface{}{key, "field", value}}
}

// HDEL Key Field
func hdel(key string) *Commands {
	return &Commands{cmd: "HDEL", args: []interface{}{key, "field"}}
}
