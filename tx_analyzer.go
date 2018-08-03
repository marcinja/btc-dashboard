package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sync"
	"time"
)

var N_WORKERS int
var DB_USED string
var BACKUP_JSON bool

const N_WORKERS_DEFAULT = 2
const DB_WAIT_TIME = 30
const MIN_DIST_FROM_TIP = 6
const MAX_ATTEMPTS = 3 // max number of DB write attempts before giving up

const CURRENT_VERSION_NUMBER = 1

const WORKER_PROGRESS_DIR_RELATIVE = "/worker-progress"

var WORKER_PROGRESS_DIR string

func main() {
	nWorkersPtr := flag.Int("workers", N_WORKERS_DEFAULT, "Number of concurrent workers.")
	startPtr := flag.Int("start", -1, "Starting blockheight.")
	endPtr := flag.Int("end", -1, "Last blockheight to analyze.")

	// Flags for different modes of operation. Default is to live analysis/back-filling.
	updateColPtr := flag.Bool("update", false, "Set to true to add a column (you need to change bits of code first)")
	insertPtr := flag.Bool("insert-json", false, "Set to true to insert .json data files into PostgreSQL")
	recoveryFlagPtr := flag.Bool("recovery", false, "Set to true to start workers on files in ./worker-progress")
	jsonPtr := flag.Bool("json", true, "Set to false to stop json logging in /db-backup")
	flag.Parse()

	// Set global variables
	N_WORKERS = *nWorkersPtr
	BACKUP_JSON = *jsonPtr

	// Create directory for json files.
	if BACKUP_JSON {
		currentDir, err := os.Getwd()
		if err != nil {
			log.Fatal(err)
		}

		JSON_DIR = currentDir + JSON_DIR_RELATIVE
		if _, err := os.Stat(JSON_DIR); os.IsNotExist(err) {
			log.Printf("Creating json backup directory at: %v\n", JSON_DIR)
			err := os.Mkdir(JSON_DIR, 0777)
			if err != nil {
				log.Fatal(err)
			}
		}
	}

	currentDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	// If there is no worker-progress directory, then there aren't any failures :)
	WORKER_PROGRESS_DIR := currentDir + WORKER_PROGRESS_DIR_RELATIVE
	if _, err := os.Stat(WORKER_PROGRESS_DIR); os.IsNotExist(err) {
		createDirIfNotExist(WORKER_PROGRESS_DIR)
	}

	if *updateColPtr {
		addColumn()
		return
	}

	if *insertPtr {
		toPostgres()
		return
	}

	var workers chan *Dashboard
	if *recoveryFlagPtr {
		workers, start := startWorkersOnLogs()
		*startPtr = start
	} else {
		workers = setupFreshWorkers()
	}

	analyze(*startPtr, *endPtr, workers)
}

func analyze(start, end int, workers chan *Dashboard) {
	// Borrow a worker to get the current block count.
	dash := <-workers
	blockCount, err := dash.client.GetBlockCount()
	if err != nil {
		log.Fatal(err)
	}
	workers <- dash // and put him back to work.

	if start > int(blockCount) || end > int(blockCount) {
		log.Fatal("Your node hasn't reached that blockheight! Try a smaller -start and/or -end value")
	}

	var nextHeight int64
	if start == -1 {
		nextHeight = blockCount - MIN_DIST_FROM_TIP
	} else {
		nextHeight = int64(start)
	}

	var wg sync.WaitGroup
	heightInRangeOfTip := (blockCount - nextHeight) <= 6
	nBusyWorkers := 0
	for int(nextHeight) <= end { // if doing live analysis, end will be set to -1
		var dash *Dashboard

		// Check if any workers are free.
		select {
		case freeDash := <-workers:
			dash = freeDash
		default:
			time.Sleep(50 * time.Millisecond)
			continue
		}

		if heightInRangeOfTip {
			time.Sleep(500 * time.Millisecond)
			blockCount, err = dash.client.GetBlockCount()
			if err != nil {
				log.Fatal(err)
			}
			workers <- dash
			continue
		}

		wg.Add(1)
		go func(dash *Dashboard, height int64) {
			dash.analyzeBlock(height)
			newLogProgressToFile(int(height), dash.workFile)
			workers <- dash
		}(dash, nextHeight)

		nextHeight += 1
		heightInRangeOfTip = (blockCount - nextHeight) <= MIN_DIST_FROM_TIP
	}

	wg.Wait()
	for i := 0; i < N_WORKERS; i++ {
		dash := <-workers
		dash.shutdown()
	}
}

// analyzeBlock uses the getblockstats RPC to compute metrics of a single block.
// It then stores the results in a batchpoint in the Dashboard's influx client.
func (dash *Dashboard) analyzeBlock(blockHeight int64) {
	// Use getblockstats RPC and merge results into the metrics struct.
	blockStatsRes, err := dash.client.GetBlockStats(blockHeight, nil)
	if err != nil {
		log.Fatal(err)
	}

	blockStats := BlockStats{blockStatsRes}

	dash.batchInsert(blockStats)

	log.Printf("Done with block %v after %v\n", blockHeight, time.Since(start))
}

func startWorkersOnLogs() (chan *Dashboard, int) {
	workers := setupFreshWorkers()

	files, err := ioutil.ReadDir(WORKER_PROGRESS_DIR)
	if err != nil {
		log.Fatal(err)
	}

	heights, maxHeight, prevFiles := getProgressInfo()

	var wg sync.WaitGroup
	wg.Add(len(files))
	i := 0 // index into files, incremented at bottom of loop.
	for i < len(heights) {
		var dash *Dashboard

		// Check if any workers are free.
		select {
		case freeDash := <-workers:
			dash = freeDash
		default:
			time.Sleep(100 * time.Millisecond)
			continue
		}

		go func(dash *Dashboard, height int, fileName string) {
			dash.analyzeBlock(int64(height))

			newLogProgressToFile(height, dash.workFile)

			if err = os.Remove(fileName); err != nil {
				log.Println("Error removing file ", prevFiles[i])
			}

			workers <- dash
			wg.Done()
		}(dash, heights[i], prevFiles[i])

		i++
	}
	wg.Wait()

	log.Println("Finished with Recovery.")

	return workers, maxHeight + 1
}

func setupFreshWorkers() chan *Dashboard {
	formattedTime := time.Now().Format("01-02:15:04")
	log.Println("Starting Workers on Logs: ", formattedTime)

	workers := make(chan *Dashboard, N_WORKERS)

	// Fill up worker channel with free Dashboards ready to go.
	for i := 0; i < N_WORKERS; i++ {
		dash := setupDashboard(formattedTime, i)
		workers <- &dash
	}

	log.Println("Finished setting up new workers.")

	return workers
}
