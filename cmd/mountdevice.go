package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
)

var (
	mountDeviceCmd = &cobra.Command{
		Use:   "mountdevice <mount dir> <mount device> <json params>",
		Short: "mount device woth dir",
		RunE:  mountDevicePlugin,
	}
)

func init() {
	RootCmd.AddCommand(mountDeviceCmd)
}

func isMounted(mntPath string) (bool, error) {
	cmd := exec.Command(fmt.Sprintf(`findmnt -n %s 2>/dev/null | cut -d' ' -f1`, mntPath))
	r, err := cmd.Output()
	if err != nil {
		return false, err
	}

	if string(r) == mntPath {
		return true, nil
	} else {
		return false, nil
	}
}

func mountDevicePlugin(cmd *cobra.Command, args []string) error {
	mntPath := args[0]
	dmDev := args[1]
	var opt = AttachOptions{}
	if err := json.Unmarshal([]byte(args[2]), &opt); err != nil {
		return err
	}

	fsType := opt.FsType

	if fInfo, err := os.Lstat(dmDev); err != nil {
		return err
	} else if fInfo.Mode() != os.ModeDevice {
		return errors.New(fmt.Sprintf("%s does not exist", dmDev))
	}

	if r, err := isMounted(mntPath); err != nil {
		return err
	} else if r == true {
		ReplyStr(SuccessResponse)
	}

	execCmd := exec.Command(fmt.Sprintf(`blkid -o udev %s 2>/dev/null|grep "ID_FS_TYPE"|cut -d"=" -f2`, dmDev))
	result, err := execCmd.Output()

	if err != nil {
		return err
	}

	if string(result) == "" {
		execCmd = exec.Command(`mkfs -t %s %s >/dev/null 2>&1`, fsType, dmDev)

		if _, err := execCmd.Output(); err != nil {
			return errors.New(fmt.Sprintf("Failed to create fs %s on device %s with cmd result %s", fsType, dmDev, err.Error()))
		}
	}

	os.MkdirAll(mntPath, os.ModePerm)

	execCmd = exec.Command(fmt.Sprintf(`mount %s %s &> /dev/null`, dmDev, mntPath))
	if _, err := execCmd.Output(); err != nil {
		return errors.New(fmt.Sprintf("Failed to mount device %s at %s with error %s", dmDev, mntPath, err.Error()))
	}

	ReplyStr(SuccessResponse)
	return nil
}
