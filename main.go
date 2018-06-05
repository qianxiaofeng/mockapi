package main

import (
	"mockapi/cmd"
	"mockapi/config"
)

func main() {
	config.BootConfig()
	cmd.Execute()
}
