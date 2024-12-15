package utils

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"sync"
)

func processFile(fileName string, resultsChan chan<- map[string]bool, wg *sync.WaitGroup) {
	defer wg.Done()
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening file")
		return
	}
	defer file.Close()

	localMap := make(map[string]bool)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		localMap[line] = true
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file")
		return
	}

	resultsChan <- localMap
}

func mergeResults(resultsChan <-chan map[string]bool, globalFileCountMap map[string]int, mapMutex *sync.Mutex, wg *sync.WaitGroup) {
	defer wg.Done()

	for localMap := range resultsChan {
		mapMutex.Lock()
		for key := range localMap {
			if len(key) >= 8 && len(key) <= 10 {
				globalFileCountMap[key]++
			}
		}
		mapMutex.Unlock()
	}
}

func getFilePaths(dir string) ([]string, error) {
	var filePaths []string

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			filePaths = append(filePaths, path)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return filePaths, nil
}

func Process() map[string]int {

	filePaths, _ := getFilePaths("coupons")

	fmt.Println("Started Processing Coupon Codes")

	globalFileCountMap := make(map[string]int)
	var mapMutex sync.Mutex

	resultsChan := make(chan map[string]bool, len(filePaths))

	var fileWg sync.WaitGroup

	for _, filePath := range filePaths {
		fileWg.Add(1)
		go processFile(filePath, resultsChan, &fileWg)
	}

	fileWg.Wait()

	close(resultsChan)

	var workerWg sync.WaitGroup
	numWorkers := runtime.NumCPU() * 2

	for i := 0; i < numWorkers; i++ {
		workerWg.Add(1)
		go mergeResults(resultsChan, globalFileCountMap, &mapMutex, &workerWg)
	}

	workerWg.Wait()

	fmt.Println("All files processed successfully")

	return globalFileCountMap

}
