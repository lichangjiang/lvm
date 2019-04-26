package main

import (
	"lvm/cmd"
	"os"
	"strings"
)

func main() {

	if err := cmd.RootCmd.Execute(); err != nil {
		if strings.HasPrefix(err.Error(), "unknown command") {
			cmd.ReplyStr(cmd.UnsupportResponse)
		} else {
			cmd.ReplyError(err)
		}

		os.Exit(1)
	}
	os.Exit(0)
}
