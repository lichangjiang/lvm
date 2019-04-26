package cmd

import "github.com/spf13/cobra"

var (
	waitforattachCmd = &cobra.Command{
		Use:   "waitforattach <mount device> <json params>",
		Short: "wait for attach volume and node ",
		RunE:  waitforattachCmdPlugin,
	}
)

func init() {
	RootCmd.AddCommand(waitforattachCmd)
}

func waitforattachCmdPlugin(cmd *cobra.Command, args []string) error {
	return attachPlugin(cmd, args[1:])
}
