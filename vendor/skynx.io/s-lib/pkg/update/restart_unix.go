//go:build !windows
// +build !windows

package update

import (
	"fmt"
	"os"
	"syscall"

	"skynx.io/s-lib/pkg/logging"
	"skynx.io/s-lib/pkg/version"
	"skynx.io/s-lib/pkg/xlog"
)

func restartProcess(app, exe string) {
	if app == version.CLI_NAME {
		fmt.Printf("Binary updated to latest version :-)\n\n")
		os.Exit(0)
	}

	RestartRequest <- struct{}{}

	xlog.Info("Binary updated to latest version, restarting...")

	<-RestartReady

	if err := syscall.Exec(exe, os.Args, os.Environ()); err != nil {
		logging.Alertf("Unable to restart main process: %v", err)
	}
}
