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
	master, _ := core.GetConfig(args[0])
	masterSSH, _ := core.GetConfig(args[0], "ssh")

	replica, _ := core.GetConfig(args[0])
	replicaSSH, _ := core.GetConfig(args[0], "ssh")

	fmt.Println(master)
	fmt.Println(masterSSH)
	fmt.Println(replica)
	fmt.Println(replicaSSH)
}

func init() {
	RootCmd.AddCommand(replicateCmd)
}
