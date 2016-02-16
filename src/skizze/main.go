package main

import (
	"flag"
	"os"
	"strconv"

	_ "net/http/pprof"

	"github.com/njpatel/loggo"
	"golang.org/x/crypto/ssh/terminal"

	"config"
	"server"
)

var logger = loggo.GetLogger("skizze")

func init() {
	setupLoggers()
}

func main() {
	var port uint
	flag.UintVar(&port, "p", 3596, "specifies the port for Skizze to run on")
	flag.Parse()

	//TODO: Add arguments for dataDir and infoDir

	err := os.Setenv("SKIZZE_PORT", strconv.Itoa(int(port)))
	config.PanicOnError(err)

	logger.Infof("Starting Skizze...")
	logger.Infof("Using data dir: %s", config.GetConfig().DataDir)
	if p, err := strconv.Atoi(os.Getenv("SKIZZE_PORT")); err == nil {
		server.Run(uint(p))
	}
}

func setupLoggers() {
	loggerSpec := os.Getenv("SKIZZE_DEBUG")

	// Setup logging and such things if we're running in a term
	if terminal.IsTerminal(int(os.Stdout.Fd())) {
		if loggerSpec == "" {
			loggerSpec = "<root>=DEBUG"
		}
		// As we're in a terminal, let's make the output a little nicer
		_, _ = loggo.ReplaceDefaultWriter(loggo.NewSimpleWriter(os.Stderr, &loggo.ColorFormatter{}))
	} else {
		if loggerSpec == "" {
			loggerSpec = "<root>=WARNING"
		}
	}

	_ = loggo.ConfigureLoggers(loggerSpec)
}
