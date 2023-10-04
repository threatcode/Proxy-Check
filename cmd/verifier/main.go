package main

import (
	"github.com/ThreatCode/Proxy-Check/internal"
	"github.com/ThreatCode/Proxy-Check/internal/model"
	"github.com/ThreatCode/Proxy-Check/internal/processer"
	logger "github.com/sirupsen/logrus"
)

func main() {
	internal.Init()

	// Create task queue and result queue
	var taskQueue = make(chan *model.Task, internal.Options.NumWorkers)
	var resultQueue = make(chan *model.Result)

	// Start loader
	logger.Info("Starting loader...")
	go processer.Loader(internal.Options.InputFilepath, taskQueue, internal.Options.NumWorkers)

	// Start workers
	logger.Infof("Starting %d workers...", internal.Options.NumWorkers)
	for i := 0; i < internal.Options.NumWorkers; i++ {
		go processer.Worker(taskQueue, resultQueue)
	}

	// Start monitor
	go processer.Monitor()

	// Start saver
	logger.Info("Starting saver...")
	processer.Saver(resultQueue, internal.Options.NumWorkers)

	// Print statistics
	internal.State.Snapshot()
}
