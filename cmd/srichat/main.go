package main

import (
	"flag"
	"log"
	"os"

	sc "github.com/sky0621/sri-chat"
)

func main() {
	os.Exit(realMain())
}

func realMain() (exitCode int) {
	// treat panic
	defer func() {
		err := recover()
		if err != nil {
			log.Printf("Panic occured. ERR: %+v", err)
			// FIXME 後始末

		}
	}()

	return wrappedMain()
}

func wrappedMain() (exitCode int) {
	configpath := flag.String("f", "./config.toml", "設定ファイル（config.toml）の格納パス")
	flag.Parse()
	ctx, ncErr := sc.NewCtx(*configpath)
	if ncErr != nil {
		return sc.ExitCodeSetupError
	}
	defer ctx.Close()

	exitCode, err := sc.Routing(ctx)
	if err != nil {
		return sc.ExitCodeError
	}

	return exitCode
}
