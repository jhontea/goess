package main

import (
	"database/sql"
	"fmt"
	"log"
	"math"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const (
	dbConnString   = "root:password@/sesi8"
	dbMaxIdleConns = 4
	dbMaxConns     = 100
	totalWorker    = 100
	csvFile        = "majestic_million.csv"
)

var dataHeaders = make([]string, 0)

func ExampleMillion() {
	start := time.Now()

	db, err := openDbConnection()
	if err != nil {
		log.Fatal(err.Error())
	}

	csvReader, csvFile, err := openCsvFile()
	if err != nil {
		log.Fatal(err.Error())
	}
	defer csvFile.Close()

	jobs := make(chan []interface{})
	wg := new(sync.WaitGroup)

	go dispatchWorkers(db, jobs, wg)
	readCsvFilePerLineThenSendToWorker(csvReader, jobs, wg)

	wg.Wait()

	duration := time.Since(start)
	fmt.Println("done in", int(math.Ceil(duration.Seconds())), "seconds")
}

func dispatchWorkers(db *sql.DB, jobs <-chan []interface{}, wg *sync.WaitGroup) {
	for workerIndex := 0; workerIndex <= totalWorker; workerIndex++ {
		go func(workerIndex int, db *sql.DB, jobs <-chan []interface{}, wg *sync.WaitGroup) {
			counter := 0

			for job := range jobs {
				doTheJob(workerIndex, counter, db, job)
				wg.Done()
				counter++
			}
		}(workerIndex, db, jobs, wg)
	}
}
