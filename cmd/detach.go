package cmd

import "github.com/spf13/cobra"

var (
	detachCmd = &cobra.Command{
		Use:   `detach <mount device> <nodename>`,
		Short: "detach volume from node",
		RunE:  detachPlugin,
	}
)

func init() {
	RootCmd.AddCommand(detachCmd)
}

func detachPlugin(cmd *cobra.Command, args []string) error {
	return ReplyStr(DetachResponse)
}
