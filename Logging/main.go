package main

import (
	log "tutorial/Logging/CustomLogger"
)

func main() {
	log.Info.Println("Test")
	log.Error.Println("Test")
	log.Warn.Println("Test")
}
