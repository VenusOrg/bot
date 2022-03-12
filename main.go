package main

import (
	"bot/config"
	"bot/store"
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/ethereum/go-ethereum/cmd/utils"
	"github.com/ethereum/go-ethereum/log"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
)

var (
	cfg *config.Config

	path    = flag.String("config", "./config.json", "config file path")
	logFlag = flag.Int("loglevel", 3, "Log level to use for proxyServer")
	logFile = flag.String("logfile", "", "Tells the bot where to write log entries")
	cycle   = flag.Int("cycle", 60, "Scan at an interval time(minute)")
)

func main() {
	// Parse the flags and set up the logger to print everything requested
	flag.Parse()
	stream := log.StreamHandler(os.Stderr, log.TerminalFormat(true))
	if len(*logFile) != 0 {
		fp, err := os.OpenFile(*logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
		if err != nil {
			fmt.Println("Failed to open log file", "err", err)
			os.Exit(1)
		}
		stream = log.StreamHandler(io.Writer(fp), log.TerminalFormat(true))
	}
	glogger := log.NewGlogHandler(stream)
	glogger.Verbosity(log.Lvl(*logFlag))
	log.Root().SetHandler(glogger)

	var err error
	cfg, err = config.NewConfig(*path)
	if err != nil {
		utils.Fatalf("config load failed", err)
	}

	ctx := context.Background()
	defer ctx.Done()

	go loop2(ctx)

	RunServer(getApis())
}

func getApis() map[string]gin.HandlerFunc {
	return map[string]gin.HandlerFunc{
		"getOrder": func(c *gin.Context) {
			data := c.Query("count")
			id, err := strconv.Atoi(data)
			if err != nil {
				log.Error("Count type is a digital", "err", err)
				c.String(http.StatusBadRequest, "Count type is a digital, err %s", err)
				return
			}
			res, err := store.GetOrder(int64(id))
			if err != nil {
				c.String(http.StatusInternalServerError, "failed to get order", err)
				return
			}
			buf, err := json.Marshal(res)
			if err != nil {
				c.String(http.StatusInternalServerError, "failed to marshal order content", err)
				return
			}
			c.String(http.StatusOK, string(buf))
		},
	}
}

func watcher() {
	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, syscall.SIGINT, syscall.SIGTERM)
	defer signal.Stop(sigc)
	<-sigc
}
