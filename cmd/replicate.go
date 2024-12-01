package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"pgscale/core"
)

var replicateCmd = &cobra.Command{
	Use:   "replicate {master_instance_name} {replica_instance_name}",
	Short: "Configure a full Postgres replica for read, backup or disaster recover",
	Args:  cobra.ExactArgs(2),
	Run:   runReplicate,
}

func runReplicate(cmd *cobra.Command, args []string) {
	masterSSH := core.GetConfig(args[0], "ssh")
	fmt.Println(masterSSH["identity"])

}

func init() {
	RootCmd.AddCommand(replicateCmd)
}
