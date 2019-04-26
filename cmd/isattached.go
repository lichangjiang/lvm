package cmd


import "github.com/spf13/cobra"

var (
	isattachedCmd = &cobra.Command{
		Use:   `isattached <json params> <nodename>`,
		Short: "check volume is isattached",
		RunE:  isAttachedPlugin,
	}
)

func init() {
	RootCmd.AddCommand(isattachedCmd)
}

func isAttachedPlugin(cmd *cobra.Command, args []string) error {
	return ReplyStr(IsAttachResponse)
}