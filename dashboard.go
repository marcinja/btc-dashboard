package main

import (
	"github.com/btcsuite/btcd/rpcclient"

	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"

	"fmt"
	"log"
	"os"
	"time"
)

const SHOW_QUERIES = true
const JSON_DIR_RELATIVE = "/db-backup"

var JSON_DIR string

// TODO: refactor all *general* methods on dashboard to collect errors from their specific implementations

// TODO: rename
// A Dashboard contains all the components necessary to make RPC calls to bitcoind, and
// to place data into PostgreSQL.
type Dashboard struct {
	client *rpcclient.Client

	// Fields specifically for PostgreSQL
	pgClient *pg.DB
	pgBatch  dataBatch

	workFile *os.File
}

// Assumes enviroment variables: DB, DB_USERNAME, DB_PASSWORD, BITCOIND_HOST, BITCOIND_USERNAME, BITCOIND_PASSWORD, are all set.
// PostgreSQL and bitcoind should already be started.
func setupDashboard(time string, id int) Dashboard {

	BITCOIND_HOST, ok := os.LookupEnv("BITCOIND_HOST")
	if !ok {
		BITCOIND_HOST = "localhost:8332"
	}

	// Connect to local bitcoin core RPC server using HTTP POST mode.
	connCfg := &rpcclient.ConnConfig{
		Host: BITCOIND_HOST,
		User: os.Getenv("BITCOIND_USERNAME"),
		Pass: os.Getenv("BITCOIND_PASSWORD"),

		HTTPPostMode: true, // Bitcoin core only supports HTTP POST mode
		DisableTLS:   true, // Bitcoin core does not provide TLS by default
	}
	// Notice the notification parameter is nil since notifications are
	// not supported in HTTP POST mode.
	client, err := rpcclient.New(connCfg, nil)
	if err != nil {
		log.Fatal(err)
	}

	DB_ADDR, ok := os.LookupEnv("DB_ADDR")
	if !ok {
		DB_ADDR = "localhost:5432"
	}

	db := pg.Connect(&pg.Options{
		Addr:     DB_ADDR,
		User:     os.Getenv("DB_USERNAME"),
		Password: os.Getenv("DB_PASSWORD"),
		Database: os.Getenv("DB"),
	})

	model := interface{}((*DashboardData)(nil))
	err = db.CreateTable(model, &orm.CreateTableOptions{
		Temp:        false,
		IfNotExists: true,
	})
	if err != nil {
		log.Fatal(err)
	}

	// Prints out the queries created by go-pg.
	if SHOW_QUERIES {
		db.OnQueryProcessed(func(event *pg.QueryProcessedEvent) {
			query, err := event.FormattedQuery()
			if err != nil {
				log.Fatal(err)
			}

			log.Printf("%s %s", time.Since(event.StartTime), query)
		})
	}

	dash := Dashboard{
		client:   client,
		pgClient: db,
		pgBatch: dataBatch{
			versions:          make([]int64, 0),
			dashboardDataRows: make([]DashboardData, 0),
		},
	}

	workFile := fmt.Sprintf("%v/worker-%v-%v", WORKER_PROGRESS_DIR, time, id)

	// Create file to record progress in.
	file, err := os.Create(workFile)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}

	dash.workFile = file

	return dash
}

func (dash *Dashboard) shutdown() {
	dash.workFile.Close()
	dash.client.Shutdown()
	dash.pgClient.Close()
}

// TODO
// Sneak in the batch insert here!!
// inserts a data from a single getblockstats call into the dashboard's DB
func (dash *Dashboard) insert(stats BlockStats) bool {
	data := Data{
		Version:          CURRENT_VERSION_NUMBER,
		DashboardDataRow: stats.transformToDashboardData(),
	}

	err := dash.pgClient.Insert(&data.DashboardDataRow)
	if err != nil {
		log.Fatal("PG database insert failed! ", err)
	}

	log.Printf("\n\n STORED INTO POSTGRESQL \n\n")

	if BACKUP_JSON {
		storeDataAsFile(data)
	}

	return true
}

// setup the insertion of many BlockStats (stored internally)
// uses batch insertion / bulk insertion capabilities of DB_USED whenever possible
func (dash *Dashboard) batchInsert(stats BlockStats) {
	dash.pgBatch.versions = append(dash.pgBatch.versions, CURRENT_VERSION_NUMBER)
	dash.pgBatch.dashboardDataRows = append(dash.pgBatch.dashboardDataRows, stats.transformToDashboardData())
}

// actually do the write of batch created
func (dash *Dashboard) commitBatchInsert() bool {
	err := dash.pgClient.Insert(&dash.pgBatch.dashboardDataRows)
	if err != nil {
		log.Fatal("PG Commit Batch insert failed! ", err)
	}

	log.Printf("\n\n STORED INTO POSTGRESQL \n\n")

	if BACKUP_JSON {
		for i, dashDataRow := range dash.pgBatch.dashboardDataRows {
			storeDataAsFile(Data{
				Version:          dash.pgBatch.versions[i],
				DashboardDataRow: dashDataRow,
			})
		}
	}

	// Reset batch.
	dash.pgBatch.versions = make([]int64, 0)
	dash.pgBatch.dashboardDataRows = make([]DashboardData, 0)

	return true
}
