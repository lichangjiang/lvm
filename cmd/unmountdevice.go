package cmd

import (
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
)

var (
	unmountDeviceCmd = &cobra.Command{
		Use:   "unmountdevice <mount dir>",
		Short: "unmount dir",
		RunE:  unmountDevicePlugin,
	}
)

func init() {
	RootCmd.AddCommand(unmountDeviceCmd)
}

func unmountDevicePlugin(cmd *cobra.Command, args []string) error {

	mntPath := args[0]
	if fInfo, err := os.Lstat(mntPath); err != nil {
		return err
	} else if !fInfo.IsDir() {
		ReplyStr(SuccessResponse)
		return nil
	}

	if r, err := isMounted(mntPath); err != nil || r == false {
		ReplyStr(SuccessResponse)
		return nil
	}

	execCmd := exec.Command(`umount %s &> /dev/null`, mntPath)
	if _, err := execCmd.Output(); err != nil {
		return errors.New(fmt.Sprintf("Failed to unmount volume at %s", mntPath))
	}

	return ReplyStr(SuccessResponse)
}
