package cmd

import "github.com/spf13/cobra"

var RootCmd = &cobra.Command{
	Use: 		   "lvm",
	Short:         "LVM Flex volume plugin",
	SilenceErrors: true,
	SilenceUsage:  true,
}
