package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

func storeDataAsFile(data Data) {
	dataFileName := fmt.Sprintf("%v/%v.json", JSON_DIR, data.DashboardDataRow.Height)
	dataFile, err := os.Create(dataFileName)
	if err != nil {
		fmt.Println(err)
	}

	enc := json.NewEncoder(dataFile)
	enc.Encode(data)

	dataFile.Close()
}

// parseProgress takes in the contents of a worker-progress file
// and returns the starting height, the last height completed, and the end height.
func parseProgress(contents string) []int {
	lines := strings.Split(contents, "\n")
	result := make([]int, 0)

	for _, line := range lines {
		split := strings.Split(line, "=")

		if len(split) < 2 {
			continue
		}
		height, err := strconv.Atoi(split[1])
		if err != nil {
			log.Fatal(err)
		}

		result = append(result, height)
	}

	return result
}

// logProgressToFile records the progress of a worker to a given file.
func logProgressToFile(start, last, end int, file *os.File) {
	// Record progress in file.
	progress := fmt.Sprintf("Start=%v\nLast=%v\nEnd=%v", start, last, end)
	_, err := file.WriteAt([]byte(progress), 0)

	if err != nil {
		log.Fatal("Error logging progress: ", err)
	}
}

func newLogProgressToFile(lastBlockAnalyzed int, file *os.File) {
	// Record progress in file.
	progress := fmt.Sprintf("last_block_analyzed=%v", lastBlockAnalyzed)
	_, err := file.WriteAt([]byte(progress), 0)

	if err != nil {
		log.Fatal("Error logging progress: ", err, lastBlockAnalyzed, file.Name())
	}
}

// createDirIfNotExist creates a directory at a given path, unless it already exists.
func createDirIfNotExist(dirPath string) {
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		log.Printf("Creating worker progress directory at: %v\n", dirPath)
		err := os.Mkdir(dirPath, 0777)
		if err != nil {
			log.Fatal(err)
		}
	}
}

// Returns a slice of heights, and the max height in the slice.
func getProgressInfo() ([]int, int, []string) {
	files, err := ioutil.ReadDir(WORKER_PROGRESS_DIR)
	if err != nil {
		log.Fatal(err)
	}

	maxHeight := -1
	heights := make([]int, 0)
	fileNames := make([]int, 0)
	for i, file := range files {
		fileName := WORKER_PROGRESS_DIR + "/" + file.Name()

		contentsBytes, err := ioutil.ReadFile(fileName)
		if err != nil {
			log.Fatal(err)
		}

		if len(contentsBytes) == 0 {
			if err = os.Remove(fileName); err != nil {
				log.Println("Error removing file: ", fileName)
			}
			continue
		}

		contents := string(contentsBytes)
		blockHeightStr := strings.Split(contents, "=")[0]
		blockHeight, err := strconv.Atoi(blockHeightStr)
		if err != nil {
			log.Fatal(err)
		}

		files = append(files, fileName)
		heights = append(heights, blockHeight)
		if blockHeight > maxHeight {
			maxHeight = blockHeight
		}
	}

	return heights, maxHeight, fileNames
}
