package cmd

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "pgscale",
	Short: "Step-by-step cli tool help assistant for Postgres scaling tasks",
}

func Execute() {
	cobra.CheckErr(RootCmd.Execute())
}
