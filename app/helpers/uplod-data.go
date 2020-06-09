package helpers

import (
	"context"
	"database/sql"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"sync"
)

const totalWorker = 100
const csvFile = "majestic_million.csv"

//connection for upload massal file
const connString = "root@tcp(127.0.0.1:3306)/gormdb"

//const connString = "root@gormdb"
const dbMaxIddleConn = 4
const dbMaxConns = 100

func OpenConnection() (*sql.DB, error) {
	log.Println("open connection..")

	db, err := sql.Open("mysql", connString)
	if err != nil {
		return nil, err
	}

	db.SetMaxIdleConns(dbMaxConns)
	db.SetMaxIdleConns(dbMaxIddleConn)

	return db, nil
}

//for flexibility inserting data
func GenerateQuestionsMark(n int) []string {
	s := make([]string, 0)
	for i := 0; i < n; i++ {
		s = append(s, "?")
	}
	return s
}

//open csv file
func OpenCsvFile() (*csv.Reader, *os.File, error) {
	log.Println("open csv file..")

	f, err := os.Open(csvFile)
	if err != nil {
		return nil, nil, err
	}

	reader := csv.NewReader(f)
	return reader, f, nil
}

func RunWorker(db *sql.DB, jobs <-chan []interface{}, wg *sync.WaitGroup) {
	for workerIndex := 0; workerIndex <= totalWorker; workerIndex++ {
		go func(workerIndex int, db *sql.DB, jobs <-chan []interface{}, wg *sync.WaitGroup) {
			counter := 0

			for job := range jobs {
				DoTheJob(workerIndex, counter, db, job)
				wg.Done()
				counter++
			}
		}(workerIndex, db, jobs, wg)
	}

}

var dataHeaders = []string{
	"GlobalRank",
	"TldRank",
	"Domain",
	"TLD",
	"RefSubNets",
	"RefIPs",
	"IDN_Domain",
	"IDN_TLD",
	"PrevGlobalRank",
	"PrevTldRank",
	"PrevRefSubNets",
	"PrevRefIPs",
}

func DoTheJob(workerIndex, counter int, db *sql.DB, values []interface{}) {
	for {
		var outerError error
		func(outerError *error) {
			defer func() {
				if err := recover(); err != nil {
					*outerError = fmt.Errorf("%v", err)
				}
			}()

			conn, err := db.Conn(context.Background())
			query := fmt.Sprintf("INSERT INTO domain (%s) VALUES (%s)",
				strings.Join(dataHeaders, ","),
				strings.Join(GenerateQuestionsMark(len(dataHeaders)), ","),
				)

			//ins := "insert into domain (GlobalRank,TldRank,Domain,TLD,RefSubNets,RefIPs,IDN_Domain,IDN_TLD,PrevGlobalRank,PrevTldRank,PrevRefSubNets,PrevRefIPs)"
			//valIns :="values (?,?,?,?,?,?,?,?,?,?,?,?)"
			//query := ins+valIns

			_, err = conn.ExecContext(context.Background(),query,values)
			if err != nil {
				log.Fatal(err.Error())
			}
		}(&outerError)
		if outerError == nil {
			break
		}
	}
	if counter%100 == 0 {
		log.Println("=>woeker", workerIndex, "inserted", counter, "data")
	}
}

//variabel untuk menampung row pertama saat rad file sebagai header
var dataheaders = make([]string, 0)

func ReadCSVPerLineTheSendToWorker(csvReader *csv.Reader, jobs chan<- []interface{}, wg *sync.WaitGroup) {
	for {
		row, err := csvReader.Read()
		if err != nil {
			if err == io.EOF {
				err = nil
			}
			break
		}

		if len(dataHeaders) == 0 {
			dataHeaders = row
			continue
		}

		rowOrdered := make([]interface{}, 0)
		for _, each := range row {
			rowOrdered = append(rowOrdered, each)
		}

		wg.Add(1)
		jobs <- rowOrdered
	}

	close(jobs)
}
