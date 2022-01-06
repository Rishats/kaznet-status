package main

import (
	"fmt"
	"kaznet-status/database"

	"github.com/jasonlvhit/gocron"
)

func tasks() {
	database.CheckAllRegionsStatus()
	gocron.Every(5).Minute().Do(database.CheckAllRegionsStatus)

	// remove, clear and next_run
	_, time := gocron.NextRun()
	fmt.Println(time)

	// function Start start all the pending jobs
	<-gocron.Start()
}

// Endpoint
func main() {
	tasks()
}
