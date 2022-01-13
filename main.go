package main

import (
	"fmt"
	"github.com/jasonlvhit/gocron"
	"kaznet-status/database"
	"kaznet-status/pkg/utils"
)

func main() {
	// database init
	memDb := database.InitMemDb()

	go utils.InitPrometheus()
	go database.CheckAllRegionsStatus(memDb)

	// cron jobs
	go gocron.Every(5).Minute().Do(database.CheckAllRegionsStatus, memDb)
	go gocron.Every(30).Second().Do(utils.InitWorker, memDb)

	// remove, clear and next_run
	_, time := gocron.NextRun()
	fmt.Println(time)

	// function Start start all the pending jobs
	<-gocron.Start()
}
