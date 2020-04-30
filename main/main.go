package main

import (
	"database/sql"
	"flag"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	nideshop "nideshop-go"
	"os"
)

const (
	defaultPort              = "8360"
	dbsource = "root:123456@tcp(127.0.0.1:3306)/nideshop"
)
func main() {
	var (
		httpAddr = flag.String("http.addr",":"+defaultPort, "HTTP listen address")
	)
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewSyncLogger(logger)
		logger = log.With(logger,
			"service", "account",
			"time", log.DefaultTimestampUTC,
			"caller", log.DefaultCaller,
		)
	}
	level.Info(logger).Log("msg", "service started")
	defer level.Info(logger).Log("msg", "service ended")
	flag.Parse()
	db, _ := sql.Open("mysql",
		dbsource)

	defer db.Close()
	repo := nideshop.NewRepo(db, logger)
	srv := nideshop.NewLoadMainPageService(repo)

	errs := make(chan error)
	var h http.Handler
	{
		h = nideshop.MakeHTTPHandler(srv)
	}
	errs <-http.ListenAndServe(*httpAddr, h)
	level.Error(logger).Log("exit", <-errs)
}