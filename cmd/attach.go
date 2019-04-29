package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"strings"
)

type AttachOptions struct {
	VolumeID    string `json:"volumeID"`
	Size        string `json:"size"`
	Volumegroup string `json:"volumegroup"`
	FsType      string `json:"kubernetes.io/fsType"`
}

var (
	attachCmd = &cobra.Command{
		Use:   `attach <json params> <nodename>`,
		Short: "attach vnode with volume",
		RunE:  attachPlugin,
	}
)

func init() {
	RootCmd.AddCommand(attachCmd)
}

func getdevice(attachOptions *AttachOptions) string {
	vid := strings.Replace(attachOptions.VolumeID, "-", "--", -1)
	vg := strings.Replace(attachOptions.Volumegroup, "-", "--", -1)

	return "/dev/mapper/" + vg + "-" + vid
}

func attachPlugin(cmd *cobra.Command, args []string) error {

	var opt = AttachOptions{}

	if err := json.Unmarshal([]byte(args[0]), &opt); err != nil {
		return err
	}

	//size := opt.Size
	device := getdevice(&opt)

	//使用Stat()方法跟踪连接文件
	if devFileInfo, err := os.Stat(device); err != nil {
		return err
	} else if devFileInfo.Mode() & os.ModeDevice == 0 {
		return errors.New(fmt.Sprintf("Volume %s is not block device", device))
	}

	return ReplyStr(fmt.Sprintf(AttachResponse, device))
}
